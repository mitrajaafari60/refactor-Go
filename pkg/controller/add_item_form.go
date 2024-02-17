package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"interview/pkg/service"
	"net/http"
)

func (cc *CartController) ShowAddItemForm(c *gin.Context) {
	_, err := c.Request.Cookie("ice_session_id")
	if errors.Is(err, http.ErrNoCookie) {
		cc.SetNewSessionCookie(c)
	}

	service.GetCartData(c)
}
