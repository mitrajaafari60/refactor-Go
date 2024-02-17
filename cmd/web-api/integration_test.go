package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntegration(t *testing.T) {
	ginEngine := setupApplication()

	// Simulate a GET request to the "/" endpoint
	req1, err := http.NewRequest("GET", "/", nil)
	assert.NoError(t, err)

	w1 := httptest.NewRecorder()
	ginEngine.ServeHTTP(w1, req1)

	// Check if the response status code is http.StatusOK
	assert.Equal(t, http.StatusOK, w1.Code)

	// Simulate a POST request to the "/add-item" endpoint
	formData := "product=shoe&quantity=2"
	req2, err := http.NewRequest("POST", "/add-item", bytes.NewBufferString(formData))
	assert.NoError(t, err)

	req2.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	w2 := httptest.NewRecorder()
	ginEngine.ServeHTTP(w2, req2)

	// Check if the response status code is http.StatusFound
	assert.Equal(t, http.StatusFound, w2.Code)

	// Simulate a GET request to the "/remove-cart-item" endpoint
	req3, err := http.NewRequest("GET", "/remove-cart-item?cart_item_id=123", nil)
	assert.NoError(t, err)

	w3 := httptest.NewRecorder()
	ginEngine.ServeHTTP(w3, req3)

	// Check if the response status code is http.StatusFound
	assert.Equal(t, http.StatusFound, w3.Code)
}
