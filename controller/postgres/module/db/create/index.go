package dbCreate

import (
	"github.com/martzing/go-boilerplate/controller/postgres/module/db/models"
	dbHelper "github.com/martzing/go-boilerplate/helpers/db"
)

func CreateCustomer(customer *models.Customer, dbTxn dbHelper.DatabaseTransaction) {
	db := dbTxn.Get()
	err := db.Create(customer).Error

	if err != nil {
		panic(err)
	}
}
