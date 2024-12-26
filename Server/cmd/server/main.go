package main

import (
	"errors"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"

	"rent-server/config"
	"rent-server/initializers"
	"rent-server/internal/api"
)

func run() error {
	var err error
	cfg := config.Load()

	if cfg.JwtKey == "" {
		return errors.New("please set environment variables for jwt key")
	}

	db, err := initializers.GetNewDBConnection(cfg)
	if err != nil {
		return err
	}

	r := api.SetupRoutes(db, cfg)

	log.Printf("Starting server on port %s", cfg.Port)
	err = http.ListenAndServe(":"+cfg.Port, r)
	if err != nil {
		return err
	}

	return nil
}

func main() {

	err := run()
	if err != nil {
		log.Fatal("Server error:", err.Error())
		os.Exit(1)
	}

}
