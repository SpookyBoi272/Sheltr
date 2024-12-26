package initializers

import (
	"database/sql"
	"fmt"
	"log"
	"rent-server/config"
)

func GetNewDBConnection(cfg *config.Config) (*sql.DB, error) {

	connString := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=disable", cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)
	db, err := sql.Open("postgres", connString)

	if err != nil {
		log.Fatal("Error connecting to the database:", err)
		return nil, err
	}

	return db, nil
}