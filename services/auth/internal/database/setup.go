package database

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/terenceodonoghue/golang/libs/env"
)

func New() (*pgx.Conn, error) {
	connStr, err := connStr()
	if err != nil {
		return nil, err
	}

	conn, err := pgx.Connect(context.Background(), connStr)
	return conn, err
}

func connStr() (string, error) {

	var (
		POSTGRES_HOST          = env.GetOrDefault("POSTGRES_HOST", "localhost")
		POSTGRES_PORT          = env.GetOrDefault("POSTGRES_PORT", 5432)
		POSTGRES_USER_FILE     = env.GetOrDefault("POSTGRES_USER_FILE", "./.secrets/postgres_user.txt")
		POSTGRES_PASSWORD_FILE = env.GetOrDefault("POSTGRES_PASSWORD_FILE", "./.secrets/postgres_password.txt")
		POSTGRES_DB_FILE       = env.GetOrDefault("POSTGRES_DB_FILE", "./.secrets/postgres_db.txt")
	)

	user, err := os.ReadFile(POSTGRES_USER_FILE)
	if err != nil {
		return "", err
	}

	password, err := os.ReadFile(POSTGRES_PASSWORD_FILE)
	if err != nil {
		return "", err
	}

	db, err := os.ReadFile(POSTGRES_DB_FILE)
	if err != nil {
		return "", err
	}

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", user, password, POSTGRES_HOST, POSTGRES_PORT, db)
	return connStr, err
}
