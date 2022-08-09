package handlers

import (
	"gobasic/services"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type PromotionHandler interface {
	CalculateDiscount(c *fiber.Ctx) error
}
type promotionHander struct {
	promoService services.PromotionService
}

func NewPromotionHandler(promoSrv services.PromotionService) promotionHander {
	return promotionHander{promoService: promoSrv}
}

func (h promotionHander) CalculateDiscount(c *fiber.Ctx) error {
	amountStr := c.Query("amount")
	amount, err := strconv.Atoi(amountStr)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	disc, err := h.promoService.CalculateDiscount(amount)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	return c.SendString(strconv.Itoa(disc))
}
