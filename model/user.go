package model

import "time"

// User is a user model
type User struct {
	ID          uint   `gorm:"primary_key"`
	Login       string `json:"login"`
	Password    string `json:"password,omitempty"`
	Email       string `json:"email"`
	AvatarURL   string `json:"avatar_url"`
	DateOfBirth int64  `json:"date_of_birth"`

	CreatedAt time.Time `gorm:"default:NULL DEFAULT CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"default:NULL ON UPDATE CURRENT_TIMESTAMP"`
	Address
}

// Address user address
type Address struct {
	Country string `json:"country"`
	City    string `json:"city"`
}
