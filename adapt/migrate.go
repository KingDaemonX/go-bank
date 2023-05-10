package adapt

import (
	"gorm.io/gorm"
)

func autoMigrate(conn *gorm.DB) error {
	return conn.AutoMigrate()
}
