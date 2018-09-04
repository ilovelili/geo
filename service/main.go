package main

import (
	"github.com/ilovelili/geo/service/app"
)

// entry
func main() {
	app := &app.App{}
	app.Initialize()
	app.Serve(":8080")
}
