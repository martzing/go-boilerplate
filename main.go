package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/martzing/go-boilerplate/boot"
	"github.com/martzing/go-boilerplate/configs"
)

func main() {
	app := fiber.New()

	boot.NewBoot(app)

	err := app.Listen(fmt.Sprintf(":%d", *configs.Port))

	if err != nil {
		panic(err)
	}
}
