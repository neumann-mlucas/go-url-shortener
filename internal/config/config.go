package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "modernc.org/sqlite"
)

// Config holds the database configuration values
type Config struct {
	DB *sql.DB // Pointer to the SQLite in-memory database connection
}

// Global variable to hold the app configuration
var AppConfig *Config

// LoadConfig initializes the application configuration and opens the SQLite in-memory DB
func LoadConfig() error {
	AppConfig = &Config{}

	db, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		return fmt.Errorf("failed to open SQLite database: %w", err)
	}
	AppConfig.DB = db

	if err := createTables(db); err != nil {
		return fmt.Errorf("failed to create tables: %w", err)
	}
	return nil
}

// LoadTestConfig initializes the application configuration and opens the SQLite in-memory DB
func LoadTestConfig() error {
	AppConfig = &Config{}

	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		return fmt.Errorf("failed to open SQLite database: %w", err)
	}
	AppConfig.DB = db

	_, err = db.Exec("DROP TABLE IF EXISTS urls")
	if err != nil {
		return fmt.Errorf("failed to drop table: %w", err)
	}

	if err := createTables(db); err != nil {
		return fmt.Errorf("failed to create tables: %w", err)
	}
	return nil
}

// createTables creates necessary tables in the SQLite in-memory database
func createTables(db *sql.DB) error {
	query := `
    CREATE TABLE IF NOT EXISTS urls (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        hash TEXT NOT NULL,
        url TEXT UNIQUE NOT NULL,
		active BOOLEAN DEFAULT 1 NOT NULL
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
