package database

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

func Connect() (*pgx.Conn, error) {
	connStr, err := connectionString()
	if err != nil {
		return nil, err
	}

	conn, err := pgx.Connect(context.Background(), connStr)
	return conn, err
}

func connectionString() (string, error) {
	var (
		DB_USER = os.Getenv("DB_USER")
		DB_PASS = os.Getenv("DB_PASS")
		DB_HOST = os.Getenv("DB_HOST")
		DB_PORT = os.Getenv("DB_PORT")
		DB_NAME = os.Getenv("DB_NAME")
	)

	user, err := os.ReadFile(DB_USER)
	if err != nil {
		return "", err
	}

	password, err := os.ReadFile(DB_PASS)
	if err != nil {
		return "", err
	}

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", string(user), string(password), DB_HOST, DB_PORT, DB_NAME)
	return connStr, err
}
