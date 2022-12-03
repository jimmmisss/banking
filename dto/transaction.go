package dto

import "github.com/jimmmisss/banking-lib/errs"

const WITHDRAWAL = "withdrawal"
const DEPOSIT = "deposit"

type TransactionRequest struct {
	AccountId       string  `json:"account_id,omitempty"`
	Amount          float64 `json:"amount,omitempty"`
	TransactionType string  `json:"transaction_type,omitempty"`
	TransactionDate string  `json:"transaction_date,omitempty"`
	CustomerId      string  `json:"customer_id,omitempty"`
}

func (r TransactionRequest) IsTransactionTypeWithdrawal() bool {
	if r.TransactionType == WITHDRAWAL {
		return true
	}
	return false
}

func (r TransactionRequest) Validate() *errs.AppError {
	if r.TransactionType != WITHDRAWAL && r.TransactionType != DEPOSIT {
		return errs.NewValidationError("Transaction type can only be deposit or withdrawal")
	}
	if r.Amount < 0 {
		return errs.NewValidationError("Amount cannot be less than zero")
	}
	return nil
}

type TransactionResponse struct {
	TransactionId   string  `json:"transaction_id,omitempty"`
	AccountId       string  `json:"account_id,omitempty"`
	Amount          float64 `json:"amount,omitempty"`
	TransactionType string  `json:"transaction_type,omitempty"`
	TransactionDate string  `json:"transaction_date,omitempty"`
}
