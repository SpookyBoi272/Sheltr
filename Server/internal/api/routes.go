package api

import (
	"rent-server/internal/api/handlers"

	"github.com/gorilla/mux"
)

func SetupRoutes(userHandler *handlers.UserHandler) *mux.Router{

	r := mux.NewRouter()
	apiRouter := r.PathPrefix("/api/v1").Subrouter()
	
	apiRouter.HandleFunc("/user/{id:[0-9]+}", userHandler.GetUser).Methods("GET")
	apiRouter.HandleFunc("/user", userHandler.GetAllUsers).Methods("GET")

	return r
}


