package routes

import (
	"ecommerce/controller"
	"net/http"
)

var welcomeRoute = []Route{
	Route{
		URI:          "/",
		Method:       http.MethodGet,
		Handler:      controller.Welcome,
		AuthRequired: false,
	},
	Route{
		URI:          "/add",
		Method:       http.MethodGet,
		Handler:      controller.Add,
		AuthRequired: false,
	},
	Route{
		URI:          "/products",
		Method:       http.MethodGet,
		Handler:      controller.GetAll,
		AuthRequired: false,
	},
}
