package repository

import (
	"interview/pkg/entity"
)

// UpdateCartItem updates the quantity and price of a cart item.
func (d *MySQLDatabase) UpdateCartItem(cartItemEntity *entity.CartItem, quantity int64, itemPrice float64) {
	db := d.GetDatabase()

	cartItemEntity.Quantity += int(quantity)
	cartItemEntity.Price += itemPrice * float64(quantity)
	db.Save(cartItemEntity)
}
