package entity

import "time"

type Account struct {
	ID        uint        `gorm:"primaryKey"`
	UID       uint        `json:"uid"`
	Number    string      `json:"accountNumber" gorm:"account_number"`
	Active    bool        `json:"active" gorm:""`
	Balance   float64     `json:"accountBalance" gorm:"account_balance"`
	Pin       *AccountPin `json:"-"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type AccountPin struct {
	Pin string `json:"pin,omitempty"`
}
