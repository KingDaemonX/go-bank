package entity

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint   `gorm:"primaryKey"`
	FirstName string `gorm:"" json:"" validate:""`
	LastName  string `gorm:"" json:"" validate:""`
	Email     string `gorm:"" json:"" validate:""`
	Password  string `gorm:"" json:"" validate:""`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
