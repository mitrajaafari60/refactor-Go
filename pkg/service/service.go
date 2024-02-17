package service

import (
	"interview/pkg/repository"
)

// CartService handles cart-related operations.
type CartService struct {
	repo repository.Database
}

// NewCartService creates a new instance of CartService.
func NewCartService(r repository.Database) CartService {
	return CartService{
		repo: r,
	}
}

// CartServiceInterface is an interface for cart-related operations.
type CartServiceInterface interface {
	GetCartData(sessionID string, qErr string) (string, error)
	AddItemToCart(sessionID string, product string, quantity int64) error
	DeleteItem(sessionID string, cartItemID int) error
	GetCartItemData(sessionID string) (items []map[string]interface{})
}

var itemPriceMapping = map[string]float64{
	"shoe":  100,
	"purse": 200,
	"bag":   300,
	"watch": 300,
}

type CartItemForm struct {
	Product  string `form:"product"   binding:"required"`
	Quantity string `form:"quantity"  binding:"required"`
}
