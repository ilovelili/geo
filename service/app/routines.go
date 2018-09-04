// Package app application container and routing config module
package app

import (
	"net/http"

	"github.com/ilovelili/geo/service/util"
	"golang.org/x/net/context"
	"googlemaps.github.io/maps"
)

// healthcheck health check
func (app *App) healthcheck(w http.ResponseWriter, r *http.Request) {
	util.RespondWithJSON(w, http.StatusOK, struct{ Healthy bool }{true})
}

func (app *App) geo(w http.ResponseWriter, r *http.Request) {
	c, err := maps.NewClient(maps.WithAPIKey(app.Config.APIKey), maps.WithRateLimit(app.Config.RateLimit))
	if err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, err.Error())
	}

	req := &maps.DirectionsRequest{
		Origin:       "Sangenjaya station, Tokyo",
		Destination:  "Shibuya station, Tokyo",
		TrafficModel: maps.TrafficModelOptimistic,
		Mode:         maps.TravelModeDriving,
		Region:       "JP",
		Language:     "ja",
		Optimize:     true,
	}

	route, _, err := c.Directions(context.Background(), req)
	if err != nil {
		util.RespondWithError(w, http.StatusBadGateway, err.Error())
	}

	util.RespondWithJSON(w, http.StatusOK, route)
}
