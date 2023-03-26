package dbGet

import (
	"github.com/martzing/go-boilerplate/controller/postgres/module/db/models"
	dbHelper "github.com/martzing/go-boilerplate/helpers/db"
)

type ListDemoParams struct {
	Search string
	Offset int
	Limit  int
	DBTxn  dbHelper.DatabaseTransaction
}

type ListDemoResult struct {
	Count int64
	Demos []models.Demo
}

type ListCustomerParams struct {
	Search string
	Offset int
	Limit  int
	DBTxn  dbHelper.DatabaseTransaction
}

type ListCustomerResult struct {
	Count     int64
	Customers []models.Customer
}

type ListFilmParams struct {
	Search string
	Offset int
	Limit  int
	DBTxn  dbHelper.DatabaseTransaction
}

type ListFilmResult struct {
	Count int64
	Films []models.Film
}
