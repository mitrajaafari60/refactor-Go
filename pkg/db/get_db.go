package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

func GetDatabase() *gorm.DB {
	// MySQL connection string
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
