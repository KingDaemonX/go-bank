package repository

import "github.com/KingDaemonX/evolve-mod-ddd-sample/services/transactions/domain/entity"

type Transaction interface {
	CreateTransaction(entity.Transfer) (any, error)
}
