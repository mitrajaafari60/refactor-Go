package service

import (
	"errors"
)

// AddItemToCart implements the AddItemToCart method of the CartServiceInterface.
func (cs *CartService) AddItemToCart(sessionID string, product string, quantity int64) error {
	cartEntity, _, err := cs.repo.GetOrCreateCart(sessionID)
	if err != nil {
		return err
	}

	itemPrice, ok := itemPriceMapping[product]
	if !ok {
		return errors.New("invalid item name")
	}

	cartItemEntity, isNew, err := cs.repo.GetOrCreateCartItem(cartEntity.ID, product, quantity, itemPrice)
	if err != nil {
		return err
	}
	if !isNew {
		cs.repo.UpdateCartItem(cartItemEntity, quantity, itemPrice)
	}

	return nil
}
