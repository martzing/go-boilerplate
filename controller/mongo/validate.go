package mongo

import (
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	mongoType "github.com/martzing/go-boilerplate/controller/mongo/type"
	utilityHelper "github.com/martzing/go-boilerplate/helpers/utility"
)

func listCustomerValidate(c *fiber.Ctx) *mongoType.ListCustomerParams {
	validate := validator.New()
	validateStruct := new(mongoType.ListCustomerValidate)

	if err := c.QueryParser(validateStruct); err != nil {
		panic(err)
	}

	err := validate.Struct(validateStruct)

	if err != nil {
		panic(utilityHelper.ValidationError(err.(validator.ValidationErrors)))
	}

	data := mongoType.ListCustomerParams{
		Search: validateStruct.Search,
		Offset: utilityHelper.StrToInt64(validateStruct.Offset, 0),
		Limit:  utilityHelper.StrToInt64(validateStruct.Limit, -1),
	}

	return &data
}

func createCustomerValidate(c *fiber.Ctx) *mongoType.CreateCustomerParams {
	validate := validator.New()
	validateStruct := new(mongoType.CreateCustomerValidate)

	if err := c.BodyParser(validateStruct); err != nil {
		panic(err)
	}

	err := validate.Struct(validateStruct)

	if err != nil {
		panic(utilityHelper.ValidationError(err.(validator.ValidationErrors)))
	}

	data := mongoType.CreateCustomerParams{
		StoreId:    validateStruct.StoreId,
		FirstName:  validateStruct.FirstName,
		LastName:   validateStruct.LastName,
		Email:      validateStruct.Email,
		AddressId:  validateStruct.AddressId,
		Activebool: validateStruct.Activebool,
	}

	return &data
}

func updateCustomerValidate(c *fiber.Ctx) *mongoType.UpdateCustomerParams {
	validate := validator.New()
	validateStruct := new(mongoType.UpdateCustomerValidate)

	if err := c.BodyParser(validateStruct); err != nil {
		panic(err)
	}

	err := validate.Struct(validateStruct)

	if err != nil {
		panic(utilityHelper.ValidationError(err.(validator.ValidationErrors)))
	}

	data := mongoType.UpdateCustomerParams{
		ID:         validateStruct.ID,
		StoreId:    validateStruct.StoreId,
		FirstName:  validateStruct.FirstName,
		LastName:   validateStruct.LastName,
		Email:      validateStruct.Email,
		AddressId:  validateStruct.AddressId,
		Activebool: validateStruct.Activebool,
	}

	return &data
}

func deleteCustomerValidate(c *fiber.Ctx) *mongoType.DeleteCustomerParams {
	validate := validator.New()
	validateStruct := new(mongoType.DeleteCustomerValidate)

	if err := c.BodyParser(validateStruct); err != nil {
		panic(err)
	}

	err := validate.Struct(validateStruct)

	if err != nil {
		panic(utilityHelper.ValidationError(err.(validator.ValidationErrors)))
	}

	data := mongoType.DeleteCustomerParams{
		ID: validateStruct.ID,
	}

	return &data
}
