package web

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"go-rest-project/db"
	"go-rest-project/model"
)

func TestGetUsers(t *testing.T) {
	db.Mock()
	expectedUser := model.User{ID: 1, Login: "test"}

	model.MockUsers(t, expectedUser)

	req, err := http.NewRequest("GET", withVersion(users), nil)
	assert.NoError(t, err)

	r := httptest.NewRecorder()
	handler := http.HandlerFunc(GetUsers)

	handler.ServeHTTP(r, req)

	if status := r.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var user []model.User
	err = json.NewDecoder(r.Body).Decode(&user)
	assert.NoError(t, err)
	assert.ObjectsAreEqual(expectedUser, user)
}
