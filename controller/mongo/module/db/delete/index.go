package dbDelete

import (
	"context"

	mongoHelper "github.com/martzing/go-boilerplate/helpers/mongo"
	utilityHelper "github.com/martzing/go-boilerplate/helpers/utility"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func DeleteCustomer(customerId string, dbTxn *mongoHelper.DatabaseTransaction) *mongo.DeleteResult {
	logger := utilityHelper.NewLogger()
	logger.Infof("Mongo Transaction (%s): Execute DeleteCustomer", dbTxn.GetTransactionId())

	coll := dbTxn.DB.Collection("customer")
	id, _ := primitive.ObjectIDFromHex(customerId)
	filter := bson.D{{Key: "_id", Value: id}}
	result, err := coll.DeleteOne(context.TODO(), filter)
	if err != nil {
		panic(err)
	}
	return result
}
