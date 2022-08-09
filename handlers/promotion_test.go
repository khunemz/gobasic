package handlers

import (
	"fmt"
	"gobasic/services"
	"io"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestPromotionCalculateDiscount(t *testing.T) {
	t.Run("Success case", func(t *testing.T) {

		amount := 100
		expected := 80

		promoService := services.NewPromotionServiceMock()
		promoService.On("CalculateDiscount", amount).Return(expected, nil)

		promoHandler := NewPromotionHandler(promoService)
		// test call api
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
