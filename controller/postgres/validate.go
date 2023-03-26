package postgres

import (
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	postgresType "github.com/martzing/go-boilerplate/controller/postgres/type"
	utilityHelper "github.com/martzing/go-boilerplate/helpers/utility"
)

func listCustomerValidate(c *fiber.Ctx) *postgresType.ListCustomerParams {
	validate := validator.New()
	validateStruct := new(postgresType.ListCustomerValidate)

	if err := c.QueryParser(validateStruct); err != nil {
		panic(err)
	}

	err := validate.Struct(validateStruct)

	if err != nil {
		panic(utilityHelper.ValidationError(err.(validator.ValidationErrors)))
	}

	data := postgresType.ListCustomerParams{
		Search: validateStruct.Search,
		Offset: utilityHelper.StrToInt(validateStruct.Offset, 0),
		Limit:  utilityHelper.StrToInt(validateStruct.Limit, -1),
	}

	return &data
}

func listFilmValidate(c *fiber.Ctx) *postgresType.ListFilmParams {
	validate := validator.New()
	validateStruct := new(postgresType.ListFilmValidate)

	if err := c.QueryParser(validateStruct); err != nil {
		panic(err)
	}

	err := validate.Struct(validateStruct)

	if err != nil {
		panic(utilityHelper.ValidationError(err.(validator.ValidationErrors)))
	}

	data := postgresType.ListFilmParams{
		Search: validateStruct.Search,
		Offset: utilityHelper.StrToInt(validateStruct.Offset, 0),
		Limit:  utilityHelper.StrToInt(validateStruct.Limit, -1),
	}

	return &data
}

func createCustomerValidate(c *fiber.Ctx) *postgresType.CreateCustomerParams {
	validate := validator.New()
	validateStruct := new(postgresType.CreateCustomerValidate)

	if err := c.BodyParser(validateStruct); err != nil {
		panic(err)
	}

	err := validate.Struct(validateStruct)

	if err != nil {
		panic(utilityHelper.ValidationError(err.(validator.ValidationErrors)))
	}

	data := postgresType.CreateCustomerParams{
		StoreId:    validateStruct.StoreId,
		FirstName:  validateStruct.FirstName,
		LastName:   validateStruct.LastName,
		Email:      validateStruct.Email,
		AddressId:  validateStruct.AddressId,
		Activebool: validateStruct.Activebool,
	}

	return &data
}

func updateCustomerValidate(c *fiber.Ctx) *postgresType.UpdateCustomerParams {
	validate := validator.New()
	validateStruct := new(postgresType.UpdateCustomerValidate)

	if err := c.BodyParser(validateStruct); err != nil {
		panic(err)
	}

	err := validate.Struct(validateStruct)

	if err != nil {
		panic(utilityHelper.ValidationError(err.(validator.ValidationErrors)))
	}

	data := postgresType.UpdateCustomerParams{
		CustomerId: validateStruct.CustomerId,
		StoreId:    validateStruct.StoreId,
		FirstName:  validateStruct.FirstName,
		LastName:   validateStruct.LastName,
		Email:      validateStruct.Email,
		AddressId:  validateStruct.AddressId,
		Activebool: validateStruct.Activebool,
	}

	return &data
}

func deleteCustomerValidate(c *fiber.Ctx) *postgresType.DeleteCustomerParams {
	validate := validator.New()
	validateStruct := new(postgresType.DeleteCustomerValidate)

	if err := c.BodyParser(validateStruct); err != nil {
		panic(err)
	}

	err := validate.Struct(validateStruct)

	if err != nil {
		panic(utilityHelper.ValidationError(err.(validator.ValidationErrors)))
	}

	data := postgresType.DeleteCustomerParams{
		CustomerId: validateStruct.CustomerId,
	}

	return &data
}
