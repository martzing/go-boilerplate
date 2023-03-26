package postgres

import (
	dbCreate "github.com/martzing/go-boilerplate/controller/postgres/module/db/create"
	dbDelete "github.com/martzing/go-boilerplate/controller/postgres/module/db/delete"
	dbGet "github.com/martzing/go-boilerplate/controller/postgres/module/db/get"
	"github.com/martzing/go-boilerplate/controller/postgres/module/db/models"
	dbUpdate "github.com/martzing/go-boilerplate/controller/postgres/module/db/update"
	postgresType "github.com/martzing/go-boilerplate/controller/postgres/type"
	dbHelper "github.com/martzing/go-boilerplate/helpers/db"
)

func listCustomer(params *postgresType.ListCustomerParams) postgresType.ListCustomerResponse {
	dbTxn := dbHelper.NewTransaction()
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
	response := listCustomerFormat(listCustomerResult.Count, &listCustomerResult.Customers)
	dbTxn.Commit()
	return response
}

func listFilm(params *postgresType.ListFilmParams) postgresType.ListFilmResponse {
	dbTxn := dbHelper.NewTransaction()
	defer func() {
		if err := recover(); err != nil {
			dbTxn.Rollback()
			panic(err)
		}
	}()
	dbTxn.Begin()

	listFilmResult := dbGet.ListFilm(dbGet.ListFilmParams{
		Search: params.Search,
		Offset: params.Offset,
		Limit:  params.Limit,
		DBTxn:  dbTxn,
	})
	response := listFilmFormat(listFilmResult.Count, &listFilmResult.Films)
	dbTxn.Commit()
	return response
}

func createCustomer(params *postgresType.CreateCustomerParams) models.Customer {
	dbTxn := dbHelper.NewTransaction()
	defer func() {
		if err := recover(); err != nil {
			dbTxn.Rollback()
			panic(err)
		}
	}()
	dbTxn.Begin()
	customer := models.Customer{
		StoreId:    params.StoreId,
		FirstName:  params.FirstName,
		LastName:   params.LastName,
		Email:      params.Email,
		AddressId:  params.AddressId,
		Activebool: params.Activebool,
	}
	dbCreate.CreateCustomer(&customer, dbTxn)
	dbTxn.Commit()
	return customer
}

func updateCustomer(params *postgresType.UpdateCustomerParams) models.Customer {
	dbTxn := dbHelper.NewTransaction()
	defer func() {
		if err := recover(); err != nil {
			dbTxn.Rollback()
			panic(err)
		}
	}()
	dbTxn.Begin()
	customer := dbGet.GetCustomer(params.CustomerId, dbTxn)
	if customer == nil {
		panic("Customer not found")
	}
	if params.StoreId != 0 {
		customer.StoreId = params.StoreId
	}
	if params.FirstName != "" {
		customer.FirstName = params.FirstName
	}
	if params.LastName != "" {
		customer.LastName = params.LastName
	}
	if params.Email != "" {
		customer.Email = params.Email
	}
	if params.AddressId != 0 {
		customer.AddressId = params.AddressId
	}
	if params.Activebool != nil {
		customer.Activebool = *params.Activebool
	}
	dbUpdate.UpdateCustomer(customer, dbTxn)
	dbTxn.Commit()
	return *customer
}

func deleteCustomer(params *postgresType.DeleteCustomerParams) postgresType.DeleteCustumerResponse {
	dbTxn := dbHelper.NewTransaction()
	defer func() {
		if err := recover(); err != nil {
			dbTxn.Rollback()
			panic(err)
		}
	}()
	dbTxn.Begin()
	customer := dbGet.GetCustomer(params.CustomerId, dbTxn)
	if customer == nil {
		panic("Customer not found")
	}
	dbDelete.DeleteCustomer(customer, dbTxn)
	dbTxn.Commit()
	return postgresType.DeleteCustumerResponse{
		Status: "success",
	}
}

func createCustomerQueue(params *postgresType.CreateCustomerParams) {
	dbTxn := dbHelper.NewTransaction()
	defer func() {
		if err := recover(); err != nil {
			dbTxn.Rollback()
			panic(err)
		}
	}()
	dbTxn.Begin()
	customer := models.Customer{
		StoreId:    params.StoreId,
		FirstName:  params.FirstName,
		LastName:   params.LastName,
		Email:      params.Email,
		AddressId:  params.AddressId,
		Activebool: params.Activebool,
	}
	dbCreate.CreateCustomer(&customer, dbTxn)
	dbTxn.Commit()
}

func createCustomerJob(params *postgresType.CreateCustomerParams) {
	dbTxn := dbHelper.NewTransaction()
	defer func() {
		if err := recover(); err != nil {
			dbTxn.Rollback()
			panic(err)
		}
	}()
	dbTxn.Begin()
	customer := models.Customer{
		StoreId:    params.StoreId,
		FirstName:  params.FirstName,
		LastName:   params.LastName,
		Email:      params.Email,
		AddressId:  params.AddressId,
		Activebool: params.Activebool,
	}
	dbCreate.CreateCustomer(&customer, dbTxn)
	dbTxn.Commit()
}
