package mongoHelper

import (
	"context"

	"github.com/google/uuid"
	"github.com/martzing/go-boilerplate/configs"
	utilityHelper "github.com/martzing/go-boilerplate/helpers/utility"
	"go.mongodb.org/mongo-driver/mongo"
)

type DatabaseTransaction struct {
	client        *mongo.Client
	DB            *mongo.Database
	ctx           *mongo.SessionContext
	transactionId *string
}

func NewTransaction() *DatabaseTransaction {
	client, db := connect(configs.MongoDB)

	dbTxn := DatabaseTransaction{
		client,
		db,
		nil,
		nil,
	}

	return &dbTxn
}

func (dbTxn *DatabaseTransaction) Begin() {
	logger := utilityHelper.NewLogger()
	transactionId := uuid.New().String()
	if dbTxn.ctx == nil {
		session, err := dbTxn.client.StartSession()

		if err != nil {
			panic(err)
		}

		if err := session.StartTransaction(); err != nil {
			panic(err)
		}

		ctx := mongo.NewSessionContext(context.TODO(), session)

		dbTxn.ctx = &ctx
	}
	dbTxn.transactionId = &transactionId
	logger.Infof("Mongo Transaction (%s): Start Transaction", *dbTxn.transactionId)
}

func (dbTx *DatabaseTransaction) Get() *mongo.SessionContext {
	if dbTx.ctx == nil {
		panic("Getting dbTxn without begin transaction")
	}
	return dbTx.ctx
}

func (dbTx *DatabaseTransaction) GetTransactionId() string {
	if dbTx.transactionId == nil {
		panic("Transaction not begin")
	}
	return *dbTx.transactionId
}

func (dbTxn *DatabaseTransaction) Commit() {
	logger := utilityHelper.NewLogger()
	if dbTxn.ctx != nil {
		ctx := *dbTxn.ctx
		dbTxn.ctx = nil

		err := ctx.CommitTransaction(context.TODO())

		if err != nil {
			panic(err)
		}
		ctx.EndSession(context.TODO())
	}
	logger.Infof("Mongo Transaction (%s): Commit", *dbTxn.transactionId)
}

func (dbTxn *DatabaseTransaction) Rollback() {
	logger := utilityHelper.NewLogger()
	if dbTxn.ctx != nil {
		ctx := *dbTxn.ctx
		dbTxn.ctx = nil

		err := ctx.AbortTransaction(context.TODO())

		if err != nil {
			panic(err)
		}
		ctx.EndSession(context.TODO())
	}
	logger.Infof("Mongo Transaction (%s): Rollback", *dbTxn.transactionId)
}
