package router

import (
	"ecommerce/router/routes"

	"github.com/gorilla/mux"
)

func New() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)

	return routes.SetupRouteWithMiddleware(r)
}
