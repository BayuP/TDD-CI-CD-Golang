package controller

import (
	"ecommerce/payload"
	"ecommerce/repository"
	"net/http"
)

func Welcome(w http.ResponseWriter, r *http.Request) {

	payload.MessageResponse(w, "Testing", nil, http.StatusForbidden)

}

func Add(w http.ResponseWriter, r *http.Request) {
	payload.MessageResponse(w, "Sucessfull", nil, http.StatusOK)
}

func GetAll(w http.ResponseWriter, r *http.Request) {
	allData, err := repository.GetAllProduct()
	if err != nil {
		payload.MessageResponse(w, "Something Wrong", err, http.StatusInternalServerError)
		return
	}
	payload.MessageResponse(w, "Sucessfull", allData, http.StatusOK)
}
