package infrastructure

import (
	"errors"

	"github.com/KingDaemonX/evolve-mod-ddd-sample/services/transactions/domain/entity"
	"github.com/KingDaemonX/evolve-mod-ddd-sample/services/transactions/domain/repository"
	entity2 "github.com/KingDaemonX/evolve-mod-ddd-sample/services/users/domain/entity"
	"gorm.io/gorm"
)

type transactionInfra struct {
	conn *gorm.DB
}

func NewTrasactionInfra(conn *gorm.DB) *transactionInfra {
	return &transactionInfra{conn: conn}
}

var _ repository.TransactionRepository = &transactionInfra{}

func (t *transactionInfra) CreateTransfer(userId string, transfer entity.Transfer) (any, error) {
	var (
		sourceUser, targetUser entity2.User
		transactions           entity.Transactions
		err                    error
	)

	if err = t.conn.Debug().Find(&sourceUser, "id ?=", userId).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			err = errors.New("this user record does not exist")
			return nil, err
		}
		return nil, err
	}

	// we assume that the account number was gotten from the read account number logic
	// so it is save to pass
	if err = t.conn.Debug().Find(&targetUser, "user.account_number ?=", transfer.TargetAccountNumber).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			err = errors.New("cannot find user with this account number")
			return nil, err
		}
		return nil, err
	}

	// * add pin security here later when it works
	// todo : add pin check here
	// sourceUserCB := user.Account.Balance - transfer.Amount
	// targetUserCB := targetUser.Account.Balance + transfer.Amount
	sourceUser.Account.Balance -= transfer.Amount
	targetUser.Account.Balance += transfer.Amount

	SourceTransaction := &entity.Transactions{
		SourceUserID: sourceUser.ID,
		// SourceAccountID:     sourceUser.Account.Number,
		SourceAccountNumber: sourceUser.Account.Number,
		TargetUserID:        targetUser.ID,
		// TargetAccountID: ,
		TargetAccountNumber:           targetUser.Account.Number,
		TransactionType:               "debit",
		SourceTransactionAmount:       transfer.Amount,
		TargetTransactionAmount:       transfer.Amount,
		SourceBalanceAfterTransaction: sourceUser.Account.Balance,
		Description:                   transfer.Description,
		Status:                        "success",
	}

	TargetTransaction := &entity.Transactions{
		SourceUserID: sourceUser.ID,
		// SourceAccountID:     sourceUser.Account.Number,
		SourceAccountNumber: sourceUser.Account.Number,
		TargetUserID:        targetUser.ID,
		// TargetAccountID: ,
		TargetAccountNumber:           targetUser.Account.Number,
		TransactionType:               "credit",
		SourceTransactionAmount:       transfer.Amount,
		TargetTransactionAmount:       transfer.Amount,
		TargetBalanceAfterTransaction: targetUser.Account.Balance,
		Description:                   transfer.Description,
		Status:                        "success",
	}

	go func() (any, error) {
		if err := t.conn.Debug().Find(&transactions, "user.transaction ?=", userId).Save(SourceTransaction).Error; err != nil {
			return nil, err
		}
		return nil, err
	}()

	go func() (any, error) {
		if err := t.conn.Debug().Find(&targetUser, "user.transaaction ?=", targetUser.ID).Save(TargetTransaction).Error; err != nil {
			return nil, err
		}
		return nil, err
	}()

	return "transfer successful", nil
}
