package dbGet

import (
	"errors"
	"fmt"

	"github.com/martzing/go-boilerplate/controller/postgres/module/db/models"
	dbHelper "github.com/martzing/go-boilerplate/helpers/db"
	"gorm.io/gorm"
)

func GetCustomer(customerId int, dbTxn dbHelper.DatabaseTransaction) *models.Customer {
	customer := models.Customer{CustomerId: customerId}

	db := dbTxn.Get()

	err := db.First(&customer).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		} else {
			panic(err)
		}
	}
	return &customer
}

func ListCustomer(params ListCustomerParams) *ListCustomerResult {
	customers := []models.Customer{}

	db := params.DBTxn.Get()

	if params.Search != "" {
		db = db.Where("first_name ILIKE ?", fmt.Sprintf("%%%s%%", params.Search))
	}

	var count int64 = 0

	err := db.Offset(params.Offset).Limit(params.Limit).Find(&customers).Count(&count).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		} else {
			panic(err)
		}
	}

	return &ListCustomerResult{
		Count:     count,
		Customers: customers,
	}
}

func ListFilm(params ListFilmParams) *ListFilmResult {
	films := []models.Film{}
	db := params.DBTxn.Get()

	if params.Search != "" {
		db = db.Where("title ILIKE ?", fmt.Sprintf("%%%s%%", params.Search))
	}

	var count int64 = 0

	err := db.Preload("FilmCategory.Category").Find(&films).Count(&count).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		} else {
			panic(err)
		}
	}

	return &ListFilmResult{
		Count: count,
		Films: films,
	}
}
