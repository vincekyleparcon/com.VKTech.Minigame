package main

import (
	"encoding/json"
	"net/http"

	sdkconnmgr "github.com/flarehotspot/sdk/api/connmgr"
	sdkhttp "github.com/flarehotspot/sdk/api/http"
	sdkplugin "github.com/flarehotspot/sdk/api/plugin"
)

type MyData struct {
	Highscore int            `json:"highscore"`
	Records   map[string]int `json:"records"`
}

type MyScore struct {
	HighScore int `json:"highScore"`
	Score     int `json:"score"`
}

func main() {}

func Init(api sdkplugin.PluginApi) {
	// define the portal route
	portalRoute := sdkhttp.VuePortalRoute{
		RouteName: "portal.game",
		RoutePath: "/game",
		Component: "portal/Game.vue",
	}

	// register portal route
	api.Http().VueRouter().RegisterPortalRoutes(portalRoute)

	api.Http().VueRouter().PortalItemsFunc(func(clnt sdkconnmgr.ClientDevice) []sdkhttp.VuePortalItem {
		portalItem := sdkhttp.VuePortalItem{
			Label:     "Play Game",
			IconPath:  "icons/UFO.png",
			RouteName: "portal.game",
		}
		return []sdkhttp.VuePortalItem{portalItem}
	})

	Route := api.Http().HttpRouter().PluginRouter().Post("/score/save", func(w http.ResponseWriter, r *http.Request) {
		var data MyScore
		if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
			api.Http().VueResponse().Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		my_key := "my_score"

		// var oldData MyData
		// err := api.Config().Plugin(my_key).Get(&oldData)
		// if err != nil {
		// 	api.Http().VueResponse().Error(w, err.Error(), http.StatusInternalServerError)
		// 	return
		// }

		var oldData MyScore
		err := api.Config().Plugin(my_key).Get(&oldData)
		if err != nil {
			api.Http().VueResponse().Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if data.Score >= oldData.HighScore {
			data.HighScore = data.Score
		} else {
			data.HighScore = oldData.HighScore
		}

		// score, ok := oldData.Records[data.Name]
		// if ok {
		// 	oldData.Records[data.Name] = data.Score
		// 	err = api.Config().Plugin(my_key).Save(oldData)
		// }

		if err := api.Config().Plugin(my_key).Save(data); err != nil {
			api.Http().VueResponse().Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		api.Http().VueResponse().Json(w, nil, http.StatusOK)

	})

	RouteGet := api.Http().HttpRouter().PluginRouter().Get("/score/get", func(w http.ResponseWriter, r *http.Request) {

		my_key := "my_score"

		var oldData MyScore
		err := api.Config().Plugin(my_key).Get(&oldData)
		if err != nil {
			api.Http().VueResponse().Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		api.Http().VueResponse().Json(w, oldData, http.StatusOK)

	})

	Route.Name("score.save")
	RouteGet.Name("score.get")

}
