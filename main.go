//go:build !mono

package main

import (
	"encoding/json"
	"net/http"

	sdkconnmgr "github.com/flarehotspot/sdk/api/connmgr"
	sdkhttp "github.com/flarehotspot/sdk/api/http"
	sdkplugin "github.com/flarehotspot/sdk/api/plugin"
)

const (
	cfgKey = "game_record"
)

type Player struct {
	MacAddr string `json:"mac"`
	Name    string `json:"name"`
	Score   int    `json:"score"`
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
		var data Player
		if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
			api.Http().VueResponse().Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		clnt, err := api.Http().GetClientDevice(r)
		if err != nil {
			api.Http().VueResponse().Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		data.MacAddr = clnt.MacAddr()

		var records []Player

		err = api.Config().Custom(cfgKey).Get(&records)
		if err != nil {
			api.Http().VueResponse().Error(w, err.Error(), http.StatusInternalServerError)
			records = []Player{}
		}

		if PlayerExist(records, data.Name) {
			records = updatePlayerScore(records, data)
		} else {
			records = append(records, data)
		}

		// Sort the records highest to lowest score
		SortPlayers(records)

		if err := api.Config().Custom(cfgKey).Save(records); err != nil {
			api.Http().VueResponse().Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		api.Http().VueResponse().Json(w, nil, http.StatusOK)

	})

	RouteGet := api.Http().HttpRouter().PluginRouter().Get("/score/get", func(w http.ResponseWriter, r *http.Request) {

		var records []Player
		err := api.Config().Custom(cfgKey).Get(&records)
		if err != nil {
			records = []Player{}
		}

		SortPlayers(records)

		api.Http().VueResponse().Json(w, records, http.StatusOK)

	})

	MyScoreRoute := api.Http().HttpRouter().PluginRouter().Post("/score/getMyScore", func(w http.ResponseWriter, r *http.Request) {

		var data Player
		if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
			api.Http().VueResponse().Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var records []Player
		err := api.Config().Custom(cfgKey).Get(&records)
		if err != nil {
			records = []Player{}
		}

		myScore := getPlayerScore(records, data.Name)

		api.Http().VueResponse().Json(w, myScore, http.StatusOK)

	})

	Route.Name("score.save")
	RouteGet.Name("score.get")
	MyScoreRoute.Name("score.getMyScore")
}

// sort players from highest score to lowest
func SortPlayers(players []Player) []Player {
	for i := 0; i < len(players); i++ {
		for j := i + 1; j < len(players); j++ {
			if players[i].Score < players[j].Score {
				players[i], players[j] = players[j], players[i]
			}
		}
	}
	return players
}

func updatePlayerScore(players []Player, newData Player) []Player {
	for i, player := range players {
		if player.Name == newData.Name {
			if newData.Score >= player.Score {
				players[i].Score = newData.Score
			} else {
				return players
			}
		}
	}
	return players
}

func PlayerExist(players []Player, playerName string) bool {
	for _, player := range players {
		if player.Name == playerName {
			return true
		}
	}
	return false
}

func getPlayerScore(players []Player, name string) int {
	for _, player := range players {
		if player.Name == name {
			return player.Score
		}
	}
	return 0
}
