package dbUpdate

import (
	"context"

	mongoHelper "github.com/martzing/go-boilerplate/helpers/mongo"
	utilityHelper "github.com/martzing/go-boilerplate/helpers/utility"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func UpdateCustomer(customerId string, update primitive.D, dbTxn *mongoHelper.DatabaseTransaction) *mongo.UpdateResult {
	logger := utilityHelper.NewLogger()
	logger.Infof("Mongo Transaction (%s): Execute UpdateCustomer", dbTxn.GetTransactionId())

	coll := dbTxn.DB.Collection("customer")
	id, _ := primitive.ObjectIDFromHex(customerId)
	filter := bson.D{{Key: "_id", Value: id}}
	result, err := coll.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		panic(err)
	}
	return result
}
