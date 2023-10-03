// app/main.go

package main

import (
	"html/template"
	"net/http"
	// "github.com/gorilla/mux"
	// "github.com/gorilla/sessions"
	// "app/sessions"
	// "app/handlers"
)

var store = sessions.NewCookieStore([]byte("session-key"))

func main() {
	router := mux.NewRouter()

	// Register Handlers
	router.HandleFunc("/", handlers.IndexHandler)
	router.HandleFunc("/login", handlers.LoginHandler).Methods("GET", "POST")
	router.HandleFunc("/dashboard", handlers.DashboardHandler)
	router.HandleFunc("/logout", handlers.LogoutHandler)

	//Serve static files - css, js, images
	router.PathPrefix("/static").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	
	//Templates
	templates := template.Must(template.ParseGlob("templates/*.html"))

	//Set up the session store
	sessionManager := session.newSessionManager(store)

	http.Handle("/", router)

	http.ListenAndServe(":8080", nil)
}