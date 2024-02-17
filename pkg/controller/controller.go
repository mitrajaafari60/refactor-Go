package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type CartController struct {
}

const ICE_SESSION_ID = "ice_session_id"

// setNewSessionCookie sets a new session cookie.
func (cc *CartController) SetNewSessionCookie(c *gin.Context) {
	c.SetCookie(ICE_SESSION_ID, time.Now().String(), 3600, "/", "localhost", false, true)
}

// RedirectTo redirects to the page.
func (cc *CartController) RedirectTo(c *gin.Context, page string) {
	c.Redirect(http.StatusFound, page)
}
