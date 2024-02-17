package repository

import (
	"errors"
	"gorm.io/gorm"
	"interview/pkg/entity"
)

// GetOrCreateCartItem retrieves or creates a cart item based on the cartID and product.
func (d *MySQLDatabase) GetOrCreateCartItem(cartID uint, product string, quantity int64, itemPrice float64) (*entity.CartItem, error) {
	db := d.GetDatabase()

	var cartItemEntity entity.CartItem
	result := db.Where("cart_id = ? and product_name = ?", cartID, product).First(&cartItemEntity)

	if result.Error != nil {
		if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, result.Error
		}

		newCartItemEntity := entity.CartItem{
			CartID:      cartID,
			ProductName: product,
			Quantity:    int(quantity),
			Price:       itemPrice * float64(quantity),
		}

		if err := db.Create(&newCartItemEntity).Error; err != nil {
			return nil, err
		}

		return &newCartItemEntity, nil
	}

	return &cartItemEntity, nil
}
