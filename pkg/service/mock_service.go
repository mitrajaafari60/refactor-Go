package service

import "interview/pkg/repository"

// MockCartService is a mock implementation of CartServiceInterface.
type MockCartService struct {
	RepoMock repository.Database
}

// NewMockCartService creates a new instance of MockCartService.
func NewMockCartService(repoMock repository.Database) *MockCartService {
	return &MockCartService{
		RepoMock: repoMock,
	}
}

// GetCartData is a mock implementation for testing purposes.
func (m *MockCartService) GetCartData(sessionID string, qErr string) (string, error) {
	// Mock implementation goes here
	return "", nil
}

// AddItemToCart is a mock implementation for testing purposes.
func (m *MockCartService) AddItemToCart(sessionID string, product string, quantity int64) error {
	// Mock implementation goes here
	return nil
}

// DeleteItem is a mock implementation for testing purposes.
func (m *MockCartService) DeleteItem(sessionID string, cartItemID int) error {
	// Mock implementation goes here
	return nil
}

// GetCartItemData is a mock implementation for testing purposes.
func (m *MockCartService) GetCartItemData(sessionID string) (items []map[string]interface{}) {
	// Mock implementation goes here
	return nil
}
