package entity

import (
	"time"

	"github.com/KingDaemonX/evolve-mod-ddd-sample/services/accounts/domain/entity"
	transacEntity "github.com/KingDaemonX/evolve-mod-ddd-sample/services/transactions/domain/entity"
	"gorm.io/gorm"
)

type User struct {
	ID           uint                          `gorm:"primaryKey"`
	FirstName    string                        `json:"firstName" validate:"gte=2,lte=100,required"`
	LastName     string                        `json:"lastName" validate:"gte=2,lte=100,required"`
	Email        string                        `gorm:"unique,email" json:"" validate:"required,email"`
	Password     string                        `json:"password" validate:"gte=8"` // todo: add regex
	Account      *entity.Account               `json:"account,omitempty" gorm:"foreignKey:UID;constraint:OnUpdate:CASCADE,Ondelete:SET NULL;"`
	Transactions []*transacEntity.Transactions `json:"transactions,omitempty" gorm:"foreignKey:UID;constraint:OnUpdate:CASCADE,Ondelete:SET NULL, "`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt
}
