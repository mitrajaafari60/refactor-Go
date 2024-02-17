package controller_test

import (
	"interview/pkg/service"
	"interview/web/controller"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCartController(t *testing.T) {
	// Given
	cartService := &service.CartService{}

	// When
	cartController := controller.NewCartController(cartService)

	// Then
	assert.Equal(t, cartService, cartController.CartService)
}
