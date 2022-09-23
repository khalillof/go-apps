package main

import (
	"log"
	"net/http"
	routes "github.com/khalillof/go-apps/jwt-app/routes"
)

func main() {
	// "Signin" and "Welcome" are the handlers that we will implement
	http.HandleFunc("/signin", routes.Signin)
	http.HandleFunc("/welcome",routes.Welcome)
	http.HandleFunc("/refresh", routes.Refresh)

	// start the server on port 8000
	log.Fatal(http.ListenAndServe(":8000", nil))
}