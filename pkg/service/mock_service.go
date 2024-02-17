package service

import (
	"errors"
)

func NewMockCartService() CartServiceInterface {
	return &MockCartService{}
}

type MockCartService struct {
}

func (s *MockCartService) GetCartData(sessionID string, qErr string) (string, error) {
	data := map[string]interface{}{
		"Error":     qErr,
		"CartItems": s.GetCartItemData(sessionID),
	}
	return renderTemplate(data)
}

func (s *MockCartService) AddItemToCart(sessionID string, product string, quantity int64) error {
	_, ok := itemPriceMapping[product]
	if !ok {
		return errors.New("invalid item name")
	}
	return nil
}

func (s *MockCartService) DeleteItem(sessionID string, cartItemID int) error {
	return nil
}

func (s *MockCartService) GetCartItemData(sessionID string) (items []map[string]interface{}) {

	return nil
}
