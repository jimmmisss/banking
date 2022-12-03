package service

import (
	"github.com/jimmmisss/banking-lib/errs"
	"github.com/jimmmisss/banking/domain"
	"github.com/jimmmisss/banking/dto"
	"time"
)

const dbTSLayout = "2006-01-02 15:04:05"

type AccountService interface {
	NewAccount(dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError)
	MakeTransaction(request dto.TransactionRequest) (*dto.TransactionResponse, *errs.AppError)
}

type DefaultAccountRepository struct {
	repository domain.AccountRepository
}

func NewAccountService(repository domain.AccountRepository) DefaultAccountRepository {
	return DefaultAccountRepository{repository}
}

func (r DefaultAccountRepository) NewAccount(req dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	account := domain.NewAccount(req.CustomerId, req.AccountType, req.Amount)

	if newAccount, err := r.repository.Save(account); err != nil {
		return nil, err
	} else {
		return newAccount.ToNewAccountResponseDTO(), nil
	}
}

func (r DefaultAccountRepository) MakeTransaction(req dto.TransactionRequest) (*dto.TransactionResponse, *errs.AppError) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	if req.IsTransactionTypeWithdrawal() {
		account, err := r.repository.FindById(req.AccountId)
		if err != nil {
			return nil, err
		}
		if !account.CanWithdrawal(req.Amount) {
			return nil, errs.NewValidationError("Insufficient balance in the account")
		}
	}

	t := domain.Transaction{
		AccountId:       req.AccountId,
		Amount:          req.Amount,
		TransactionType: req.TransactionType,
		TransactionDate: time.Now().Format(dbTSLayout),
	}

	transaction, errSaveTransaction := r.repository.SaveTransaction(t)
	if errSaveTransaction != nil {
		return nil, errSaveTransaction
	}

	response := transaction.ToDTO()
	return &response, nil

}
