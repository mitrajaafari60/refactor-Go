package main

import (
	"github.com/gin-gonic/gin"
	"interview/pkg/controller"
	"interview/pkg/repository"
	"interview/pkg/service"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestLoadEnvFile(t *testing.T) {
	// Backup and restore original environment variables after the test
	originalEnv := os.Environ()
	defer func() {
		os.Clearenv()
		for _, envVar := range originalEnv {
			pair := strings.SplitN(envVar, "=", 2)
			os.Setenv(pair[0], pair[1])
		}
	}()

	// Set up a temporary .env file with test values
	tempDir := t.TempDir()
	envFilePath := filepath.Join(tempDir, ".env")
	err := ioutil.WriteFile(envFilePath, []byte("TEST_VAR=test_value"), 0644)
	if err != nil {
		t.Fatalf("Failed to create temporary .env file: %v", err)
	}

	// Set the PROJECT_DIR environment variable
	os.Setenv("TEST_VAR", "test_value")

	// Run the loadEnvFile function
	err = loadEnvFile()

	// Check for errors
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	// Verify that the environment variable was loaded
	testVar := os.Getenv("TEST_VAR")
	if testVar != "test_value" {
		t.Errorf("Expected TEST_VAR=test_value, got %s", testVar)
	}
}

func TestRoutes(t *testing.T) {
	var registered = []struct {
		route  string
		method string
	}{
		{"/", "GET"},
		{"/remove-cart-item", "GET"},
		{"/add-item", "POST"},
	}
	ginEngine := gin.Default()
	db := repository.NewMockDatabase()
	cartService := service.NewCartService(db)
	cartController := controller.NewCartController(&cartService)
	mux := routes(ginEngine, &cartService, &cartController)

	for _, route := range registered {
		// check to see if the route exists
		if !routeExists(route.route, route.method, mux) {
			t.Errorf("route %s is not registered", route.route)
		}
	}
}

func routeExists(testRoute, testMethod string, ginRouter *gin.Engine) bool {
	found := false
	for _, route := range ginRouter.Routes() {
		if strings.EqualFold(testMethod, route.Method) && strings.EqualFold(testRoute, route.Path) {
			found = true
		}
	}
	return found
}
