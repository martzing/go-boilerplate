package mongo

import (
	"github.com/martzing/go-boilerplate/controller/mongo/module/db/models"
	mongoType "github.com/martzing/go-boilerplate/controller/mongo/type"
)

func getCustomerFormat(id string, customer models.Customer) mongoType.CustomerType {
	result := mongoType.CustomerType{
		ID:         id,
		StoreId:    customer.StoreId,
		FirstName:  customer.FirstName,
		LastName:   customer.LastName,
		Email:      customer.Email,
		AddressId:  customer.AddressId,
		Activebool: customer.Activebool,
		CreatedAt:  customer.CreatedAt,
		UpdatedAt:  customer.UpdatedAt,
	}

	return result
}
