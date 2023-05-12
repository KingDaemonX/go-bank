package persistent

import (
	"github.com/KingDaemonX/evolve-mod-ddd-sample/services/accounts/domain/entity"
	"github.com/KingDaemonX/evolve-mod-ddd-sample/services/accounts/domain/repository"
	"gorm.io/gorm"
)

type AccountInfra struct {
	conn *gorm.DB
}

func NewAccountInfra(conn *gorm.DB) *AccountInfra {
	return &AccountInfra{conn: conn}
}

var _ repository.AcccountRepository = &AccountInfra{}

func (a *AccountInfra) CreateAccount(account *entity.Account) (any, error) {
	// if userId / email does not exist ? \ cancel account creation
	// else continue and build with logic
	return "", nil
}

func (a *AccountInfra) ReadAccountByAccountNumber(account string) (any, error) {
	return "", nil
}

func (a *AccountInfra) UpdateAccountBalance(fundsInfo *entity.UpdateBalance) (any, error) {
	return "", nil
}
