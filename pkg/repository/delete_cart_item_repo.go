package repository

import (
	"errors"
	entity2 "interview/pkg/entity"
)

// DeleteCartItem deletes a cart item based on the sessionID and cartItemID.
func (d *MySQLDatabase) DeleteCartItem(sessionID string, cartItemID int) error {
	db := d.GetDatabase()

	var cartEntity entity2.CartEntity
	result := db.Where("status = ? AND session_id = ?", entity2.CartOpen, sessionID).First(&cartEntity)
	if result.Error != nil {
		return result.Error
	}

	if cartEntity.Status == entity2.CartClosed {
		return errors.New("CartClosed")
	}

	var cartItemEntity entity2.CartItem

	result = db.Where(" ID  = ?", cartItemID).First(&cartItemEntity)
	if result.Error != nil {
		return result.Error
	}

	db.Delete(&cartItemEntity)
	return nil
}
