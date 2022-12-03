package service

import (
	"github.com/jimmmisss/banking-lib/errs"
	"github.com/jimmmisss/banking/domain"
	"github.com/jimmmisss/banking/dto"
)

//go:generate mockgen -destination=../mocks/service/mockCustomerService.go -package=service github.com/jimmmisss/banking/service CustomerService
type CustomerService interface {
	GetAllCustomer(status string) ([]dto.CustomerResponse, *errs.AppError)
	GetByIdCustomer(string) (*dto.CustomerResponse, *errs.AppError)
}

type DefaultCustomerRepository struct {
	repository domain.CustomerRepository
}

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerRepository {
	return DefaultCustomerRepository{repository: repository}
}

func (r DefaultCustomerRepository) GetAllCustomer(status string) ([]dto.CustomerResponse, *errs.AppError) {
	if status == "active" {
		status = "1"
	} else if status == "inactive" {
		status = "0"
	} else {
		status = ""
	}

	customers, err := r.repository.FindAll(status)
	if err != nil {
		return nil, err
	}

	response := make([]dto.CustomerResponse, 0)
	for _, c := range customers {
		response = append(response, c.ToDto())
	}
	return response, nil
}

func (r DefaultCustomerRepository) GetByIdCustomer(id string) (*dto.CustomerResponse, *errs.AppError) {
	c, err := r.repository.FindById(id)
	if err != nil {
		return nil, err
	}
	response := c.ToDto()
	return &response, nil
}
