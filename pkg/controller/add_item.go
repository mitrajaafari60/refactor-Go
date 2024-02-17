package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"interview/pkg/service"
	"net/http"
	"strconv"
)

func (cc *CartController) AddItem(c *gin.Context) {
	cookie, err := c.Request.Cookie(ICE_SESSION_ID)

	if err != nil || errors.Is(err, http.ErrNoCookie) || (cookie != nil && cookie.Value == "") {
		cc.RedirectTo(c, "/")
		return
	}

	if c.Request.Body == nil {
		cc.RedirectTo(c, "/?error="+"body cannot be nil")
		return
	}

	form := &service.CartItemForm{}
	if err := binding.FormPost.Bind(c.Request, form); err != nil {
		cc.RedirectTo(c, "/?error="+err.Error())
		return
	}

	quantity, err := strconv.ParseInt(form.Quantity, 10, 0)
	if err != nil {
		cc.RedirectTo(c, "/?error=invalid quantity")
		return
	}

	err = service.AddItemToCart(cookie.Value, form.Product, quantity)
	if err != nil {
		cc.RedirectTo(c, "/?error="+err.Error())
		return
	}
	cc.RedirectTo(c, "/")
}
