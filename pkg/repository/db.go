package repository

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"interview/pkg/entity"
	"os"
)

// Database is an interface for database operations.
type Database interface {
	MigrateDatabase()
	GetDatabase() *gorm.DB
	GetOrCreateCart(sessionID string) (*entity.CartEntity, bool, error)
	GetOrCreateCartItem(cartID uint, product string, quantity int64, itemPrice float64) (*entity.CartItem, bool, error)
	UpdateCartItem(cartItemEntity *entity.CartItem, quantity int64, itemPrice float64)
	DeleteCartItem(sessionID string, cartItemID int) error
	GetCartData(sessionID string) ([]map[string]interface{}, error)
}

// MySQLDatabase is an implementation of the Database interface.
type MySQLDatabase struct {
}

// NewMySQLDatabase creates a new instance of MySQLDatabase.
func NewMySQLDatabase() Database {
	return &MySQLDatabase{}
}

func (d *MySQLDatabase) GetDatabase() *gorm.DB {
	// Retrieve MySQL connection details from environment variables
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	// Construct MySQL connection string
	dsn := username + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"

	// Open the connection to the database
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	return db
}

func (d *MySQLDatabase) MigrateDatabase() {
	db := d.GetDatabase()

	// AutoMigrate will create or update the tables based on the models
	err := db.AutoMigrate(&entity.CartEntity{}, &entity.CartItem{})
	if err != nil {
		panic(err)
	}
}
