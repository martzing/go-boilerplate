package boot

import (
	"github.com/gofiber/fiber/v2"
	"github.com/martzing/go-boilerplate/middleware"
)

func bootHelmet(app *fiber.App) {
	app.Use(middleware.Helmet)
}
