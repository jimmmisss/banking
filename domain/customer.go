package domain

import (
	"github.com/jimmmisss/banking-lib/errs"
	"github.com/jimmmisss/banking/dto"
)

type Customer struct {
	Id          string `db:"customer_id,omitempty"`
	Name        string `db:"name,omitempty"`
	City        string `db:"city,omitempty"`
	Zipcode     string `db:"zipcode,omitempty"`
	DateOfBirth string `db:"date_of_birth,omitempty"`
	Status      string `db:"status,omitempty"`
}

func (c Customer) statusAsText() string {
	statusText := "active"
	if c.Status == "0" {
		statusText = "inactive"
	}
	return statusText
}

func (c Customer) ToDto() dto.CustomerResponse {
	return dto.CustomerResponse{
		Id:          c.Id,
		Name:        c.Name,
		City:        c.City,
		Zipcode:     c.Zipcode,
		DateOfBirth: c.DateOfBirth,
		Status:      c.statusAsText(),
	}
}

type CustomerRepository interface {
	FindAll(status string) ([]Customer, *errs.AppError)
	FindById(string string) (*Customer, *errs.AppError)
}
