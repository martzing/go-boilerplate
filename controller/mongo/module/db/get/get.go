package dbGet

import (
	"context"

	"github.com/martzing/go-boilerplate/controller/mongo/module/db/models"
	mongoHelper "github.com/martzing/go-boilerplate/helpers/mongo"
	utilityHelper "github.com/martzing/go-boilerplate/helpers/utility"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetCustomer(customerId string, dbTxn *mongoHelper.DatabaseTransaction) models.Customer {
	logger := utilityHelper.NewLogger()
	logger.Infof("Mongo Transaction (%s): Execute GetCustomer", dbTxn.GetTransactionId())

	coll := dbTxn.DB.Collection("customer")
	id, err := primitive.ObjectIDFromHex(customerId)
	if err != nil {
		panic("Invalid customer id")
	}
	filter := bson.M{"_id": id}
	customer := models.Customer{}
	err = coll.FindOne(context.TODO(), filter).Decode(&customer)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			panic("Customer not found")
		}
		panic(err)
	}
	return customer
}

func ListCustomer(params ListCustomerParams) ListCustomerResult {
	logger := utilityHelper.NewLogger()
	logger.Infof("Mongo Transaction (%s): Execute ListCustomer", params.DBTxn.GetTransactionId())

	coll := params.DBTxn.DB.Collection("customer")
	opts := options.Find()
	filter := bson.D{}

	if params.Search != "" {
		filter = bson.D{{
			Key: "firstname", Value: primitive.Regex{
				Pattern: params.Search,
				Options: "i",
			},
		}}
	}
	if params.Offset != 0 {
		opts.SetSkip(params.Offset)
	}
	if params.Limit != 0 {
		opts.SetLimit(params.Limit)
	}
	cursor, err := coll.Find(context.TODO(), filter, opts)
	if err != nil {
		panic(err)
	}

	count, _err := coll.CountDocuments(context.TODO(), filter)
	if _err != nil {
		panic(_err)
	}

	customers := []models.Customer{}
	if err = cursor.All(context.TODO(), &customers); err != nil {
		panic(err)
	}

	return ListCustomerResult{
		Count:     count,
		Customers: customers,
	}
}
