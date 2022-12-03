package domain

import (
	"github.com/jimmmisss/banking-lib/errs"
	"github.com/jimmmisss/banking/dto"
)

const dbTSLayout = "2006-01-02 15:04:05"

type Account struct {
	AccountId   string  `db:"account_id,omitempty"`
	CustomerId  string  `db:"customer_id,omitempty"`
	OpeningDate string  `db:"opening_date,omitempty"`
	AccountType string  `db:"account_type,omitempty"`
	Amount      float64 `db:"amount,omitempty"`
	Status      string  `db:"status,omitempty"`
}

func (a Account) ToNewAccountResponseDTO() *dto.NewAccountResponse {
	return &dto.NewAccountResponse{AccountId: a.AccountId}
}

//go:generate mockgen -destination=../mocks/domain/mockAccountRepository.go -package=domain github.com/jimmmisss/banking/domain AccountRepository
type AccountRepository interface {
	Save(account Account) (*Account, *errs.AppError)
	SaveTransaction(transaction Transaction) (*Transaction, *errs.AppError)
	FindById(accountId string) (*Account, *errs.AppError)
}

func (a Account) CanWithdrawal(amount float64) bool {
	if a.Amount < amount {
		return false
	}
	return true
}

func NewAccount(customerId, accountType string, amount float64) Account {
	return Account{
		CustomerId:  customerId,
		OpeningDate: dbTSLayout,
		AccountType: accountType,
		Amount:      amount,
		Status:      "1",
	}
}
