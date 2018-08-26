package main

import (
	"log"

	"github.com/kr/pretty"
	"golang.org/x/net/context"
	"googlemaps.github.io/maps"
)

var (
	public = []maps.TransitMode{maps.TransitModeBus, maps.TransitModeRail, maps.TransitModeSubway}
)

func main() {
	config := GetConfig()
	c, err := maps.NewClient(maps.WithAPIKey(config.APIKey), maps.WithRateLimit(config.RateLimit))
	if err != nil {
		log.Fatalf("fatal error: %s", err)
	}
	r := &maps.DirectionsRequest{
		Origin:      "New York",
		Destination: "Boston",
		Mode:        maps.TravelModeTransit,
		TransitMode: public,
	}
	route, _, err := c.Directions(context.Background(), r)
	if err != nil {
		log.Fatalf("fatal error: %s", err)
	}

	pretty.Println(route)
}
