package route

import (
	"github.com/samber/lo"
)

func GatherRoute() (map[string][]Route, []map[string][]Route) {
	baseRoutes := lo.Assign(
		map[string][]Route{"/health": healthRoutes},
	)

	v1 := lo.Assign(
		map[string][]Route{"/mongo": mongoRoutes},
		map[string][]Route{"/postgres": postresRoutes},
	)

	versionRoutes := []map[string][]Route{v1}

	return baseRoutes, versionRoutes
}
