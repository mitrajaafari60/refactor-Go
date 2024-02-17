package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type CartController struct {
}

// setNewSessionCookie sets a new session cookie.
func (cc *CartController) SetNewSessionCookie(c *gin.Context) {
	c.SetCookie("ice_session_id", time.Now().String(), 3600, "/", "localhost", false, true)
}

// RedirectTo redirects to the page.
func (cc *CartController) RedirectTo(c *gin.Context, page string) {
	c.Redirect(http.StatusFound, page)
}
