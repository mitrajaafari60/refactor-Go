package repository

import (
	"gorm.io/gorm"
	"interview/pkg/entity"
)

// MockDatabase is a mock implementation of the Database interface for testing purposes.
type MockDatabase struct {
	CartEntities  map[string]*entity.CartEntity
	CartItems     map[uint]*entity.CartItem
	CartItemIndex uint
}

// NewMockDatabase creates a new instance of MockDatabase.
func NewMockDatabase() Database {
	return &MockDatabase{
		CartEntities: make(map[string]*entity.CartEntity),
		CartItems:    make(map[uint]*entity.CartItem),
	}
}

func (d *MockDatabase) GetDatabase() *gorm.DB {
	// Mock implementation, not used in testing
	return nil
}

func (d *MockDatabase) MigrateDatabase() {
	// Mock implementation, not used in testing
}

func (d *MockDatabase) GetOrCreateCart(sessionID string) (*entity.CartEntity, bool, error) {
	cart, exists := d.CartEntities[sessionID]
	if !exists {
		cart = &entity.CartEntity{SessionID: sessionID, Status: entity.CartOpen}
		d.CartEntities[sessionID] = cart
	}
	return cart, !exists, nil
}

func (d *MockDatabase) GetOrCreateCartItem(cartID uint, product string, quantity int64, itemPrice float64) (*entity.CartItem, error) {
	d.CartItemIndex++
	cartItem := &entity.CartItem{
		Model:       gorm.Model{ID: d.CartItemIndex},
		CartID:      cartID,
		ProductName: product,
		Quantity:    int(quantity),
		Price:       itemPrice * float64(quantity),
	}
	d.CartItems[d.CartItemIndex] = cartItem
	return cartItem, nil
}

func (d *MockDatabase) UpdateCartItem(cartItemEntity *entity.CartItem, quantity int64, itemPrice float64) {
	// Mock implementation, not used in testing
}

func (d *MockDatabase) DeleteCartItem(sessionID string, cartItemID int) error {
	// Mock implementation, not used in testing
	return nil
}

func (d *MockDatabase) GetCartData(sessionID string) ([]map[string]interface{}, error) {
	// Mock implementation, not used in testing
	return nil, nil
}
