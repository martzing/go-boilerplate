package route

import "github.com/martzing/go-boilerplate/controller/postgres"

var postresRoutes = []Route{
	{
		Method: "GET",
		Path:   "/customer",
		Configs: RouteConfig{
			Limiter: &Limiter{
				Max: 4,
				Exp: 1,
			},
		},
		Handler: postgres.GetCustomerAdapter,
	},
	{
		Method: "POST",
		Path:   "/customer",
		Configs: RouteConfig{
			Limiter: &Limiter{
				Max: 4,
				Exp: 1,
			},
		},
		Handler: postgres.CreateCustomerAdapter,
	},
	{
		Method: "PATCH",
		Path:   "/customer",
		Configs: RouteConfig{
			Limiter: &Limiter{
				Max: 4,
				Exp: 1,
			},
		},
		Handler: postgres.UpdateCustomerAdapter,
	},
	{
		Method: "DELETE",
		Path:   "/customer",
		Configs: RouteConfig{
			Limiter: &Limiter{
				Max: 4,
				Exp: 1,
			},
		},
		Handler: postgres.DeleteCustomerAdapter,
	},
	{
		Method: "GET",
		Path:   "/film",
		Configs: RouteConfig{
			Limiter: &Limiter{
				Max: 4,
				Exp: 1,
			},
		},
		Handler: postgres.ListFilmAdapter,
	},
}
