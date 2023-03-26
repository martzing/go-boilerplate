package middleware

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/storage/redis"
	"github.com/martzing/go-boilerplate/configs"
	utilityHelper "github.com/martzing/go-boilerplate/helpers/utility"
)

func Limiter(max int, exp float64) func(*fiber.Ctx) error {
	redisConfig := configs.Redis
	store := redis.New(redis.Config{
		Host:     redisConfig.Host,
		Port:     redisConfig.Port,
		Database: 1,
	})

	return limiter.New(limiter.Config{
		Max:        max,
		Expiration: time.Duration(exp) * time.Second,
		KeyGenerator: func(c *fiber.Ctx) string {
			projectName := *configs.ServiceName
			key := fmt.Sprintf("%s:%s:%s:%s", c.Get("Device-Id"), projectName, c.Path(), c.Method())

			return key
		},
		LimitReached: func(c *fiber.Ctx) error {
			err := utilityHelper.CustomError{
				StatusCode: fiber.StatusTooManyRequests,
				Code:       "too-many-requests",
				Message:    "share.error.too many requests",
			}

			return err.SendError(c)
		},
		Storage: store,
	})
}
