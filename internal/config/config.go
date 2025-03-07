package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

// Config holds the database configuration values
type Config struct {
	Port   string
	Driver string
	URI    string
	DB     *sql.DB
}

// Global variable to hold the app configuration
var AppConfig *Config

// LoadConfig initializes the application configuration and opens the SQLite in-memory DB
func LoadConfig() error {
	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = ":8080"
	}

	driver, ok := os.LookupEnv("DRIVER")
	if !ok {
		driver = "sqlite3"
	}

	uri, ok := os.LookupEnv("DATABASE_URI")
	if !ok {
		uri = ":memory:"
	}

	db, err := sql.Open(driver, uri)
	if err != nil {
		return fmt.Errorf("failed to open SQLite database: %w", err)
	}

	AppConfig = &Config{Port: port, Driver: driver, URI: uri, DB: db}

	if err := createTables(db); err != nil {
		return fmt.Errorf("failed to create tables: %w", err)
	}
	return nil
}

// LoadTestConfig initializes the application configuration and opens the SQLite in-memory DB
func LoadTestConfig() error {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		return fmt.Errorf("failed to open SQLite database: %w", err)
	}

	_, err = db.Exec("DROP TABLE IF EXISTS urls")
	if err != nil {
		return fmt.Errorf("failed to drop table: %w", err)
	}

	AppConfig = &Config{Port: ":8080", Driver: "sqlite3", URI: ":memory:", DB: db}

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
