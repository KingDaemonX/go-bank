package entity

import (
	"time"

	"github.com/KingDaemonX/evolve-mod-ddd-sample/services/accounts/domain/entity"
	"gorm.io/gorm"
)

type User struct {
	ID        uint            `gorm:"primaryKey"`
	FirstName string          `gorm:"" json:"" validate:""`
	LastName  string          `gorm:"" json:"" validate:""`
	Email     string          `gorm:"" json:"" validate:""`
	Password  string          `gorm:"" json:"" validate:""`
	Account   *entity.Account `gorm:""`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

type PublicAccInfo struct {
	ID            uint   `gorm:"primaryKey"`
	FirstName     string `gorm:"" json:"" validate:""`
	LastName      string `gorm:"" json:"" validate:""`
	AccountNumber string `gorm:"" json:"" validate:""`
}

type Users []User

func (u *User) PublicAccInfo() interface{} {
	return &PublicAccInfo{
		ID:            u.ID,
		FirstName:     u.FirstName,
		LastName:      u.LastName,
		AccountNumber: u.Account.Number,
	}
}

func (users Users) PublicAccInfos() []interface{} {
	result := make([]interface{}, len(users))

	for index, info := range users {
		result[index] = info.PublicAccInfo()
	}
	return result
}
