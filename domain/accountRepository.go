package domain

import (
	"github.com/jimmmisss/banking-lib/errs"
	"github.com/jimmmisss/banking-lib/logger"
	"github.com/jmoiron/sqlx"
	"strconv"
)

type AccountRepositoryDB struct {
	client *sqlx.DB
}

func NewAccountRepositoryDB(dbClient *sqlx.DB) AccountRepositoryDB {
	return AccountRepositoryDB{dbClient}
}

func (r AccountRepositoryDB) Save(a Account) (*Account, *errs.AppError) {
	sql := "INSERT INTO accounts(customer_id, opening_date, account_type, amount, status) values (?, ?, ?, ?, ?)"

	result, err := r.client.Exec(sql, a.CustomerId, a.OpeningDate, a.AccountType, a.Amount, a.Status)
	if err != nil {
		logger.Info("Error while creating new account" + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from database")
	}

	id, err := result.LastInsertId()
	if err != nil {
		logger.Info("Error while getting last insert id for new account" + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from database")
	}

	a.AccountId = strconv.FormatInt(id, 10)
	return &a, nil
}

func (r AccountRepositoryDB) SaveTransaction(t Transaction) (*Transaction, *errs.AppError) {
	tx, err := r.client.Begin()
	if err != nil {
		logger.Error("Error while starting a new transaction for bank account transaction: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	result, _ := tx.Exec(`INSERT INTO transactions(account_id, amount, transaction_type, transaction_date) values (?, ?, ?, ?)`, t.AccountId, t.Amount, t.TransactionType, t.TransactionDate)

	if t.IsWithdrawal() {
		_, err = tx.Exec(`UPDATE accounts SET amount = amount - ? where account_id = ?`, t.Amount, t.AccountId)
	} else {
		_, err = tx.Exec(`UPDATE accounts SET amount = amount + ? where account_id = ?`, t.Amount, t.AccountId)
	}

	if err != nil {
		errRoll := tx.Rollback()
		if errRoll != nil {
			logger.Error("Error rollback insert: " + errRoll.Error())
			return nil, errs.NewUnexpectedError("Unexpected database error")
		}
		logger.Error("Error while saving transaction: " + errRoll.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	err = tx.Commit()
	if err != nil {
		errRoll := tx.Rollback()
		if errRoll != nil {
			logger.Error("Error rollback insert: " + errRoll.Error())
			return nil, errs.NewUnexpectedError("Unexpected database error")
		}
		logger.Error("Error while committing transaction for bank account: " + errRoll.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	transactionId, errLastInsertId := result.LastInsertId()
	if errLastInsertId != nil {
		logger.Error("Error while getting the last transaction id: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	account, errFindById := r.FindById(t.AccountId)
	if errFindById != nil {
		return nil, errFindById
	}

	t.TransactionId = strconv.FormatInt(transactionId, 10)
	t.Amount = account.Amount
	return &t, nil
}

func (r AccountRepositoryDB) FindById(accountId string) (*Account, *errs.AppError) {
	sql := "SELECT account_id, customer_id, opening_date, account_type, amount from accounts where account_id = ?"
	var account Account
	err := r.client.Get(&account, sql, accountId)
	if err != nil {
		logger.Error("Error while fetching account information: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}
	return &account, nil
}
