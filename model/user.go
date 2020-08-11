package model

import (
	"github.com/jinzhu/gorm"
)

// User is a user model
type User struct {
	gorm.Model
	Address

	Login       string `json:"login"`
	Password    string `json:"password"`
	Email       string `json:"email"`
	AvatarURL   string `json:"avatar_url"`
	DateOfBirth int64  `json:"date_of_birth"`
}

// Address user address
type Address struct {
	Country string `json:"country"`
	City    string `json:"city"`
}
