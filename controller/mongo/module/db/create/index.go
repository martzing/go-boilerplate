package dbCreate

import (
	"context"

	"github.com/martzing/go-boilerplate/controller/mongo/module/db/models"
	mongoHelper "github.com/martzing/go-boilerplate/helpers/mongo"
	utilityHelper "github.com/martzing/go-boilerplate/helpers/utility"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateCustomer(customer models.CustomerCreate, dbTxn *mongoHelper.DatabaseTransaction) mongo.InsertOneResult {
	logger := utilityHelper.NewLogger()
	logger.Infof("Mongo Transaction (%s): Execute CreateCustomer", dbTxn.GetTransactionId())

	coll := dbTxn.DB.Collection("customer")
	result, err := coll.InsertOne(context.TODO(), customer)

	if err != nil {
		panic(err)
	}
	return *result
}
