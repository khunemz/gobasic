package service

import (
	"database/sql"
	"errors"
	"go-hexagonal/repository"
)

type customerService struct {
	custeRepo repository.CustomerRepository
}

func NewCustomerService(custRepo repository.CustomerRepository) customerService {
	return customerService{custeRepo: custRepo}
}

func (s customerService) GetCustomers() ([]CustomerResponse, error) {
	customers, err := s.custeRepo.GetAll()
	if err != nil {
		return nil, err
	}

	custResList := []CustomerResponse{}

	for _, customer := range customers {
		custRes := CustomerResponse{
			CustomerID: customer.CustomerId,
			Name:       customer.Name,
			Status:     customer.Status,
		}

		custResList = append(custResList, custRes)
	}
	return custResList, nil
}

func (s customerService) GetCustomer(id int) (*CustomerResponse, error) {
	customer, err := s.custeRepo.GetById(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("Entity not found")
		}
		return nil, err
	}

	custRes := CustomerResponse{
		CustomerID: customer.CustomerId,
		Name:       customer.Name,
		Status:     customer.Status,
	}
	return &custRes, nil
}
