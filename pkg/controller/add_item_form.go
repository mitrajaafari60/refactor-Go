package controller

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (cc *CartController) ShowAddItemForm(c *gin.Context) {
	cookie, err := c.Request.Cookie("ice_session_id")
	if errors.Is(err, http.ErrNoCookie) {
		cc.SetNewSessionCookie(c)
		cookie, err = c.Request.Cookie("ice_session_id")
	}
	sessionId := ""
	if cookie != nil {
		sessionId = cookie.Value
	}
	fmt.Println("sessionId:" + sessionId)
	html, err := cc.CartService.GetCartData(sessionId, c.Query("error"))
	if err != nil {
		c.AbortWithStatus(500)
	}
	c.Header("Content-Type", "text/html")
	c.String(200, html)
}
