package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (cc *CartController) DeleteCartItem(c *gin.Context) {
	cookie, err := c.Request.Cookie(ICE_SESSION_ID)

	if err != nil || errors.Is(err, http.ErrNoCookie) || (cookie != nil && cookie.Value == "") {
		cc.RedirectTo(c, "/")
		return
	}
	cartItemIDString := c.Query("cart_item_id")
	if cartItemIDString == "" {
		cc.RedirectTo(c, "/")
		return
	}

	cartItemID, err := strconv.Atoi(cartItemIDString)
	if err != nil {
		cc.RedirectTo(c, "/")
		return
	}
	err = cc.CartService.DeleteItem(cookie.Value, cartItemID)
	if err != nil {
		cc.RedirectTo(c, "/")
		return
	}
	cc.RedirectTo(c, "/")

}
