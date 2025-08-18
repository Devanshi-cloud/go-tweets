package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/Devanshi-cloud/go-tweets/internal/routes"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterRoutes(r)

	http.Handle("/", r)

	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Could not start server: %s\n", err)
	}
}