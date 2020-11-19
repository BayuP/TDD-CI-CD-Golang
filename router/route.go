package router

import (
	"ecommerce/common"
	"ecommerce/router/routes"
	"log"

	"github.com/gorilla/mux"
)

func New() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)

	err := common.Connect()
	if err != nil {
		log.Fatal(err)
	}
	return routes.SetupRouteWithMiddleware(r)
}
