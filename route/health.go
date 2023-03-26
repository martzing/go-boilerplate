package route

import (
	"github.com/gofiber/fiber/v2"
)

var healthRoutes = []Route{
	{
		Method: "GET",
		Path:   "/ready",
		Handler: func(c *fiber.Ctx) error {
			// TODO: Call data base service
			return c.Status(fiber.StatusOK).SendString("OK")
		},
	},
	{
		Method: "GET",
		Path:   "/live",
		Handler: func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusOK).SendString("OK")
		},
	},
}
