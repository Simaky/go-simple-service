package model

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"go-rest-project/db"
)

func TestGetAllUsers(t *testing.T) {
	db.Mock()
	userLogin := "test"

	MockUsers(t, User{ID: 1, Login: userLogin})

	users, err := Users().All()
	assert.NoError(t, err)

	for _, user := range users {
		assert.Equal(t, userLogin, user.Login)
	}
	assert.Len(t, users, 1)
}

func TestGetByID(t *testing.T) {
	db.Mock()
	userID := uint(1)

	MockUsers(t, User{ID: userID, Login: "test"})

	user, err := Users().GetByID(userID)
	assert.NoError(t, err)
	assert.Equal(t, userID, user.ID)
}

func TestGetByLogin(t *testing.T) {
	db.Mock()
	userLogin := "test"

	MockUsers(t, User{ID: 1, Login: userLogin})

	user, err := Users().GetByLogin(userLogin)
	assert.NoError(t, err)
	assert.Equal(t, userLogin, user.Login)
}

func TestDeleteUser(t *testing.T) {
	db.Mock()
	user := User{ID: 1, Login: "test"}

	MockUsers(t, user)

	err := Users().Delete(user)
	assert.NoError(t, err)
}

func TestUpdateUser(t *testing.T) {
	db.Mock()
	user := User{ID: 1, Login: "test"}

	MockUsers(t, user)

	err := Users().Update(user)
	assert.NoError(t, err)
}
