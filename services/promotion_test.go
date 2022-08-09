package services_test

import (
	"errors"
	"gobasic/repositories"
	"gobasic/services"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPromotionCalculateDiscount(t *testing.T) {

	type testCases struct {
		name            string
		purchaseMin     int
		discountPercent int
		amount          int
		expected        int
	}

	cases := []testCases{
		{name: "applied 100", purchaseMin: 100, discountPercent: 20, amount: 100, expected: 80},
		{name: "applied 200", purchaseMin: 100, discountPercent: 20, amount: 200, expected: 160},
		{name: "applied 300", purchaseMin: 100, discountPercent: 20, amount: 300, expected: 240},
		{name: "applied 80 (not get promotion discount)", purchaseMin: 100, discountPercent: 20, amount: 100, expected: 80},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			promoRepoMock := repositories.NewPromotionRepositoryMock()
			promoRepoMock.On("GetPromotion").Return(repositories.Promotion{
				ID:              1,
				PurchaseMin:     c.purchaseMin,
				DiscountPercent: c.discountPercent,
			}, nil)
			promoService := services.NewPromotionService(promoRepoMock)

			disc, _ := promoService.CalculateDiscount(c.amount)

			expected := c.expected

			assert.Equal(t, expected, disc)
		})
	}

	t.Run("Purchase amount 0 should return error ", func(t *testing.T) {
		promoRepoMock := repositories.NewPromotionRepositoryMock()
		promoRepoMock.On("GetPromotion").Return(repositories.Promotion{
			ID:              1,
			PurchaseMin:     100,
			DiscountPercent: 20,
		}, nil)
		promoService := services.NewPromotionService(promoRepoMock)

		_, err := promoService.CalculateDiscount(0)

		assert.ErrorIs(t, err, services.ErrZeroAmount)
		promoRepoMock.AssertNotCalled(t, "GetPromotion")

	})

	t.Run("Repository Error", func(t *testing.T) {
		promoRepoMock := repositories.NewPromotionRepositoryMock()
		promoRepoMock.On("GetPromotion").Return(repositories.Promotion{}, errors.New("Unexpected Error"))
		promoService := services.NewPromotionService(promoRepoMock)

		_, err := promoService.CalculateDiscount(100)
		assert.ErrorIs(t, err, services.ErrRepository)
	})
}
