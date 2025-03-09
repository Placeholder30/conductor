package db

import (
	"database/sql"
	"log/slog"
	"os"

	"github.com/placeholder30/conductor/cmd/config"
)

func NewPostgresStorage(cfg config.PgConfig) (*sql.DB, error) {

	db, err := sql.Open("postgres", cfg.FormatDSN())
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	return db, nil
}
