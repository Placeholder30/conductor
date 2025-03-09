package api

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/placeholder30/conductor/cmd/config"
)

type apiServer struct {
	port string
	db   *config.PgConfig
}

func NewApiServer(port string, cfg *config.PgConfig) *apiServer {
	return &apiServer{
		port: port,
		db:   cfg,
	}
}

func (a *apiServer) Run() error {
	mux := http.Server{
		Addr:    a.port,
		Handler: nil,
	}
	slog.Info(fmt.Sprintf("server is running on port %s", a.port))
	return mux.ListenAndServe()

}
