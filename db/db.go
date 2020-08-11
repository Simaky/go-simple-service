package db

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"

	"go-rest-project/config"
)

// db contains information for current db connection
var db *gorm.DB

// Load loads mysql db
func Load() (err error) {
	db, err = gorm.Open("mysql", loadAccess())
	return
}

// GetConnection returns database connection object
func GetConnection() *gorm.DB {
	return db
}

// Migrate create or update tables for each model
func Migrate(values ...interface{}) {
	db.AutoMigrate(values...)
}

func loadAccess() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		config.Config.Database.Username,
		config.Config.Database.Password,
		config.Config.Database.Server,
		config.Config.Database.Port,
		config.Config.Database.DBName)
}
