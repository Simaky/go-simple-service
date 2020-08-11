package web

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"

	"go-rest-project/config"
)

const startedMessage = "web server has been started"

// ListenAndServe load routes and run server
func ListenAndServe(log *logrus.Logger) error {
	log.Info(startedMessage)
	fmt.Println(startedMessage)
	return http.ListenAndServe(config.Config.Port, InitRouter())
}

// SendResponse send JSON response
func SendResponse(w http.ResponseWriter, header int, response interface{}, logEntry *logrus.Entry) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(header)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		SendInternalServerError(w, err, logEntry)
	}
	logEntry.Debugf("sent response: %d, %s", header, response)
}

// SendInternalServerError send internal serve error
func SendInternalServerError(w http.ResponseWriter, err error, logEntry *logrus.Entry) {
	sendError(w, http.StatusInternalServerError, err, logEntry)
}

// SendBadRequest send bad request
func SendBadRequest(w http.ResponseWriter, err error, logEntry *logrus.Entry) {
	sendError(w, http.StatusBadRequest, err, logEntry)
}

// SendNotFound send not found
func SendNotFound(w http.ResponseWriter, err error, logEntry *logrus.Entry) {
	sendError(w, http.StatusNotFound, err, logEntry)
}

// SendUnauthorized send unauthorized error
func SendUnauthorized(w http.ResponseWriter, logEntry *logrus.Entry) {
	sendError(w, http.StatusUnauthorized, errors.New("you are not authorized"), logEntry)
}

func sendError(w http.ResponseWriter, header int, error error, logEntry *logrus.Entry) {
	w.WriteHeader(header)
	_, err := w.Write([]byte(error.Error()))
	if err != nil {
		logEntry.Error(error)
	}
	logEntry.Debugf("http %d error: %s", header, error.Error())
}
