package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestCartEntity(t *testing.T) {
	// Test creating a new CartEntity
	cart := CartEntity{
		Model:     gorm.Model{},
		Total:     100.0,
		SessionID: "test_session",
		Status:    CartOpen,
	}

	// Assertions for the created CartEntity
	assert.Equal(t, float64(100.0), cart.Total)
	assert.Equal(t, "test_session", cart.SessionID)
	assert.Equal(t, CartOpen, cart.Status)

	// Test updating the CartEntity status
	cart.Status = CartClosed

	// Assertions for the updated CartEntity
	assert.Equal(t, CartClosed, cart.Status)
}
