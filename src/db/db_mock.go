package db

import (
	"github.com/jinzhu/gorm"
	mocket "github.com/selvatico/go-mocket"
)

func Mock() {
	mocket.Catcher.Register()
	mocket.Catcher.Logging = true

	dbMock, _ := gorm.Open(mocket.DriverName, "gorm_mock")
	db = dbMock
}
