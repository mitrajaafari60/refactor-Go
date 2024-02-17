package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"interview/pkg/service"
	"net/http"
)

func (cc *CartController) AddItem(c *gin.Context) {
	cookie, err := c.Request.Cookie("ice_session_id")

	if err != nil || errors.Is(err, http.ErrNoCookie) || (cookie != nil && cookie.Value == "") {
		cc.RedirectTo(c, "/")
		return
	}

	service.AddItemToCart(c)
}
