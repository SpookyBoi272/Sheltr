package initializers

import (
	"database/sql"
	"fmt"
	"log"
	"rent-server/config"
)

func GetNewDBConnection(cfg *config.Config) (*sql.DB, error) {

	connString := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=disable", cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)
	log.Print("Connecting to the database...")

	db, err := sql.Open("postgres", connString)
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
		return nil, err
	}
	
	err = db.Ping()
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
		return nil, err
	}

	log.Print("Successfully connected to database server")
	return db, nil
}