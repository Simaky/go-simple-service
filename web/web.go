package web

import (
	"encoding/json"
	"log"
	"net/http"

	"go-rest-project/config"
)

// ListenAndServe load routes and run server
func ListenAndServe() error {
	log.Println("web server has been started")
	return http.ListenAndServe(config.Config.Port, InitRouter())
}

// SendResponse send JSON response
func SendResponse(w http.ResponseWriter, header int, response interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(header)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		SendInternalServerError(w, err)
	}
}

// SendInternalServerError send internal serve error
func SendInternalServerError(w http.ResponseWriter, err error) {
	sendError(w, http.StatusInternalServerError, err)
}

// SendBadRequest send bad request
func SendBadRequest(w http.ResponseWriter, err error) {
	sendError(w, http.StatusBadRequest, err)
}

// SendNotFound send not found
func SendNotFound(w http.ResponseWriter, err error) {
	sendError(w, http.StatusNotFound, err)
}

// SendUnauthorized send unauthorized error
func SendUnauthorized(w http.ResponseWriter, err error) {
	sendError(w, http.StatusUnauthorized, err)
}

func sendError(w http.ResponseWriter, header int, error error) {
	w.WriteHeader(header)
	w.Write([]byte(error.Error()))
}
