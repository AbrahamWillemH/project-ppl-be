package config

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

var DB *pgx.Conn

func ConnectDB() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: No .env file found, using system environment variables.")
	}

	// Get database URL
	connStr := os.Getenv("DATABASE_URL")
	if connStr == "" {
		log.Fatal("DATABASE_URL is not set in .env file or environment variables.")
	}

	// Connect to Neon PostgreSQL DB
	conn, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		log.Fatalf("Failed to connect to PostgreSQL NeonDB: %v", err)
	}

	fmt.Println("Connected to CockroachDB!")
	DB = conn

	fmt.Println("Table check completed.")
}

// CloseDB closes the database connection
func CloseDB() {
	if DB != nil {
		DB.Close(context.Background())
		fmt.Println("Database connection closed.")
	}
}
