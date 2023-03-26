package route

import (
	"github.com/martzing/go-boilerplate/controller/mongo"
)

var mongoRoutes = []Route{
	{
		Method: "GET",
		Path:   "/customer",
		Configs: RouteConfig{
			Limiter: &Limiter{
				Max: 4,
				Exp: 1,
			},
			Auth: Auth{
				Name: "chomchob-frontend",
				Scope: []string{
					"boilerplate.mongo.read",
				},
			},
		},
		Handler: mongo.ListCustomerAdapter,
	},
	{
		Method: "POST",
		Path:   "/customer",
		Configs: RouteConfig{
			Limiter: &Limiter{
				Max: 4,
				Exp: 1,
			},
			Auth: Auth{
				Name: "chomchob-frontend",
				Scope: []string{
					"boilerplate.mongo.create",
				},
			},
		},
		Handler: mongo.CreateCustomerAdapter,
	},
	{
		Method: "PATCH",
		Path:   "/customer",
		Configs: RouteConfig{
			Limiter: &Limiter{
				Max: 4,
				Exp: 1,
			},
			Auth: Auth{
				Name: "chomchob-frontend",
				Scope: []string{
					"boilerplate.mongo.update",
				},
			},
		},
		Handler: mongo.UpdateCustomerAdapter,
	},
	{
		Method: "DELETE",
		Path:   "/customer",
		Configs: RouteConfig{
			Limiter: &Limiter{
				Max: 4,
				Exp: 1,
			},
			Auth: Auth{
				Name: "chomchob-frontend",
				Scope: []string{
					"boilerplate.mongo.delete",
				},
			},
		},
		Handler: mongo.DeleteCustomerAdapter,
	},
}
