package database

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"os"
)

var DB *pgxpool.Pool

func Connect() error {
	_ = godotenv.Load()
	databaseUrl := os.Getenv("DATABASE_URL")
	if databaseUrl == "" {
		return fmt.Errorf("Path variable by database not defined")
	}
	var err error

	DB, err = pgxpool.New(context.Background(), databaseUrl)

	if err != nil {
		return fmt.Errorf("connection error %w", err)
	}

	err = DB.Ping(context.Background())

	if err != nil {
		return fmt.Errorf("ping database error: %w, databaseUrl: %s", err, databaseUrl)
	}
	fmt.Println("Connection successful")

	return nil
}

func Close() {
	if DB != nil {
		DB.Close()
	}
}
