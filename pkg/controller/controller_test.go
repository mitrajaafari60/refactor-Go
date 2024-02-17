package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"interview/pkg/service"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestRedirectTo(t *testing.T) {
	// Setup
	gin.SetMode(gin.TestMode)
	mockCartService := service.NewMockCartService()
	// create the controller
	cartController := CartController{CartService: mockCartService}

	// Set up Gin context
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest("GET", "/", nil)

	// Invoke the handler
	cartController.RedirectTo(c, "/")

	// Assertions
	assert.Equal(t, http.StatusFound, w.Code)
	assert.Equal(t, "/", w.Header().Get("Location"))
}

func TestSCartController_ShowAddItemForm(t *testing.T) {
	// mock dependencies
	cartService := service.NewMockCartService()
	cartService.GetCartData("temp_session", "")
	// mock the request
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	// create the controller
	controller := CartController{CartService: cartService}

	// call the method under test
	controller.ShowAddItemForm(c)

	// verify the response
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "text/html", w.Header().Get("Content-Type"))
	assert.NotNil(t, "cart data", w.Body.String())
}

func TestSCartController_AddItem(t *testing.T) {
	// mock dependencies
	cartService := service.NewMockCartService()
	// Mock Request
	formData := strings.NewReader("product=testProduct&quantity=2")
	req, err := http.NewRequest("POST", "/add-item", formData)
	assert.NoError(t, err)

	// Set cookie in the request
	req.AddCookie(&http.Cookie{
		Name:  "ice_session_id",
		Value: "your-session-id-value",
	})

	// Set up Gin context
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	// create the controller
	controller := CartController{CartService: cartService}

	// call the method under test
	controller.AddItem(c)

	// verify the response
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestAddItemRequestBodyNil(t *testing.T) {
	// Setup
	gin.SetMode(gin.TestMode)
	mockCartService := service.NewMockCartService()
	cartController := CartController{CartService: mockCartService}

	// Mock Request with nil body
	req, err := http.NewRequest("POST", "/add-item", nil)
	assert.NoError(t, err)

	// Set cookie in the request
	req.AddCookie(&http.Cookie{
		Name:  "ice_session_id",
		Value: "your-session-id-value",
	})

	// Set up Gin context
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	// Invoke the handler
	cartController.AddItem(c)

	// Assertions
	assert.Equal(t, "", w.Body.String())

}
func TestDeleteCartItem(t *testing.T) {
	// Setup
	gin.SetMode(gin.TestMode)
	mockCartService := service.NewMockCartService()
	cartController := CartController{CartService: mockCartService}

	// Set up Gin context with a valid session cookie
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest("GET", "/?cart_item_id=123", nil)
	c.Request.AddCookie(&http.Cookie{Name: ICE_SESSION_ID, Value: "valid_session_id"})

	// Invoke the handler
	cartController.DeleteCartItem(c)

	// Assertions
	assert.Equal(t, http.StatusFound, w.Code)
	// Add more assertions based on your specific requirements
}

func TestSetNewSessionCookie(t *testing.T) {
	// Setup
	gin.SetMode(gin.TestMode)
	mockCartService := service.NewMockCartService()
	// create the controller
	cartController := CartController{CartService: mockCartService}

	// Set up Gin context
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	// Invoke the handler
	cartController.SetNewSessionCookie(c)

	// Assertions
	cookies := w.Result().Cookies()
	assert.Len(t, cookies, 1)

	cookie := cookies[0]
	assert.Equal(t, ICE_SESSION_ID, cookie.Name)
	assert.NotEmpty(t, cookie.Value)
	assert.Equal(t, "/", cookie.Path)
	assert.Equal(t, "localhost", cookie.Domain)
	assert.False(t, cookie.Secure)
	assert.True(t, cookie.HttpOnly)
	assert.Equal(t, int(3600), cookie.MaxAge)
}
