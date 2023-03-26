package boot

import (
	"github.com/gofiber/fiber/v2"
	"github.com/martzing/go-boilerplate/configs"
	redisHelper "github.com/martzing/go-boilerplate/helpers/redis"
	"github.com/martzing/go-boilerplate/locales"
)

func NewBoot(app *fiber.App) {
	configs.BootConfig()
	locales.BootLocalize()

	redisHelper.New()

	bootCors(app)
	bootHelmet(app)
	bootLogger(app)
	bootRoute(app)

	if *configs.EnableQueue {
		bootQueue()
	}

	if *configs.EnableJob {
		bootJob()
	}
}
