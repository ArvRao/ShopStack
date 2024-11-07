package database

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
)

var DB *pgxpool.Pool

func InitDB() error {
	// Read database configuration from environment variables
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	// Construct the database connection string
	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", dbUser, dbPassword, dbHost, dbPort, dbName)

	// Create a connection pool
	config, err := pgxpool.ParseConfig(dbURL)
	if err != nil {
		return fmt.Errorf("error parsing database config: %v", err)
	}

	// Set max connections in the pool
	config.MaxConns = 10

	// Create the connection pool
	pool, err := pgxpool.ConnectConfig(context.Background(), config)
	if err != nil {
		return fmt.Errorf("error connecting to the database: %v", err)
	}

	// Verify the connection
	if err := pool.Ping(context.Background()); err != nil {
		return fmt.Errorf("error pinging database: %v", err)
	}

	DB = pool
	fmt.Println("Successfully connected to the database")
	return nil
}

func CloseDB() {
	if DB != nil {
		DB.Close()
	}
}
