package boot

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func bootCors(app *fiber.App) {
	app.Use(cors.New())
}
