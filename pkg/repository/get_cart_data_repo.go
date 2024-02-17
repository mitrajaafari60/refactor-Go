package repository

import (
	entity2 "interview/pkg/entity"
)

// GetCartData retrieves cart data based on the sessionID.
func (d *MySQLDatabase) GetCartData(sessionID string) ([]map[string]interface{}, error) {
	db := d.GetDatabase()

	var cartEntity entity2.CartEntity
	result := db.Where("status = ? AND session_id = ?", entity2.CartOpen, sessionID).First(&cartEntity)
	if result.Error != nil {
		return nil, result.Error
	}

	var cartItems []entity2.CartItem
	result = db.Where("cart_id = ?", cartEntity.ID).Find(&cartItems)

	if result.Error != nil {
		return nil, result.Error
	}

	var items []map[string]interface{}
	for _, cartItem := range cartItems {
		item := map[string]interface{}{
			"ID":       cartItem.ID,
			"Quantity": cartItem.Quantity,
			"Price":    cartItem.Price,
			"Product":  cartItem.ProductName,
		}

		items = append(items, item)
	}
	return items, nil
}
