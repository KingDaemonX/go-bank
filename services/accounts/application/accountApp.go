package application

import (
	"github.com/KingDaemonX/evolve-mod-ddd-sample/services/accounts/domain/entity"
	"github.com/KingDaemonX/evolve-mod-ddd-sample/services/accounts/domain/repository"
)

type accountApp struct {
	account repository.AcccountRepository
}

type AccountApplication interface {
	CreateAccount(*entity.Account) (any, error)
	ReadAccountByAccountNumber(string) (any, error)
	UpdateAccountBalance(*entity.UpdateBalance) (any, error)
}

var _ repository.AcccountRepository = &accountApp{}

func (aa *accountApp) CreateAccount(account *entity.Account) (any, error) {
	return aa.account.CreateAccount(account)
}

func (aa *accountApp) ReadAccountByAccountNumber(accountNumber string) (any, error) {
	return aa.account.ReadAccountByAccountNumber(accountNumber)
}

func (aa *accountApp) UpdateAccountBalance(fundsInfo *entity.UpdateBalance) (any, error) {
	return aa.account.UpdateAccountBalance(fundsInfo)
}
