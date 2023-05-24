package repository

import "github.com/KingDaemonX/evolve-mod-ddd-sample/services/accounts/domain/entity"

type AcccountRepository interface {
	CreateAccount(*entity.Account) (any, error)
	ReadAccountByAccountNumber(string) (any, error)
	CreatePin(*entity.AccountPin) (any, error)
}
