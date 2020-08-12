package web

import (
	"io"
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"

	"go-rest-project/logger"
	"go-rest-project/model"
	"go-rest-project/service"
)

func GetAvatar(w http.ResponseWriter, r *http.Request) {
	logEntry := logger.LogEntry("get-avatar-handler")

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

	image, err := service.ImageOpen(user.AvatarURL)
	if err != nil {
		SendInternalServerError(w, err, logEntry)
		return
	}

	if _, err := io.Copy(w, image); err != nil {
		SendInternalServerError(w, err, logEntry)
	}
}

// SetAvatar set avatar by user ID
func SetAvatar(w http.ResponseWriter, r *http.Request) {
	logEntry := logger.LogEntry("set-avatar-handler")

	file, _, err := r.FormFile("avatar")
	if err != nil {
		SendBadRequest(w, err, logEntry)
	}
	defer file.Close()

	userID, err := getUserIDFromRequest(r)
	if err != nil {
		SendBadRequest(w, err, logEntry)
		return
	}
	cookieUserID := GetAuthorizedUserID(r, logEntry)

	if cookieUserID != userID {
		SendBadRequest(w, errors.Errorf("you can edit only your account, your id=%d, id you want to edit=%d", cookieUserID, userID), logEntry)
		return
	}

	avatarPath, err := service.ImageUpload(file, userID)
	if err != nil {
		SendInternalServerError(w, err, logEntry)
		return
	}

	user := model.User{ID: userID, AvatarURL: avatarPath}

	err = model.Users().Update(model.User{ID: userID, AvatarURL: avatarPath})
	if err != nil {
		SendInternalServerError(w, err, logEntry)
		return
	}
	renderUser(w, http.StatusOK, user, logEntry)
}
