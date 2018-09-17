// Package app application container and routing config module
package app

import (
	"encoding/json"
	"net/http"

	"github.com/ilovelili/geo/server/util"
	"golang.org/x/net/context"
	"googlemaps.github.io/maps"
)

type georequest struct {
	Origin      string   `json:"origin"`
	Destination string   `json:"destination"`
	WayPoints   []string `json:"waypoints"`
}

// healthcheck health check
func (app *App) healthcheck(w http.ResponseWriter, r *http.Request) {
	util.RespondWithJSON(w, http.StatusOK, struct{ Healthy bool }{true})
}

// geo main geo api
func (app *App) geo(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var gr *georequest
	err := decoder.Decode(&gr)
	if err != nil {
		util.RespondWithError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	c, err := maps.NewClient(maps.WithAPIKey(app.Config.APIKey), maps.WithRateLimit(app.Config.RateLimit))
	if err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, "failed to init geo client")
		return
	}

	req := &maps.DirectionsRequest{
		Origin:      gr.Origin,
		Destination: gr.Destination,
		Waypoints:   gr.WayPoints,
		Mode:        maps.TravelModeDriving,
		Region:      "JP",
		Language:    "ja",
		Optimize:    true,
	}

	route, _, err := c.Directions(context.Background(), req)
	if err != nil {
		util.RespondWithError(w, http.StatusBadGateway, "cannot get direction info")
		return
	}

	util.RespondWithJSON(w, http.StatusOK, route)
}
