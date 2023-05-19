package repository

import "github.com/KingDaemonX/evolve-mod-ddd-sample/services/transactions/domain/entity"

type TransactionRepository interface {
	CreateTransfer(string, entity.Transfer) (any, error)
}
