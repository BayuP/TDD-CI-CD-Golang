package controller

import (
	"ecommerce/payload"
	"net/http"
)

func Welcome(w http.ResponseWriter, r *http.Request) {
	payload.MessageResponse(w, "Testing", nil, http.StatusForbidden)
}

func Add(w http.ResponseWriter, r *http.Request) {
	payload.MessageResponse(w, "Sucessfull", nil, http.StatusOK)
}
