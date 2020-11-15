package payload

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ResponseToken struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Token   string      `json:"token"`
	Data    interface{} `json:"data"`
}

type FieldResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func ErrorResponse(w http.ResponseWriter, status int, err error) {
	w.WriteHeader(status)

	res := Response{
		Message: err.Error(),
	}

	err = json.NewEncoder(w).Encode(res)

	if err != nil {
		log.Fatal(err)
	}
}

func MessageResponse(w http.ResponseWriter, message string, data interface{}, status int) {
	res := Response{
		Status:  status,
		Message: message,
		Data:    data,
	}

	fmt.Println(res)
	//err := json.NewEncoder(w).Encode(res)
	body, err := json.Marshal(res)
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(body)

}

func ErrorField(w http.ResponseWriter, message string, status int) {
	w.WriteHeader(status)
	res := FieldResponse{
		Message: message,
		Status:  status,
	}

	err := json.NewEncoder(w).Encode(res)

	if err != nil {
		log.Fatal(err)
	}
}

func MessageToken(w http.ResponseWriter, message string, token string, data interface{}, status int) {
	w.WriteHeader(status)
	res := ResponseToken{
		Status:  status,
		Message: message,
		Token:   token,
		Data:    data,
	}

	err := json.NewEncoder(w).Encode(res)

	if err != nil {
		log.Fatal(err)
	}
}
