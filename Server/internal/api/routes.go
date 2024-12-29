package api

import (
	"database/sql"
	"rent-server/config"
	"rent-server/internal/api/handlers"
	"rent-server/internal/api/middleware"
	"rent-server/internal/repositories"
	"rent-server/internal/services"

	"github.com/gorilla/mux"
)

func SetupRoutes(db *sql.DB, cfg *config.Config) *mux.Router {

	r := mux.NewRouter()
	apiRouter := r.PathPrefix("/api/v1").Subrouter()
	
	userRepo := repositories.NewUserRepository(db)

	authService := services.NewAuthService(userRepo, cfg)
	authHandler := handlers.NewAuthHandler(authService)
	apiRouter.HandleFunc("/auth/login", authHandler.LoginHandler).Methods("POST")
	apiRouter.HandleFunc("/auth/signup", authHandler.SignupHandler).Methods("POST")

	
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)
	detailsRouter := apiRouter.PathPrefix("/details").Subrouter()
	detailsRouter.Use(middleware.AuthMiddleware(cfg))
	detailsRouter.HandleFunc("/me", userHandler.GetUser).Methods("GET")

	return r
}
