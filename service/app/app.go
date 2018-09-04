// Package app application container and routing config module
package app

import (
	"log"
	"net/http"

	"github.com/ilovelili/geo/service/config"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

// App They call me God object. So I think I am very cool
type App struct {
	Router         *mux.Router
	Config         *config.Config
	Request        *http.Request
	ResponseWriter http.ResponseWriter
}

// Serve serve Non-TLS with cros origin support
func (app *App) Serve(addr string) {
	handler := cors.Default().Handler(app.Router)
	log.Fatal(http.ListenAndServe(addr, handler))
}

// Initialize init the app
func (app *App) Initialize() {
	// set up config
	config := config.GetConfig()
	app.Config = config

	// set up new router
	app.Router = mux.NewRouter()

	// init routes
	app.InitializeRoutes()
}
