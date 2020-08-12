package web

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"

	"go-rest-project/logger"
	"go-rest-project/model"
)

// Login creates new user
func Login(w http.ResponseWriter, r *http.Request) {
	logEntry := logger.LogEntry("login-handler")

	var user model.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		SendBadRequest(w, err, logEntry)
		return
	}

	userDB, err := model.Users().GetByLogin(user.Login)
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			SendInternalServerError(w, err, logEntry)
		}
		SendBadRequest(w, errors.Errorf("user with login: '%s' is not exist", user.Login), logEntry)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(userDB.Password), []byte(user.Password))
	if err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			SendBadRequest(w, errors.New("wrong password"), logEntry)
			return
		}
		SendInternalServerError(w, err, logEntry)
		return
	}

	err = SaveSession(w, r, userDB.ID)
	if err != nil {
		SendInternalServerError(w, err, logEntry)
		return
	}
	renderUser(w, http.StatusOK, userDB, logEntry)
}

// Logout creates new user
func Logout(w http.ResponseWriter, r *http.Request) {
	logEntry := logger.LogEntry("login-handler")

	err := RemoveSession(w, r)
	if err != nil {
		SendInternalServerError(w, err, logEntry)
		return
	}
	SendUnauthorized(w, logEntry)
}

// Registration creates new user
func Registration(w http.ResponseWriter, r *http.Request) {
	logEntry := logger.LogEntry("registration-handler")
	var user model.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		SendBadRequest(w, err, logEntry)
		return
	}

	err = userValidation(user)
	if err != nil {
		SendBadRequest(w, err, logEntry)
		return
	}

	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		SendInternalServerError(w, err, logEntry)
		return
	}
	user.Password = string(password)

	err = model.Users().Add(user)
	if err != nil {
		SendInternalServerError(w, err, logEntry)
		return
	}

	err = SaveSession(w, r, user.ID)
	if err != nil {
		SendInternalServerError(w, err, logEntry)
		return
	}
	renderUser(w, http.StatusCreated, user, logEntry)
}

func userValidation(user model.User) error {
	if strings.TrimSpace(user.Login) == "" {
		return errors.New("user login can't be empty")
	}
	if strings.TrimSpace(user.Password) == "" || strings.Count(user.Password, "")-1 < 6 {
		return errors.New("user password should contains 6 and more symbols")
	}

	_, err := model.Users().GetByLogin(user.Login)
	if err != gorm.ErrRecordNotFound {
		return errors.New("user with that login already exist")
	}
	return nil
}
