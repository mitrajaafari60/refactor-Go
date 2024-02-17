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
	// Initialize the application and start the HTTP server
	startServer(setupApplication())
}

func setupApplication() *gin.Engine {
	// load env file
	if err := loadEnvFile(); err != nil {
		log.Fatalf("Error during initialization: %v", err)
	}

	// create a new instance of the cart service
	cartService := service.NewCartService(repository.NewMySQLDatabase())

	// create a new instance of the cart controller
	cartController := controller.NewCartController(&cartService)
	ginEngine := gin.Default()
	// get application routes
	routes := routes(ginEngine, &cartService, &cartController)

	// Migrate database
	repository.NewMySQLDatabase().MigrateDatabase()

	return routes
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

func routes(engine *gin.Engine, cartService *service.CartService, cartController *controller.CartController) *gin.Engine {

	engine.GET("/", cartController.ShowAddItemForm)
	engine.POST("/add-item", cartController.AddItem)
	engine.GET("/remove-cart-item", cartController.DeleteCartItem)

	return engine
}

func startServer(routs *gin.Engine) {
	srv := &http.Server{
		Addr:    ":" + os.Getenv("LISTEN_PORT"),
		Handler: routs,
	}

	err := srv.ListenAndServe()
	if err != nil {
		fmt.Println("server error:", err)
	}
}
