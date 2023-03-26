package dbUpdate

import (
	"github.com/martzing/go-boilerplate/controller/postgres/module/db/models"
	dbHelper "github.com/martzing/go-boilerplate/helpers/db"
)

func UpdateCustomer(customer *models.Customer, dbTxn dbHelper.DatabaseTransaction) {
	db := dbTxn.Get()
	err := db.Save(customer).Error
	if err != nil {
		panic(err)
	}
}
