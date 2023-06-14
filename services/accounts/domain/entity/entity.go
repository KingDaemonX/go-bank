package entity

import "time"

type Account struct {
	ID        uint        `gorm:"primaryKey"`
	UID       uint        `json:"uid"`
	Number    string      `json:"accountNumber" gorm:"account_number"`
	Active    bool        `json:"active" gorm:""`
	Balance   float64     `json:"accountBalance" gorm:"account_balance"`
	Pin       *AccountPin `json:"-" gorm:"foreignKey:UID;"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type AccountPin struct {
	UID uint
	Pin string `json:"pin,omitempty"`
}
