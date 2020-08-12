package web

import (
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/sirupsen/logrus"

	"go-rest-project/logger"
)

const (
	sessionName = "sessionid"
	sessionKey  = "user_id"
)

var (
	key   = []byte("super-secret-key")
	store = sessions.NewCookieStore(key)

	excludePaths = []string{login, registration}
)

// IsAuthorized middleware that checks if user is authenticated
func IsAuthorized(next http.Handler) http.Handler {
	logEntry := logger.LogEntry("is-authorized-middleware")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		for _, path := range excludePaths {
			if r.URL.Path == path {
				next.ServeHTTP(w, r)
				return
			}
		}
		_, isAuthorized := getSessionValue(r, logEntry)
		if !isAuthorized {
			SendUnauthorized(w, logEntry)
			return
		}
		next.ServeHTTP(w, r)
	})
}

// GetAuthorizedUserID returns authorized user
func GetAuthorizedUserID(r *http.Request, logEntry *logrus.Entry) uint {
	user, _ := getSessionValue(r, logEntry)
	return user
}

// SaveSession saves session to cookies store
func SaveSession(w http.ResponseWriter, r *http.Request, userID uint) error {
	session, err := store.Get(r, sessionName)
	if err != nil {
		return err
	}
	session.Values[sessionKey] = userID
	return session.Save(r, w)
}

// RemoveSession removes session from cookies store
func RemoveSession(w http.ResponseWriter, r *http.Request) error {
	session, err := store.Get(r, sessionName)
	if err != nil {
		return err
	}
	session.Values[sessionKey] = nil
	return session.Save(r, w)
}

func getSessionValue(r *http.Request, logEntry *logrus.Entry) (uint, bool) {
	session, err := store.Get(r, sessionName)
	if err != nil {
		logEntry.Error(err)
		return 0, false
	}

	userID, ok := session.Values[sessionKey].(uint)
	if !ok {
		logEntry.Debug("userID are not authorized")
		return 0, false
	}
	return userID, true
}
