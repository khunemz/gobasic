//go:build integration

package handlers_test

import (
	"fmt"
	"gobasic/handlers"
	"gobasic/repositories"
	"gobasic/services"
	"io"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestPromotionCalculateDiscIntegrationService(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		amount := 100
		expected := 80
		promoRepo := repositories.NewPromotionRepositoryMock()
		promoRepo.On("GetPromotion").Return(repositories.Promotion{
			ID:              1,
			PurchaseMin:     100,
			DiscountPercent: 20,
		}, nil)

		promoService := services.NewPromotionService(promoRepo)
		promoHandler := handlers.NewPromotionHandler(promoService)

		app := fiber.New()
		app.Get("/calculate", promoHandler.CalculateDiscount)
		req := httptest.NewRequest("GET", fmt.Sprintf("/calculate?amount=%v", amount), nil)

		res, _ := app.Test(req)
		body, _ := io.ReadAll(res.Body)
		ok := assert.Equal(t, fiber.StatusOK, res.StatusCode)
		if ok {
			assert.Equal(t, strconv.Itoa(expected), string(body))
		}

	})
}
