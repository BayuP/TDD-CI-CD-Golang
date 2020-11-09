package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	URI          string
	Method       string
	Handler      func(w http.ResponseWriter, r *http.Request)
	AuthRequired bool
}

func Load() []Route {
	routes := welcomeRoute

	return routes
}

func SetupRouteWithMiddleware(r *mux.Router) *mux.Router {
	for _, route := range Load() {
		api := r.PathPrefix("/api/v1").Subrouter()
		api.HandleFunc(route.URI, route.Handler).Methods(route.Method)

	}

	return r
}
