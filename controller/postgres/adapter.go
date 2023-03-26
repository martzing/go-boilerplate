package postgres

import (
	"github.com/gofiber/fiber/v2"
	"github.com/martzing/go-boilerplate/job"
	"github.com/martzing/go-boilerplate/queue"
)

func GetCustomerAdapter(c *fiber.Ctx) error {
	params := listCustomerValidate(c)
	data := listCustomer(params)

	return c.JSON(data)
}

func ListFilmAdapter(c *fiber.Ctx) error {
	params := listFilmValidate(c)
	data := listFilm(params)

	return c.JSON(data)
}

func CreateCustomerAdapter(c *fiber.Ctx) error {
	params := createCustomerValidate(c)
	data := createCustomer(params)

	return c.JSON(data)
}

func UpdateCustomerAdapter(c *fiber.Ctx) error {
	params := updateCustomerValidate(c)
	data := updateCustomer(params)

	return c.JSON(data)
}

func DeleteCustomerAdapter(c *fiber.Ctx) error {
	params := deleteCustomerValidate(c)
	data := deleteCustomer(params)

	return c.JSON(data)
}

var CreateCustomerQueueAdapter = queue.CreateQueueAdapter(createCustomerQueue)

var CreateCustomerJobAdapter = job.CreateJobAdapter(createCustomerJob)
