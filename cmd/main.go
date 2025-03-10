package main

import (
	"database/sql"
	"log/slog"
	"os"

	_ "github.com/lib/pq"

	"github.com/placeholder30/conductor/cmd/api"
	"github.com/placeholder30/conductor/cmd/config"
	"github.com/placeholder30/conductor/cmd/db"
)

func main() {

	cfg := &config.PgConfig{
		Host:     config.Envs.Host,
		Port:     config.Envs.Port,
		User:     config.Envs.User,
		Password: config.Envs.Password,
		Dbname:   config.Envs.Dbname,
	}

	db, err := db.NewPostgresStorage(*cfg)
	initStorage(db)

	server := api.NewApiServer(":8081", db)

	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	if err := server.Run(); err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

}

func initStorage(db *sql.DB) {
	if err := db.Ping(); err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
	slog.Info("DB: Successfully connected!")
}
