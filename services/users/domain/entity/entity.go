package entity

import (
	"time"

	"github.com/KingDaemonX/evolve-mod-ddd-sample/services/accounts/domain/entity"
	transacEntity "github.com/KingDaemonX/evolve-mod-ddd-sample/services/transactions/domain/entity"
	"gorm.io/gorm"
)

type User struct {
	ID           uint   `gorm:"primaryKey"`
	FirstName    string `gorm:"" json:"" validate:""`
	LastName     string `gorm:"" json:"" validate:""`
	Email        string `gorm:"" json:"" validate:""`
	Password     string `gorm:"" json:"" validate:""`
	Account      *entity.Account
	Transactions []*transacEntity.Transactions
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt
}
