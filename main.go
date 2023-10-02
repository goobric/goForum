// main.go

package main

import (
	"fmt"
	//import external "github.com/gorilla/sessions"
	"net/http"
)

func main() {
	// Initialize the application
	app := &App{}
	store := sessions.NewCookieStore([]byte("secret-key"))
	app.Initialize("database-url", store)

	// Start the web server
	addr := ":8080"
	fmt.Printf("Server is listening on %s...\n", addr)
	err := http.ListenAndServe(addr, app.Router)
	if err != nil {
		fmt.Println(err)
	}
}