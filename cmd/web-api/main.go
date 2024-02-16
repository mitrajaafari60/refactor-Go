package main

import (
	"github.com/gin-gonic/gin"
	"interview/pkg/controllers"
	"interview/pkg/db"
	"net/http"
)

func main() {
	db.MigrateDatabase()

	ginEngine := gin.Default()

	var cartController controllers.CartController
	ginEngine.GET("/", cartController.ShowAddItemForm)
	ginEngine.POST("/add-item", cartController.AddItem)
	ginEngine.GET("/remove-cart-item", cartController.DeleteCartItem)
	srv := &http.Server{
		Addr:    ":8088",
		Handler: ginEngine,
	}

	srv.ListenAndServe()
}
