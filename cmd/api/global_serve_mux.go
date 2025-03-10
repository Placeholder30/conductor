package api

import (
	"database/sql"
	"net/http"

	"github.com/placeholder30/conductor/cmd/services/user"
)

func Muxer(db *sql.DB) *http.ServeMux {
	mux := http.NewServeMux()
	user.NewUserHandler(db, mux).UserRoutes()
	return mux
}
