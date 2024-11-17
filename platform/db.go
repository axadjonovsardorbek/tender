package platform

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/axadjonovsardorbek/tender/config"
	"github.com/axadjonovsardorbek/tender/internal/auth"
	"github.com/axadjonovsardorbek/tender/internal/tender"

	_ "github.com/lib/pq"
)

// Database wraps the SQL DB connection
type Storage struct {
	Db      *sql.DB
	TenderS tender.TenderI
	AuthS   auth.AuthI
}

// ConnectDatabase initializes the Postgres connection
func ConnectDatabase(cfg *config.Config) (*Storage, error) {
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

	tender := tender.NewTenderRepository(db)
	auth := auth.NewAuthRepo(db)

	log.Println("Connected to the Postgres database.")
	return &Storage{
		Db:      db,
		TenderS: tender,
		AuthS:   auth}, nil
}

// Close closes the database connection
func (db *Storage) Close() {
	if err := db.Db.Close(); err != nil {
		log.Printf("Error closing database: %v", err)
	}
}
