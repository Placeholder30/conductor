package api

import (
	"database/sql"
	"fmt"
	"log/slog"
	"net/http"
)

type apiServer struct {
	port string
	db   *sql.DB
}

func NewApiServer(port string, db *sql.DB) *apiServer {
	return &apiServer{
		port: port,
		db:   db,
	}
}

func (a *apiServer) Run() error {
	mux := http.Server{
		Addr:    a.port,
		Handler: Muxer(a.db),
	}
	slog.Info(fmt.Sprintf("server is running on port %s", a.port))
	return mux.ListenAndServe()

}
