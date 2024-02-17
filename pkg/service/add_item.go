package service

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"interview/pkg/entity"
	db2 "interview/pkg/repository"
)

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

func AddItemToCart(sessionID string, product string, quantity int64) error {

	itemPrice, ok := itemPriceMapping[product]
	if !ok {
		return errors.New("invalid item name")
	}
	db := db2.GetDatabase()
	var cartEntity entity.CartEntity
	var cartItemEntity entity.CartItem
	result := db.Where(fmt.Sprintf("status = '%s' AND session_id = '%s'", entity.CartOpen, sessionID)).First(&cartEntity)
	if result.Error != nil {
		if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return result.Error
		}
		cartEntity = entity.CartEntity{
			SessionID: sessionID,
			Status:    entity.CartOpen,
		}
		db.Create(&cartEntity)
		cartItemEntity = entity.CartItem{
			CartID:      cartEntity.ID,
			ProductName: product,
			Quantity:    int(quantity),
			Price:       itemPrice * float64(quantity),
		}
		db.Create(&cartItemEntity)
	} else {
		result = db.Where(" cart_id = ? and product_name  = ?", cartEntity.ID, product).First(&cartItemEntity)

		if result.Error != nil {
			if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
				return result.Error
			}
			cartItemEntity = entity.CartItem{
				CartID:      cartEntity.ID,
				ProductName: product,
				Quantity:    int(quantity),
				Price:       itemPrice * float64(quantity),
			}
			db.Create(&cartItemEntity)
		} else {
			cartItemEntity.Quantity += int(quantity)
			cartItemEntity.Price += itemPrice * float64(quantity)
			db.Save(&cartItemEntity)
		}
	}

	return nil
}
