package config

import (
	"database/sql"
	"fmt"
	"internal/model"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

// Config holds the database configuration values
type Config struct {
	DB *sql.DB // Pointer to the SQLite in-memory database connection
}

// Global variable to hold the app configuration
var AppConfig *Config

// LoadConfig initializes the application configuration and opens the SQLite in-memory DB
func LoadConfig() error {
	// Initialize the config struct
	AppConfig = &Config{}

	// Open an SQLite in-memory database
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		return fmt.Errorf("failed to open SQLite database: %w", err)
	}

	// Assign the *sql.DB instance to the AppConfig
	AppConfig.DB = db

	// Optionally, you can create tables here, for example:
	if err := createTables(db); err != nil {
		return fmt.Errorf("failed to create tables: %w", err)
	}

	return nil
}

// createTables creates necessary tables in the SQLite in-memory database
func createTables(db *sql.DB) error {
	// Example table creation (you can add more as per your requirements)
	query := `
    CREATE TABLE IF NOT EXISTS urls (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        hash TEXT NOT NULL,
        url TEXT NOT NULL
    );
    `
	_, err := db.Exec(query)
	if err != nil {
		return fmt.Errorf("failed to create table: %w", err)
	}
	return nil
}

// CloseDB closes the database connection when the application exits
func CloseDB() {
	if AppConfig.DB != nil {
		err := AppConfig.DB.Close()
		if err != nil {
			log.Printf("Error closing the database: %v", err)
		}
	}
}
