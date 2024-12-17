package api

import (
	"rent-server/internal/api/handlers"

	"github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router{

	r := mux.NewRouter()
	apiRouter := r.PathPrefix("/api/v1").Subrouter()


	apiRouter.HandleFunc("/user", handlers.UserHandler).Methods("GET")

	return r
}

