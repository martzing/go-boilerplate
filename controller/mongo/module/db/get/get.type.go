package dbGet

import (
	"github.com/martzing/go-boilerplate/controller/mongo/module/db/models"
	mongoHelper "github.com/martzing/go-boilerplate/helpers/mongo"
)

type ListCustomerParams struct {
	Search string
	Offset int64
	Limit  int64
	DBTxn  *mongoHelper.DatabaseTransaction
}

type ListCustomerResult struct {
	Count     int64             `json:"count"`
	Customers []models.Customer `json:"customers"`
}
