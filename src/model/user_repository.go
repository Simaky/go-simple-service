package model

import (
	"sync"

	"github.com/jinzhu/gorm"

	"go-rest-project/db"
)

var (
	initUsers      sync.Once
	userRepository *UserRepository
)

type UserRepository struct {
	db *gorm.DB
}

// Users singleton, thread-safe, returns pointer to Users repository
func Users() *UserRepository {
	initUsers.Do(func() {
		userRepository = &UserRepository{db: db.GetConnection()}
	})
	return userRepository
}

// All returns list of all users
func (u UserRepository) All() ([]*User, error) {
	var users []*User
	return users, u.db.Find(&users).Error
}

// Add adds new user
func (u UserRepository) Add(user User) error {
	return u.db.Create(&user).Error
}

// Update updates user data
func (u UserRepository) Update(user User) error {
	return u.db.Model(&user).Updates(user).Error
}

// Delete removes user
func (u UserRepository) Delete(user User) error {
	return u.db.Delete(user).Error
}

// GetByID returns user by ID
func (u UserRepository) GetByID(id uint) (User, error) {
	var user User
	return user, u.db.Where("ID = ?", id).First(&user).Error
}

// GetByLogin returns user by Login
func (u UserRepository) GetByLogin(login string) (User, error) {
	var user User
	return user, u.db.Where("login = ?", login).First(&user).Error
}
