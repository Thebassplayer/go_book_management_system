package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Thebassplayer/go_book_management_system/pkg/middleware"
	"github.com/Thebassplayer/go_book_management_system/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.Use(middleware.LoggingMiddleware)
	// Print the path hit
	r.HandleFunc("/",
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Welcome to the Go Book Management System!")
		},
	).Methods("GET")
	// Import the RegisterBookStoreRoutes function from the routes package and call it with the router object
	routes.RegisterBookStoreRoutes(r)

	fmt.Printf("Starting server at port 9010\n")

	log.Fatal(http.ListenAndServe(":8080", r))
}
