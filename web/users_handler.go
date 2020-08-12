package web

import (
	"encoding/json"
	"go-rest-project/logger"
	"go-rest-project/model"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

// GetUsers show all users
func GetUsers(w http.ResponseWriter, r *http.Request) {
	logEntry := logger.LogEntry("get-user-handler")

	users, err := model.Users().All()
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			SendInternalServerError(w, err, logEntry)
		}
		users = make([]*model.User, 0)
	}
	renderUsers(w, http.StatusOK, users, logEntry)
}

// GetUserByID show user by ID
func GetUserByID(w http.ResponseWriter, r *http.Request) {
	logEntry := logger.LogEntry("get-user-by-id-handler")
	userID, err := getUserIDFromRequest(r)
	if err != nil {
		SendBadRequest(w, err, logEntry)
		return
	}

	user, err := model.Users().GetByID(userID)
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			SendInternalServerError(w, err, logEntry)
		}
		SendNotFound(w, err, logEntry)
		return
	}
	renderUser(w, http.StatusOK, user, logEntry)
}

// ModifyUserByID change user data by ID
func ModifyUserByID(w http.ResponseWriter, r *http.Request) {
	logEntry := logger.LogEntry("modify-user-by-id-handler")
	cookieUserID := GetAuthorizedUserID(r, logEntry)

	userID, err := getUserIDFromRequest(r)
	if err != nil {
		SendBadRequest(w, err, logEntry)
		return
	}

	if cookieUserID != userID {
		SendBadRequest(w, errors.Errorf("you can edit only your account, your id=%d, id you want to edit=%d", cookieUserID, userID), logEntry)
		return
	}

	var user model.User
	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		SendBadRequest(w, err, logEntry)
		return
	}

	if strings.TrimSpace(user.Password) != "" {
		password, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			SendInternalServerError(w, err, logEntry)
			return
		}
		user.Password = string(password)
	}

	err = model.Users().Update(user)
	if err != nil {
		SendInternalServerError(w, err, logEntry)
		return
	}
	renderUser(w, http.StatusOK, user, logEntry)
}

// DeleteUserByID deletes user by ID
func DeleteUserByID(w http.ResponseWriter, r *http.Request) {
	logEntry := logger.LogEntry("get-user-handler")
	cookieUserID := GetAuthorizedUserID(r, logEntry)

	userID, err := getUserIDFromRequest(r)
	if err != nil {
		SendBadRequest(w, err, logEntry)
		return
	}

	if userID != cookieUserID {
		SendBadRequest(w, errors.Errorf("you can delete only your account, your id=%d, id you want to delete=%d", cookieUserID, userID), logEntry)
		return
	}

	user, err := model.Users().GetByID(userID)
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			SendInternalServerError(w, err, logEntry)
		}
		SendNotFound(w, err, logEntry)
		return
	}

	err = model.Users().Delete(user)
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			SendInternalServerError(w, err, logEntry)
		}
		SendNotFound(w, err, logEntry)
		return
	}
	renderUser(w, http.StatusNoContent, user, logEntry)
}

func renderUsers(w http.ResponseWriter, header int, users []*model.User, logEntry *logrus.Entry) {
	for _, user := range users {
		user.Password = ""
	}
	SendResponse(w, header, users, logEntry)
}

func renderUser(w http.ResponseWriter, header int, user model.User, logEntry *logrus.Entry) {
	user.Password = ""
	SendResponse(w, header, user, logEntry)
}

func getUserIDFromRequest(r *http.Request) (uint, error) {
	idVar, ok := mux.Vars(r)["ID"]
	if !ok {
		return 0, errors.New("user id is missed")
	}

	userID, err := strconv.Atoi(idVar)
	if err != nil || userID < 0 {
		return 0, errors.New("wrong user id")
	}
	return uint(userID), nil
}
