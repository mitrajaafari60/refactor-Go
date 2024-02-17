package repository

import (
	"github.com/stretchr/testify/assert"
	"interview/pkg/entity"
	"os"
	"testing"
)

func setupTestEnvVariables() {
	os.Setenv("DB_USERNAME", "ice_user")

	os.Setenv("DB_PASSWORD", "9xz3jrd8wf")
	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_PORT", "4001")
	os.Setenv("DB_NAME", "ice_db")
}
func TestNewMySQLDatabase(t *testing.T) {
	// Set up environment variables for testing database connection
	setupTestEnvVariables()
	// Create an instance of MySQLDatabase
	db := NewMySQLDatabase()

	// Check if the returned database instance is not nil
	assert.NotNil(t, db)
}

func TestMySQLDatabase_GetDatabase(t *testing.T) {
	// Set up environment variables for testing database connection
	setupTestEnvVariables()
	// Create an instance of MySQLDatabase
	db := NewMySQLDatabase()

	// Call GetDatabase method
	gormDB := db.GetDatabase()

	// Check if the returned gorm.DB instance is not nil
	assert.NotNil(t, gormDB)
}

func TestMySQLDatabase_MigrateDatabase(t *testing.T) {
	// Set up environment variables for testing database connection
	setupTestEnvVariables()
	// Create an instance of MySQLDatabase
	db := NewMySQLDatabase()

	// Invoke MigrateDatabase method
	db.MigrateDatabase()

	// Check if tables are created by attempting to query them
	var cartEntity entity.CartEntity
	var cartItem entity.CartItem
	result1 := db.GetDatabase().Migrator().HasTable(&cartEntity)
	result2 := db.GetDatabase().Migrator().HasTable(&cartItem)

	// Assertions
	assert.True(t, result1, "CartEntity table should exist")
	assert.True(t, result2, "CartItem table should exist")
}

func TestMySQLDatabase_GetOrCreateCart(t *testing.T) {
	// Set up environment variables for testing database connection
	setupTestEnvVariables()

	// Create an instance of MySQLDatabase
	db := NewMySQLDatabase()

	// Call GetOrCreateCart method
	sessionID := "test_session"
	cart, isNew, err := db.GetOrCreateCart(sessionID)
	cart, isNew, err = db.GetOrCreateCart(sessionID)

	// Assertions
	assert.NoError(t, err)
	assert.NotNil(t, cart)
	assert.False(t, isNew)
}

func TestMySQLDatabase_GetOrCreateCartItem(t *testing.T) {
	// Set up environment variables for testing database connection
	setupTestEnvVariables()

	// Create an instance of MySQLDatabase
	db := NewMySQLDatabase()

	// Mock data for testing
	cartID := uint(1)
	product := "test_product"
	quantity := int64(2)
	itemPrice := 50.0

	// Call GetOrCreateCartItem method
	cartItem, isNew, err := db.GetOrCreateCartItem(cartID, product, quantity, itemPrice)

	// Assertions
	assert.NoError(t, err)
	assert.NotNil(t, cartItem)
	if isNew {
		_, isNew, _ = db.GetOrCreateCartItem(cartID, product, quantity, itemPrice)
	}
	assert.False(t, isNew) // a new cart item is created

	// Call GetOrCreateCartItem method again to simulate retrieving an existing cart item
	cartItem, isNew, err = db.GetOrCreateCartItem(cartID, product, quantity, itemPrice)
	if isNew {
		cartItem, isNew, err = db.GetOrCreateCartItem(cartID, product, quantity, itemPrice)
	}
	// Assertions
	assert.NoError(t, err)
	assert.NotNil(t, cartItem)
	assert.False(t, isNew) // if isNew  run again so Assuming the cart item already exists
}

func TestMySQLDatabase_UpdateCartItem(t *testing.T) {
	// Set up environment variables for testing database connection
	setupTestEnvVariables()

	// Create an instance of MySQLDatabase
	db := NewMySQLDatabase()

	// Mock data for testing
	cartItem := &entity.CartItem{
		// Initialize with required fields for testing
	}

	// Call UpdateCartItem method
	db.UpdateCartItem(cartItem, 5, 60.0)
}
