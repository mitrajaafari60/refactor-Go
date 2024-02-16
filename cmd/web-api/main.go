package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"interview/pkg/controller"
	"interview/pkg/db"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
)

func loadEnvFile() error {
	_, b, _, _ := runtime.Caller(0)

	// Root folder of this project
	rootProject := filepath.Join(filepath.Dir(b), "../..")
	// Construct the absolute path to the .env file in the Docker folder
	envFilePath := filepath.Join(rootProject, "docker", ".env")

	// Load environment variables from .env file
	if err := godotenv.Load(envFilePath); err != nil {
		return err
	}

	log.Println(".env file loaded successfully")
	return nil
}

func routes() *gin.Engine {
	ginEngine := gin.Default()

	var cartController controller.CartController
	ginEngine.GET("/", cartController.ShowAddItemForm)
	ginEngine.POST("/add-item", cartController.AddItem)
	ginEngine.GET("/remove-cart-item", cartController.DeleteCartItem)

	return ginEngine
}

func main() {
	// load env file
	if err := loadEnvFile(); err != nil {
		log.Fatalf("Error during initialization: %v", err)
	}
	// get application routes
	routes := routes()

	db.MigrateDatabase()

	srv := &http.Server{
		Addr:    ":" + os.Getenv("LISTEN_PORT"),
		Handler: routes,
	}

	srv.ListenAndServe()
}
