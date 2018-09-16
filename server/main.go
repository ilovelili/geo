package main

import (
	"github.com/ilovelili/geo/server/app"
)

// entry
func main() {
	app := &app.App{}
	app.Initialize()
	app.Serve(":3200")
}
