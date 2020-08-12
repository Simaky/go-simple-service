package web

import "github.com/gorilla/mux"

const (
	version = "/v1"

	users        = "/users"
	usersID      = "/users/{ID}"
	avatarID     = "/avatar/{ID}"
	login        = "/login"
	logout       = "/logout"
	registration = "/registration"

	staticImages = "/images/"
)

// InitRouter returns router with all routes
func InitRouter() *mux.Router {
	r := mux.NewRouter()

	r.Use(IsAuthorized)

	r.HandleFunc(withVersion(users), GetUsers).Methods("GET")
	r.HandleFunc(withVersion(usersID), GetUserByID).Methods("GET")
	r.HandleFunc(withVersion(usersID), DeleteUserByID).Methods("DELETE")
	r.HandleFunc(withVersion(usersID), ModifyUserByID).Methods("PUT")

	r.HandleFunc(withVersion(avatarID), SetAvatar).Methods("PUT")
	r.HandleFunc(withVersion(avatarID), GetAvatar).Methods("GET")

	r.HandleFunc(withVersion(login), Login).Methods("POST")
	r.HandleFunc(withVersion(logout), Logout).Methods("POST")
	r.HandleFunc(withVersion(registration), Registration).Methods("POST")

	return r
}

func withVersion(path string) string {
	return version + path
}
