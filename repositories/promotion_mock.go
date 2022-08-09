package repositories

import "github.com/stretchr/testify/mock"

type promotionRepositoryMock struct {
	mock.Mock
}

func NewPromotionRepositoryMock() *promotionRepositoryMock {
	return &promotionRepositoryMock{}
}

func (m *promotionRepositoryMock) GetPromotion() (Promotion, error) {
	// mocking progress
	// should return promtion and err
	args := m.Called()
	return args.Get(0).(Promotion), args.Error(1)
}
