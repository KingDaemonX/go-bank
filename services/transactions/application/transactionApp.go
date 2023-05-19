package application

import (
	"github.com/KingDaemonX/evolve-mod-ddd-sample/services/transactions/domain/entity"
	"github.com/KingDaemonX/evolve-mod-ddd-sample/services/transactions/domain/repository"
)

type TransactionApp struct {
	Transacation repository.TransactionRepository
}

var _ TransactionApplication = &TransactionApp{}

type TransactionApplication interface {
	CreateTransfer(string, entity.Transfer) (any, error)
}

func (ta *TransactionApp) CreateTransfer(userId string, transfer entity.Transfer) (any, error) {
	return ta.Transacation.CreateTransfer(userId, transfer)
}
