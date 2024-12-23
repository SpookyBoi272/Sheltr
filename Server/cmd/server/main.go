package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"

	"rent-server/config"
	"rent-server/internal/api"
	"rent-server/internal/api/handlers"
	"rent-server/internal/repositories"
	"rent-server/internal/services"
)

func run() {
	cfg := config.Load()

	connString := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=disable", cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)
	db , err := sql.Open("postgres", connString)

	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}

	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)


	r := api.SetupRoutes(userHandler)

	log.Printf("Starting server on port %s", cfg.Port)
	error := http.ListenAndServe(":"+cfg.Port, r)
	if error != nil {
		log.Fatal("Server error:" ,err)
	}
}

func main() {

	run()
	
}