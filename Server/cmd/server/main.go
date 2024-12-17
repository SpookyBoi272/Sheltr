package main

import (
	"log"
	"net/http"

	"rent-server/config"
	"rent-server/internal/api"
)

func main() {

	cfg := config.Load()

	r := api.SetupRoutes()

	log.Printf("Starting server on port %s", cfg.Port)
	err := http.ListenAndServe(":"+cfg.Port, r)
	if err != nil {
		log.Fatal("Server error:" ,err)
	}
}