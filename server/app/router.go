// Package app application container and routing config module
package app

// InitializeRoutes init routes
func (app *App) InitializeRoutes() {
	app.Router.HandleFunc("/", app.healthcheck).Methods("GET")
	app.Router.HandleFunc("/geo", app.geo).Methods("POST")
}
