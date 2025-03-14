package user

import (
	"database/sql"
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/placeholder30/conductor/cmd/services/response"
)

type UserHandler struct {
	Store *sql.DB
	mux   *http.ServeMux
}

var Validator = validator.New()

func NewUserHandler(db *sql.DB, mux *http.ServeMux) *UserHandler {
	return &UserHandler{Store: db, mux: mux}
}

func (u UserHandler) UserRoutes() {
	u.mux.HandleFunc("GET /users", u.getUser)
	u.mux.HandleFunc("GET /users/:id", u.getUserByID)
	u.mux.HandleFunc("PUT /users/:id", u.updateUser)
	u.mux.HandleFunc("POST /users", u.createUser)
	u.mux.HandleFunc("DELETE /users/:id", u.deleteUser)
}

func (u UserHandler) createUser(w http.ResponseWriter, r *http.Request) {
	slog.Info(r.Method)
	var user CreateUserRequest
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		response.DispatchError(w, err.Error())
		return
	}
	//validate
	if err := Validator.Struct(user); err != nil {
		errors := err.(validator.ValidationErrors)
		response.DispatchError(w, errors.Error())
		return
	}

	response.DispatchSuccess(w, "user created successfully", nil)
	return
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
