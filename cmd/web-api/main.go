package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"interview/pkg/controllers"
	"interview/pkg/db"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func loadEnvFile() error {
	// Construct the absolute path to the .env file in the Docker folder
	envFilePath := filepath.Join("../..", "docker", ".env")

	// Load environment variables from .env file
	if err := godotenv.Load(envFilePath); err != nil {
		return err
	}

	log.Println(".env file loaded successfully")
	return nil
}

func main() {
	if err := loadEnvFile(); err != nil {
		log.Fatalf("Error during initialization: %v", err)
	}

	db.MigrateDatabase()

	ginEngine := gin.Default()

	var cartController controllers.CartController
	ginEngine.GET("/", cartController.ShowAddItemForm)
	ginEngine.POST("/add-item", cartController.AddItem)
	ginEngine.GET("/remove-cart-item", cartController.DeleteCartItem)
	srv := &http.Server{
		Addr:    ":" + os.Getenv("LISTEN_PORT"),
		Handler: ginEngine,
	}

	srv.ListenAndServe()
}
