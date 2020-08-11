package model

import (
	"go-rest-project/db"

	"github.com/sirupsen/logrus"
)

// Migrate runs migration for all models
func Migrate(log *logrus.Logger) {
	db.Migrate(User{})
	log.Info("Migration successfully completed")
}
