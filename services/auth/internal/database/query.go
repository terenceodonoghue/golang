package database

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type Table string

const (
	ExternalProviders  Table = "external_providers"
	UserAccounts       Table = "user_accounts"
	UserCredentials    Table = "user_credentials"
	UserCredentialsExt Table = "user_credentials_ext"
)

func Exists(conn *pgx.Conn, table_name Table, attr string, value string) (bool, error) {
	var exists bool

	err := conn.QueryRow(context.Background(), "SELECT EXISTS(SELECT 1 FROM "+string(table_name)+" WHERE "+attr+" = $1);", value).Scan(&exists)

	return exists, err
}
