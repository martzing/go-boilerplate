package adapterHelper

import (
	"github.com/gofiber/fiber/v2"
	utilityHelper "github.com/martzing/go-boilerplate/helpers/utility"
)

func CreateHandlerAdapter(controller func(*fiber.Ctx) error) func(*fiber.Ctx) error {
	adapter := func(c *fiber.Ctx) error {
		defer func() error {
			if err := recover(); err != nil {
				switch e := err.(type) {
				case *utilityHelper.CustomError:
					return e.SendError(c)
				default:
					logger := utilityHelper.NewLogger()
					logger.Error(err)

					return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
						"code":    "unknown",
						"message": "Something went wrong, Please contact support.",
					})
				}
			}

			return nil
		}()

		err := controller(c)

		return err
	}

	return adapter
}
