package repository

import (
	"errors"
	"gorm.io/gorm"
	"interview/pkg/entity"
)

// GetOrCreateCart retrieves or creates a cart based on the sessionID.
func (d *MySQLDatabase) GetOrCreateCart(sessionID string) (*entity.CartEntity, bool, error) {
	db := d.GetDatabase()

	var cartEntity entity.CartEntity
	result := db.Where("status = ? AND session_id = ?", entity.CartOpen, sessionID).First(&cartEntity)

	if result.Error != nil {
		if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, false, result.Error
		}

		newCartEntity := entity.CartEntity{
			SessionID: sessionID,
			Status:    entity.CartOpen,
		}

		if err := db.Create(&newCartEntity).Error; err != nil {
			return nil, false, err
		}

		return &newCartEntity, true, nil
	}

	return &cartEntity, false, nil
}
