package route

import (
	"github.com/gofiber/fiber/v2"
)

type Adapter func(*fiber.Ctx) error

type Limiter struct {
	Max int
	Exp float64
}

type Auth struct {
	Name  string
	Scope []string
}

type RouteConfig struct {
	*Limiter
	Auth
}
type Route struct {
	Method  string
	Path    string
	Configs RouteConfig
	Handler Adapter
}
