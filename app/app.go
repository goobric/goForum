// app/app.go

package app

import (
	"fmt"
	"net/http"
)

// App represents the web forum application.
type App struct {
	Router      *mux.Router
	Store       *sessions.CookieStore
	DatabaseURL string
}

// Initialize sets up and configures the application.
func (a *App) Initialize(databaseURL string, store *sessions.CookieStore) {
	a.DatabaseURL = databaseURL
	a.Store = store

	// Initialize the router
	a.Router = mux.NewRouter()

	// Set up routes and handlers
	a.initializeRoutes()
}

// Run starts the application's web server.
func (a *App) Run(addr string) {
	fmt.Printf("Server is listening on %s...\n", addr)
	err := http.ListenAndServe(addr, a.Router)
	if err != nil {
		fmt.Println(err)
	}
}