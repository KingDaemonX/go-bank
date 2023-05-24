package persistent

import (
	"errors"

	"github.com/KingDaemonX/evolve-mod-ddd-sample/services/accounts/domain/entity"
	"github.com/KingDaemonX/evolve-mod-ddd-sample/services/accounts/domain/repository"

	"github.com/KingDaemonX/evolve-mod-ddd-sample/services/accounts/infrastructure/helpers"
	entity2 "github.com/KingDaemonX/evolve-mod-ddd-sample/services/users/domain/entity"
	"gorm.io/gorm"
)

type AccountInfra struct {
	database *gorm.DB
	helper   *helpers.AccountHelper
}

func NewAccountInfra(conn *gorm.DB) *AccountInfra {
	return &AccountInfra{database: conn}
}

var _ repository.AcccountRepository = &AccountInfra{}

func (a *AccountInfra) CreateAccount(account *entity.Account) (any, error) {
	// if userId / email does not exist ? \ cancel account creation
	var user entity2.User
	if err := a.database.Debug().Find(user, "id ?=", account.UID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			err = errors.New("user cred does not exist for account creation")
			return nil, err
		}
		return nil, err
	}

	// else continue and build with logic
	account.Number = a.helper.GenerateAccNumber(a.database)
	account.Active = false
	account.Balance = 0.00
	a.database.Debug().Save(account)

	return account, nil
}

func (a *AccountInfra) ReadAccountByAccountNumber(account string) (any, error) {
	// check if user with account exist
	return "", nil
}

func (a *AccountInfra) CreatePin(pin *entity.AccountPin) (any, error) {
	return nil, nil
}
