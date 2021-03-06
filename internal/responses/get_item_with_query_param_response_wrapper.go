package responses

import "github.com/aditya109/go-server-template/internal/models"

// GetItemWithQueryParamResponseWrapper is an item returned in response
// swagger:response GetItemWithQueryParamResponse
type GetItemWithQueryParamResponseWrapper struct {
	// Item in the system
	// in: body
	Body []models.Item
}
