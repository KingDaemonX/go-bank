package entity

import "time"

type Account struct {
	UID       uint        `gorm:"foreignKey;user"`
	Number    uint        `json:"accountNumber" gorm:"account_number"`
	Active    bool        `json:"" gorm:""`
	Balance   float64     `json:"accountBalance" gorm:"account_balance"`
	Pin       *AccountPin `json:"-"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type AccountPin struct {
	Pin string `json:"omit"`
}

type UpdateBalance struct {
	ID   uint
	Fund float64
}
