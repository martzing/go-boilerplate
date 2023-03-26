package boot

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func bootLogger(app *fiber.App) {
	app.Use(logger.New(logger.Config{
		Next: func(c *fiber.Ctx) bool {
			path := c.Path()

			if path == "/health/live" || path == "/health/ready" {
				return true
			} else {
				return false
			}
		},
		Format:     "[${ip} - ${time}] - ${status} ${protocol} ${method} ${host}${path} [${ua} - ${latency}]\n",
		TimeFormat: "2006-01-02T15:04:05",
		TimeZone:   "Asia/Bangkok",
	}))
}
