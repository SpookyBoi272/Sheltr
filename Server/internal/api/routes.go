package api

import (
	"database/sql"
	"rent-server/config"
	"rent-server/internal/api/handlers"
	"rent-server/internal/repositories"
	"rent-server/internal/services"

	"github.com/gorilla/mux"
)

func SetupRoutes(db *sql.DB, cfg *config.Config) *mux.Router {

	r := mux.NewRouter()
	apiRouter := r.PathPrefix("/api/v1").Subrouter()

	userRepo := repositories.NewUserRepository(db)

	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	authService := services.NewAuthService(userRepo, cfg)
	authHandler := handlers.NewAuthHandler(authService)

	apiRouter.HandleFunc("/auth/login", authHandler.LoginHandler).Methods("POST")
	apiRouter.HandleFunc("/auth/signup", authHandler.SignupHandler).Methods("POST")

	apiRouter.HandleFunc("/user/{id:[0-9]+}", userHandler.GetUser).Methods("GET")
	apiRouter.HandleFunc("/user", userHandler.GetAllUsers).Methods("GET")

	return r
}
