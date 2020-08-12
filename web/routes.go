package web

import (
	"net/http"

	"github.com/gorilla/mux"
)

const (
	users        = "/users"
	usersID      = "/users/{ID}"
	avatarID     = "/avatar/{ID}"
	login        = "/login"
	logout       = "/logout"
	registration = "/registration"
)

// InitRouter returns router with all routes
func InitRouter() *mux.Router {
	r := mux.NewRouter()

	r.Use(IsAuthorized)

	r.HandleFunc(users, GetUsers).Methods("GET")
	r.HandleFunc(usersID, GetUserByID).Methods("GET")
	r.HandleFunc(usersID, DeleteUserByID).Methods("DELETE")
	r.HandleFunc(usersID, ModifyUserByID).Methods("PUT")

	r.HandleFunc(avatarID, SetAvatar).Methods("PUT")

	r.HandleFunc(login, Login).Methods("POST")
	r.HandleFunc(logout, Logout).Methods("POST")
	r.HandleFunc(registration, Registration).Methods("POST")

	r.PathPrefix("/images/").Handler(http.StripPrefix("/images/",
		http.FileServer(http.Dir("static/images/"))))

	return r
}
