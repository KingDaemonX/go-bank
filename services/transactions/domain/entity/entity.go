package entity

import "time"

type Transactions struct {
	ID                            uint    `gorm:"primaryKey"`
	UID                           uint    `json:"uid"`
	LinkID                        string  `json:"linkId"`
	SourceUserID                  uint    `json:"sourceUserId"`
	SourceAccountID               string  `json:"sourceAccountId"`
	SourceAccountNumber           string  `json:"sourceAccountNumber"`
	TargetUserID                  uint    `json:"targetUserId"`
	TargetAccountID               string  `json:"targetaccountId"`
	TargetAccountNumber           string  `json:"targetAccountNumber"`
	TransactionType               string  `json:"transactionType"`
	SourceTransactionAmount       float64 `json:"sourceTransactionAmount"`
	TargetTransactionAmount       float64 `json:"targetTransactionAmount"`
	SourceBalanceAfterTransaction float64 `json:"sourceBalanceAfterTransaction"`
	TargetBalanceAfterTransaction float64 `json:"targetBalanceAfterTransaction"`
	Description                   string  `json:"Description"`
	Status                        string  `json:"status"`
	TransactionDate               time.Time
	UpdatedAt                     time.Time
}
type Transfer struct {
	TargetAccountNumber string  `json:"targetAccountNumber" validate:""`
	Amount              float64 `json:"amount" validate:""`
	Description         string  `json:"description" validate:""`
	TransactionPin      string  `json:"transactionPin" validate:""`
}
