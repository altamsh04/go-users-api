package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var DB *sql.DB

func Connect() error {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found, reading environment variables")
	}

	dsn := os.Getenv("DB_DSN")
	if dsn == "" {
		return fmt.Errorf("DB_DSN environment variable is not set")
	}

	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		return fmt.Errorf("failed to open DB: %w", err)
	}

	if err = DB.Ping(); err != nil {
		return fmt.Errorf("failed to connect to DB: %w", err)
	}

	fmt.Println("MySQL database connected.")
	return nil
}
