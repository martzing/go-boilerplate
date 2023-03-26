package dbHelper

import (
	"github.com/google/uuid"
	"github.com/martzing/go-boilerplate/configs"
	utilityHelper "github.com/martzing/go-boilerplate/helpers/utility"
	"gorm.io/gorm"
)

type DatabaseTransaction struct {
	db            *gorm.DB
	transactionId string
}

func NewTransaction() DatabaseTransaction {
	dbTxn := DatabaseTransaction{}

	return dbTxn
}

func executeTransaction(db *gorm.DB, transactionId string) {
	logger := utilityHelper.NewLogger()
	statement := db.Statement
	sql := db.Explain(statement.SQL.String(), statement.Vars...)

	logger.Infof("Transaction (%s): %s", transactionId, sql)
}

func registerCallback(db *gorm.DB, transactionId string) {
	callback := db.Callback()

	callback.Create().Register("execute_transaction", func(db *gorm.DB) {
		executeTransaction(db, transactionId)
	})
	callback.Query().Register("execute_transaction", func(db *gorm.DB) {
		executeTransaction(db, transactionId)
	})
	callback.Update().Register("execute_transaction", func(db *gorm.DB) {
		executeTransaction(db, transactionId)
	})
	callback.Delete().Register("execute_transaction", func(db *gorm.DB) {
		executeTransaction(db, transactionId)
	})
	callback.Raw().Register("execute_transaction", func(db *gorm.DB) {
		executeTransaction(db, transactionId)
	})
	callback.Row().Register("execute_transaction", func(db *gorm.DB) {
		executeTransaction(db, transactionId)
	})
}

func (dbTxn *DatabaseTransaction) Begin() {
	logger := utilityHelper.NewLogger()
	transactionId := uuid.New().String()

	NewDatabase(*configs.DBConfig)

	db := DB

	registerCallback(db, transactionId)

	dbTxn.db = db.Begin()
	dbTxn.transactionId = transactionId

	logger.Infof("Transaction (%s): Start Transaction", dbTxn.transactionId)
}

func (dbTxn *DatabaseTransaction) Get() *gorm.DB {
	return dbTxn.db
}

func (dbTxn *DatabaseTransaction) Commit() {
	logger := utilityHelper.NewLogger()
	db := dbTxn.db

	db.Commit()
	logger.Infof("Transaction (%s): Commit", dbTxn.transactionId)
}

func (dbTxn *DatabaseTransaction) Rollback() {
	logger := utilityHelper.NewLogger()
	db := dbTxn.db

	db.Rollback()
	logger.Infof("Transaction (%s): Rollback", dbTxn.transactionId)
}
