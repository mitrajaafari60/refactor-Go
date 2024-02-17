package main

import (
	"github.com/joho/godotenv"
	"interview/pkg/repository"
	"log"
	"net/http"
	"os"

	"fmt"
	"github.com/gin-gonic/gin"
	"interview/pkg/controller"
	"interview/pkg/service"
	"path/filepath"
	"runtime"
)

var ProjectDir string

func main() {
	// load env file
	if err := loadEnvFile(); err != nil {
		log.Fatalf("Error during initialization: %v", err)
	}
	// get application routes
	routes := routes()

	repository.NewMySQLDatabase().MigrateDatabase()

	srv := &http.Server{
		Addr:    ":" + os.Getenv("LISTEN_PORT"),
		Handler: routes,
	}

	err := srv.ListenAndServe()
	if err != nil {
		fmt.Println("server error:", err)
	}
}

func loadEnvFile() error {
	_, b, _, _ := runtime.Caller(0)

	// Root folder of this project
	ProjectDir = filepath.Join(filepath.Dir(b), "../..")
	// Construct the absolute path to the .env file in the Docker folder
	envFilePath := filepath.Join(ProjectDir, "docker", ".env")

	// Load environment variables from .env file
	if err := godotenv.Load(envFilePath); err != nil {
		return err
	}

	log.Println(".env file loaded successfully")
	return nil
}

func routes() *gin.Engine {
	ginEngine := gin.Default()
	cartService := service.NewCartService(repository.NewMySQLDatabase())
	cartController := controller.NewCartController(&cartService)

	ginEngine.GET("/", cartController.ShowAddItemForm)
	ginEngine.POST("/add-item", cartController.AddItem)
	ginEngine.GET("/remove-cart-item", cartController.DeleteCartItem)

	return ginEngine
}
