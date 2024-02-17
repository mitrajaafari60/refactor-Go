package service

func (cs *CartService) DeleteItem(sessionID string, cartItemID int) error {
	return cs.repo.DeleteCartItem(sessionID, cartItemID)
}
