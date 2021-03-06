package router

import (
	"net/http"

	h "github.com/aditya109/go-server-template/internal/handlers"
	"github.com/go-openapi/runtime/middleware"
)

// Route is container template for server routes
type Route struct {
	Name            string
	Method          string
	Pattern         string
	HandlerFunction http.HandlerFunc
	Handler         http.Handler
}

type routes []Route

var routeList = routes{

	// swagger:route GET / home welcome
	// Returns a welcome message
	// responses:
	// 	200: WelcomeResponse
	Route{
		"Welcome",
		"GET",
		"/",
		h.WelcomeHandler,
		nil,
	},
	// swagger:route GET /items items listItems
	// Returns a list of items, no query params required
	// responses:
	// 	200: GetItemsResponse
	Route{
		"GetItems",
		"GET",
		"/items",
		h.GetItemsHandler,
		nil,
	},
	// swagger:route GET /item/{id} item listItemById
	// Returns an item with id from the existing list of items
	// responses:
	// 	200: GetItemWithIdReponse
	Route{
		"GetItemWithId",
		"GET",
		"/item/{id}",
		h.GetItemWithIDHandler,
		nil,
	},

	// swagger:route GET /item item listFilteredItems
	// Returns items filtered by query parameters from the existing list of items
	// responses:
	// 	200: GetWithQueryParamsReponse
	Route{
		"GetWithQueryParams",
		"GET",
		"/item",
		h.GetWithQueryParamsHandler,
		nil,
	},

	// swagger:route GET /docs docs swaggerDocumentation
	// Returns swagger specification uunder OpenAPIv3 documeted APIs
	Route{
		"swaggerDocumentation",
		"GET",
		"/docs",
		nil,
		middleware.Redoc(middleware.RedocOpts{
			SpecURL: "/swagger.yaml",
		}, nil),
	},

	Route{
		"Swagger JSON",
		"GET",
		"/swagger.yaml",
		nil,
		http.FileServer(http.Dir("./api/swagger")),
	},
}
