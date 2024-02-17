package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"interview/pkg/service"
	"net/http"
)

func (cc *CartController) ShowAddItemForm(c *gin.Context) {
	cookie, err := c.Request.Cookie(ICE_SESSION_ID)
	if errors.Is(err, http.ErrNoCookie) {
		cc.SetNewSessionCookie(c)
		cookie, err = c.Request.Cookie(ICE_SESSION_ID)
	}
	sessionId := ""
	if cookie != nil {
		sessionId = cookie.Value
	}
	html, err := service.GetCartData(sessionId, c.Query("error"))
	if err != nil {
		c.AbortWithStatus(500)
	}
	c.Header("Content-Type", "text/html")
	c.String(200, html)
}
