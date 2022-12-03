package domain

import "github.com/jimmmisss/banking/dto"

const WITHDRAWAL = "withdrawal"

type Transaction struct {
	TransactionId   string  `db:"transaction_id,omitempty"`
	AccountId       string  `db:"account_id,omitempty"`
	Amount          float64 `db:"amount,omitempty"`
	TransactionType string  `db:"transaction_type,omitempty"`
	TransactionDate string  `db:"transaction_date,omitempty"`
}

func (t Transaction) IsWithdrawal() bool {
	if t.TransactionType != WITHDRAWAL {
		return false
	}
	return true
}

func (t Transaction) ToDTO() dto.TransactionResponse {
	return dto.TransactionResponse{
		TransactionId:   t.TransactionId,
		AccountId:       t.AccountId,
		Amount:          t.Amount,
		TransactionType: t.TransactionType,
		TransactionDate: t.TransactionDate,
	}
}
