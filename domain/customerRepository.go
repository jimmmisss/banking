package domain

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jimmmisss/banking-lib/errs"
	"github.com/jimmmisss/banking-lib/logger"
	"github.com/jmoiron/sqlx"
)

type CustomerRepositoryDB struct {
	client *sqlx.DB
}

func (db CustomerRepositoryDB) FindAll(status string) ([]Customer, *errs.AppError) {
	var err error
	customers := make([]Customer, 0)

	if status == "" {
		query := "select customer_id, name, city, zipcode, date_of_birth, status from customers"
		err = db.client.Select(&customers, query)
	} else {
		query := "select customer_id, name, city, zipcode, date_of_birth, status from customers where status = ?"
		err = db.client.Select(&customers, query, status)
	}

	if err != nil {
		logger.Error("Error while querying customer table " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	return customers, nil
}

func (db CustomerRepositoryDB) FindById(id string) (*Customer, *errs.AppError) {
	query := "select customer_id, name, city, zipcode, date_of_birth, status from customers where customer_id = ?"

	var c Customer
	err := db.client.Get(&c, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("Customer not found")
		} else {
			logger.Error("Error while scanning customer " + err.Error())
			return nil, errs.NewUnexpectedError("Unexpected database error")
		}
	}
	return &c, nil
}

func NewCustomerRepositoryDB(dbClient *sqlx.DB) CustomerRepositoryDB {
	return CustomerRepositoryDB{dbClient}
}
