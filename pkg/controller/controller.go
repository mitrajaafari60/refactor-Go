package controller

import (
	"github.com/gin-gonic/gin"
	"interview/pkg/service"
	"net/http"
	"time"
)

const ICE_SESSION_ID = "ice_session_id"

// CartControllerInterface is an interface for shopping cart operations.
type CartControllerInterface interface {
	ShowAddItemForm(c *gin.Context)
	AddItem(c *gin.Context)
	DeleteCartItem(c *gin.Context)
	SetNewSessionCookie(c *gin.Context)
	RedirectTo(c *gin.Context, page string)
}

// CartController is an implementation of the ShoppingCartControllerInterface.
type CartController struct {
	CartService service.CartServiceInterface // Inject the cart service interface
}

func NewCartController(cs *service.CartService) CartController {
	return CartController{
		CartService: cs,
	}
}

// setNewSessionCookie sets a new session cookie.
func (cc *CartController) SetNewSessionCookie(c *gin.Context) {
	c.SetCookie(ICE_SESSION_ID, time.Now().String(), 3600, "/", "localhost", false, true)
}

// RedirectTo redirects to the page.
func (cc *CartController) RedirectTo(c *gin.Context, page string) {
	c.Redirect(http.StatusFound, page)
}
