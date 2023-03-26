package boot

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	adapterHelper "github.com/martzing/go-boilerplate/helpers/adapter"
	"github.com/martzing/go-boilerplate/middleware"
	"github.com/martzing/go-boilerplate/route"
)

func generateRoute(versionRouter fiber.Router, routes []route.Route, rootPath string) {
	router := versionRouter.Group(rootPath)

	for _, route := range routes {
		routeConfig := route.Configs
		handler := adapterHelper.CreateHandlerAdapter(route.Handler)
		handlers := []func(*fiber.Ctx) error{}

		if routeConfig.Limiter != nil {
			limiter := routeConfig.Limiter
			handlers = append(handlers, middleware.Limiter(limiter.Max, limiter.Exp))
		}

		handlers = append(handlers, handler)

		router.Add(route.Method, route.Path, handlers...)
	}
}

func bootRoute(app *fiber.App) {
	baseRoutes, versionRoutes := route.GatherRoute()

	for rootPath, routes := range baseRoutes {
		generateRoute(app, routes, rootPath)
	}

	for i, versionRoute := range versionRoutes {
		versionRouter := app.Group(fmt.Sprintf("/v%d", (i + 1)))

		for rootPath, routes := range versionRoute {
			generateRoute(versionRouter, routes, rootPath)
		}
	}
}
