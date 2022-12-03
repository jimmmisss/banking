package dto

import "github.com/jimmmisss/banking-lib/errs"

const SAVING = "saving"
const CHECKING = "checking"

type NewAccountRequest struct {
	CustomerId  string  `json:"customer_id,omitempty"`
	AccountType string  `json:"account_type,omitempty"`
	Amount      float64 `json:"amount,omitempty"`
}

func (r NewAccountRequest) Validate() *errs.AppError {
	if r.AccountType != SAVING && r.AccountType != CHECKING {
		return errs.NewValidationError("Account type should be checking or saving")
	}
	if r.Amount < 5000 {
		return errs.NewValidationError("To open a new account you need to deposit at least 5000")
	}
	return nil
}
