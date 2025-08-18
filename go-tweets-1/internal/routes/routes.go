package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/Devanshi-cloud/go-tweets/internal/handlers"
)

func RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/tweets", handlers.CreateTweet).Methods(http.MethodPost)
	r.HandleFunc("/tweets/{id:[0-9]+}", handlers.GetTweet).Methods(http.MethodGet)
	r.HandleFunc("/tweets", handlers.GetAllTweets).Methods(http.MethodGet)
	r.HandleFunc("/tweets/{id:[0-9]+}", handlers.UpdateTweet).Methods(http.MethodPut)
	r.HandleFunc("/tweets/{id:[0-9]+}", handlers.DeleteTweet).Methods(http.MethodDelete)
}