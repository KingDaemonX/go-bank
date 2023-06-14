package entity

import (
	"log"

	"github.com/KingDaemonX/evolve-mod-ddd-sample/services/users/infrastructure/encrypt"
	"gorm.io/gorm"
)

// gorm hook
func (u *User) BeforeCreate(tx *gorm.DB) error {
	password, err := encrypt.HashPassword(u.Password)
	if err != nil {
		log.Println("Error : ", err.Error())
		return nil
	}

	u.Password = string(password)
	return nil
}
