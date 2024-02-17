package service

import (
	"fmt"
	"html/template"
	"strings"
)

// GetCartData retrieves cart data based on the sessionID.
func (cs *CartService) GetCartData(sessionID string, qErr string) (string, error) {
	data := map[string]interface{}{
		"Error":     qErr,
		"CartItems": cs.GetCartItemData(sessionID),
	}

	return renderTemplate(data)
}

// getCartItemData retrieves cart item data based on the sessionID.
func (cs *CartService) GetCartItemData(sessionID string) []map[string]interface{} {
	items, err := cs.repo.GetCartData(sessionID)
	if err != nil {
		return nil
	}

	return items
}

var pathToTemplates = "../../static/add_item_form.html"

func renderTemplate(pageData interface{}) (string, error) {
	// Read and parse the HTML template file
	tmpl, err := template.ParseFiles(pathToTemplates)
	if err != nil {
		return "", fmt.Errorf("Error parsing template: %v ", err)
	}

	// Create a strings.Builder to store the rendered template
	var renderedTemplate strings.Builder

	err = tmpl.Execute(&renderedTemplate, pageData)
	if err != nil {
		return "", fmt.Errorf("Error parsing template: %v ", err)
	}

	// Convert the rendered template to a string
	resultString := renderedTemplate.String()

	return resultString, nil
}
