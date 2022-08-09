package main

import (
	"errors"
	"fmt"

	"github.com/stretchr/testify/mock"
)

func main() {
	fmt.Println("Hello world")
	c := CustomerRepositoryMock{}
	c.On("GetCustomer", 1).Return("Bond", 18, nil)
	c.On("GetCustomer", 2).Return("", 18, errors.New("Not found customer"))
	name, age, err := c.GetCustomer(1)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(name, age)
}

type CustomerRepositoryMock struct {
	mock.Mock
}

func (m CustomerRepositoryMock) GetCustomer(id int) (name string, age int, err error) {
	args := m.Called(id)
	return args.String(0), args.Int(1), args.Error(2)
}
