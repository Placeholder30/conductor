package user

import (
	"database/sql"
	"net/http"
)

type UserHandler struct {
	Store *sql.DB
	mux   *http.ServeMux
}

func NewUserHandler(db *sql.DB, mux *http.ServeMux) *UserHandler {
	return &UserHandler{Store: db, mux: mux}
}

func (u UserHandler) UserRoutes() {
	u.mux.HandleFunc("GET /user", u.getUser)
	u.mux.HandleFunc("GET /user/:id", u.getUserByID)
	u.mux.HandleFunc("PUT /user/:id", u.updateUser)
	u.mux.HandleFunc("POST /user", u.createUser)
	u.mux.HandleFunc("DELETE /user/:id", u.deleteUser)
}

func (u UserHandler) createUser(w http.ResponseWriter, r *http.Request) {

}

func (u UserHandler) getUser(w http.ResponseWriter, r *http.Request) {

	// Your handler logic here
}

func (u UserHandler) getUserByID(w http.ResponseWriter, r *http.Request) {
	// Your handler logic here
}
func (u UserHandler) updateUser(w http.ResponseWriter, r *http.Request) {
	// Your handler logic here
}

func (u UserHandler) deleteUser(w http.ResponseWriter, r *http.Request) {
	// Your handler logic here
}
