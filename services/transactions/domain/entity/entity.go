package entity

import "time"

type Transactions struct {
	UID                           uint      // a foriegnKey
	ID                            uint      `json:""`
	LinkID                        string    `json:""`
	SourceUserID                  uint      `json:""`
	SourceAccountID               string    `json:""`
	SourceAccountNumber           string    `json:""`
	TargetUserID                  uint      `json:""`
	TargetAccountID               string    `json:""`
	TargetAccountNumber           string    `json:""`
	TransactionType               string    `json:""`
	SourceTransactionAmount       float64   `json:""`
	TargetTransactionAmount       float64   `json:""`
	SourceBalanceAfterTransaction float64   `json:""`
	TargetBalanceAfterTransaction float64   `json:""`
	Description                   string    `json:""`
	Status                        string    `json:""`
	TransactionDate               time.Time `json:""`
	UpdatedAt                     time.Time `json:""`
}
type Transfer struct {
	TargetAccountNumber string  `json:"" validate:""`
	Amount              float64 `json:"" validate:""`
	Description         string  `json:"" validate:""`
	TransactionPin      string  `json:"" validate:""`
}
