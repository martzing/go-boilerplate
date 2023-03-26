package postgres

import (
	"fmt"

	"github.com/martzing/go-boilerplate/controller/postgres/module/db/models"
	postgresType "github.com/martzing/go-boilerplate/controller/postgres/type"
)

func listCustomerFormat(count int64, customers *[]models.Customer) postgresType.ListCustomerResponse {
	result := []postgresType.CustomerType{}

	for _, customer := range *customers {
		result = append(result, postgresType.CustomerType{
			ID:         customer.CustomerId,
			FirstName:  customer.FirstName,
			LastName:   customer.LastName,
			FullName:   fmt.Sprintf("%s %s", customer.FirstName, customer.LastName),
			Email:      customer.Email,
			CreateDate: customer.CreateDate,
			UpdatedAt:  customer.UpdatedAt,
		})
	}

	return postgresType.ListCustomerResponse{
		Count:     count,
		Customers: result,
	}
}

func listFilmFormat(count int64, films *[]models.Film) postgresType.ListFilmResponse {

	return postgresType.ListFilmResponse{
		Count: count,
		Films: *films,
	}
}
