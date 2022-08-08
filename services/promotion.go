package services

type PromotionService interface {
	CalculateDiscount(amount int) (int, error)
}

type promotionService struct {
}
