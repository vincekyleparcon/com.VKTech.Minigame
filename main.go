package main

import (
	"net/http"

	sdkhttp "github.com/flarehotspot/core/sdk/api/http"
	sdkplugin "github.com/flarehotspot/core/sdk/api/plugin"
)

func main() {}

func Init(api sdkplugin.PluginApi) {
	// define the portal route
	portalRoute := sdkhttp.VuePortalRoute{
		RouteName: "portal.welcome",
		RoutePath: "/welcome/:name",
		Component: "portal/Welcome.vue",
		HandlerFunc: func(w http.ResponseWriter, r *http.Request) {
			params := api.Http().MuxVars(r)
			name := params["name"]

			data := map[string]interface{}{
				"name": name,
			}

			api.Http().VueResponse().Json(w, data, 200)
		},
		Middlewares: []func(http.Handler) http.Handler{},
	}
	// register portal route
	api.Http().VueRouter().RegisterPortalRoutes(portalRoute)
}
