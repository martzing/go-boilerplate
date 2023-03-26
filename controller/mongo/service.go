package mongo

import (
	dbCreate "github.com/martzing/go-boilerplate/controller/mongo/module/db/create"
	dbDelete "github.com/martzing/go-boilerplate/controller/mongo/module/db/delete"
	dbGet "github.com/martzing/go-boilerplate/controller/mongo/module/db/get"
	"github.com/martzing/go-boilerplate/controller/mongo/module/db/models"
	dbUpdate "github.com/martzing/go-boilerplate/controller/mongo/module/db/update"
	mongoType "github.com/martzing/go-boilerplate/controller/mongo/type"
	mongoHelper "github.com/martzing/go-boilerplate/helpers/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func listCustomer(params *mongoType.ListCustomerParams) dbGet.ListCustomerResult {
	dbTxn := mongoHelper.NewTransaction()
	defer func() {
		if err := recover(); err != nil {
			dbTxn.Rollback()
			panic(err)
		}
	}()
	dbTxn.Begin()

	listCustomerResult := dbGet.ListCustomer(dbGet.ListCustomerParams{
		Search: params.Search,
		Offset: params.Offset,
		Limit:  params.Limit,
		DBTxn:  dbTxn,
	})
	dbTxn.Commit()
	return listCustomerResult
}

func createCustomer(params *mongoType.CreateCustomerParams) mongoType.CustomerType {
	dbTxn := mongoHelper.NewTransaction()
	defer func() {
		if err := recover(); err != nil {
			dbTxn.Rollback()
			panic(err)
		}
	}()
	dbTxn.Begin()
	_customer := models.CustomerCreate{
		StoreId:    params.StoreId,
		FirstName:  params.FirstName,
		LastName:   params.LastName,
		Email:      params.Email,
		AddressId:  params.AddressId,
		Activebool: params.Activebool,
	}
	result := dbCreate.CreateCustomer(_customer, dbTxn)
	customerId := result.InsertedID.(primitive.ObjectID).Hex()
	customer := dbGet.GetCustomer(customerId, dbTxn)
	dbTxn.Commit()
	s := getCustomerFormat(customerId, customer)
	return s
}

func updateCustomer(params *mongoType.UpdateCustomerParams) models.Customer {
	dbTxn := mongoHelper.NewTransaction()
	defer func() {
		if err := recover(); err != nil {
			dbTxn.Rollback()
			panic(err)
		}
	}()
	dbTxn.Begin()
	var value bson.D
	customer := dbGet.GetCustomer(params.ID, dbTxn)
	if params.StoreId != 0 {
		value = append(value, bson.E{Key: "storeid", Value: params.StoreId})
	}
	if params.FirstName != "" {
		value = append(value, bson.E{Key: "firstname", Value: params.FirstName})
	}
	if params.LastName != "" {
		value = append(value, bson.E{Key: "lastname", Value: params.LastName})
	}
	if params.Email != "" {
		value = append(value, bson.E{Key: "email", Value: params.Email})
	}
	if params.AddressId != 0 {
		value = append(value, bson.E{Key: "addressid", Value: params.AddressId})
	}
	if params.Activebool != nil {
		value = append(value, bson.E{Key: "activebool", Value: *params.Activebool})
	}
	update := bson.D{{Key: "$set", Value: value}}
	updated := dbUpdate.UpdateCustomer(customer.ID, update, dbTxn)
	if updated.ModifiedCount < 1 {
		panic("Update fail")
	}
	result := dbGet.GetCustomer(params.ID, dbTxn)
	dbTxn.Commit()
	return result
}

func deleteCustomer(params *mongoType.DeleteCustomerParams) mongoType.DeleteCustumerResponse {
	dbTxn := mongoHelper.NewTransaction()
	defer func() {
		if err := recover(); err != nil {
			dbTxn.Rollback()
			panic(err)
		}
	}()
	dbTxn.Begin()
	customer := dbGet.GetCustomer(params.ID, dbTxn)
	result := dbDelete.DeleteCustomer(customer.ID, dbTxn)
	dbTxn.Commit()
	if result.DeletedCount < 1 {
		panic("Delete fail")
	}
	return mongoType.DeleteCustumerResponse{
		Status: "success",
	}
}

// func myQueue (data string) {
//     //
// }

// func myJob (data string) {
//     //
// }
