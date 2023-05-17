package entity

import "time"

type Transactions struct {
	ID                            uint      `json:""`
	LinkID                        string    `json:""`
	SourceUserID                  string    `json:""`
	SourceAccountID               string    `json:""`
	SourceAccountNumber           string    `json:""`
	TargetUserID                  string    `json:""`
	TargetAccountID               string    `json:""`
	TargetAccountNumber           string    `json:""`
	TransactionType               string    `json:""`
	SourceTransactionAmount       float64   `json:""`
	TargetTransactionAmount       float64   `json:""`
	SourceBalanceAfterTransaction float64   `json:""`
	TargetBalanceAfterTransaction float64   `json:""`
	Status                        string    `json:""`
	TransactionDate               time.Time `json:""`
	UpdatedAt                     time.Time `json:""`
}
type Transfer struct {
	SourceUserID        string  `json:"" validate:""`
	SourceAccountID     string  `json:"" validate:""`
	SourceAccountNumber string  `json:"" validate:""`
	TargetAccountNumber string  `json:"" validate:""`
	Amount              float64 `json:"" validate:""`
	TransactionPin      string  `json:"" validate:""`
}
