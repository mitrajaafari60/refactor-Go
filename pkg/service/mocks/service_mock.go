package mocks

import (
	"github.com/stretchr/testify/mock"
	"interview/pkg/entity"
)

type CartServiceInterface interface {
	GetCartData(sessionId string) (*entity.CartEntity, error)
	AddItem(sessionId string, item *entity.CartItem) error
	UpdateCartItem(sessionId string, item *entity.CartItem) error
	DeleteCartItem(sessionId string, id int) error
	CreateCart(sessionId string) (*entity.CartEntity, error)
}

type CartService struct {
	mock.Mock
}

func (s *CartService) GetCartData(sessionId string) (*entity.CartEntity, error) {
	args := s.Called(sessionId)
	return args.Get(0).(*entity.CartEntity), args.Error(1)
}

func (s *CartService) AddItem(sessionId string, item *entity.CartItem) error {
	args := s.Called(sessionId, item)
	return args.Error(0)
}

func (s *CartService) UpdateCartItem(sessionId string, item *entity.CartItem) error {
	args := s.Called(sessionId, item)
	return args.Error(0)
}

func (s *CartService) DeleteCartItem(sessionId string, id int) error {
	args := s.Called(sessionId, id)
	return args.Error(0)
}

func (s *CartService) CreateCart(sessionId string) (*entity.CartEntity, error) {
	args := s.Called(sessionId)
	return args.Get(0).(*entity.CartEntity), args.Error(1)
}
