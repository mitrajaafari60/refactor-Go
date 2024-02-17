package service

import (
	"errors"
	"fmt"
	"interview/pkg/entity"
	db2 "interview/pkg/repository"
)

func DeleteCartItem(sessionID string, cartItemID int) error {

	db := db2.GetDatabase()

	var cartEntity entity.CartEntity
	result := db.Where(fmt.Sprintf("status = '%s' AND session_id = '%s'", entity.CartOpen, sessionID)).First(&cartEntity)
	if result.Error != nil {
		return result.Error
	}

	if cartEntity.Status == entity.CartClosed {
		return errors.New("CartClosed")
	}

	var cartItemEntity entity.CartItem

	result = db.Where(" ID  = ?", cartItemID).First(&cartItemEntity)
	if result.Error != nil {
		return result.Error
	}

	db.Delete(&cartItemEntity)
	return nil
}
