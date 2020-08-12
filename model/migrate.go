package model

import (
	"github.com/sirupsen/logrus"

	"go-rest-project/db"
)

// Migrate runs migration for all models
func Migrate(log *logrus.Logger) {
	db.Migrate(User{})
	log.Info("Migration successfully completed")
}
