package service

import (
	"github.com/stretchr/testify/assert"
	"interview/pkg/repository"
	"testing"
)

func TestCartService_GetCartData(t *testing.T) {
	// Setup
	mockRepo := repository.NewMockDatabase()
	cartService := NewCartService(mockRepo)
	data, err := cartService.GetCartData("sessionID", "")
	assert.NoError(t, err)
	assert.NotEmpty(t, data)
}

func TestCartService_AddItemToCart(t *testing.T) {
	// Setup
	mockRepo := repository.NewMockDatabase()
	cartService := NewCartService(mockRepo)

	// Test case 1: Valid item
	err := cartService.AddItemToCart("sessionID", "shoe", 2)
	assert.NoError(t, err)

	// Test case 2: Invalid item
	err = cartService.AddItemToCart("sessionID", "invalidItem", 2)
	assert.Error(t, err)
}

func TestCartService_DeleteItem(t *testing.T) {
	// Setup
	mockRepo := repository.NewMockDatabase()
	cartService := NewCartService(mockRepo)

	// Test case: Successful deletion
	err := cartService.DeleteItem("sessionID", 123)
	assert.NoError(t, err)
}

func TestCartService_GetCartItemData(t *testing.T) {
	// Setup
	mockRepo := repository.NewMockDatabase()
	cartService := NewCartService(mockRepo)

	// Test case: Valid cart item data retrieval
	items := cartService.GetCartItemData("sessionID")
	assert.NotNil(t, items)
}
