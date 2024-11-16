package platform

import (
	"database/sql"
	"fmt"
	"log"
	"github.com/axadjonovsardorbek/tender/config"

	_ "github.com/lib/pq"
)

// Database wraps the SQL DB connection
type Database struct {
	Client *sql.DB
}

// ConnectDatabase initializes the Postgres connection
func ConnectDatabase(cfg *config.Config) (*Database, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName,
	)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	// Verify connection
	if err = db.Ping(); err != nil {
		return nil, err
	}

	log.Println("Connected to the Postgres database.")
	return &Database{Client: db}, nil
}

// Close closes the database connection
func (db *Database) Close() {
	if err := db.Client.Close(); err != nil {
		log.Printf("Error closing database: %v", err)
	}
}
