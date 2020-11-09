package main

import (
	"fmt"
	"log"

	//"log"
	"net/http"

	"ecommerce/router"
)

func main() {
	r := router.New()

	log.Fatal(http.ListenAndServe(":8000", r))
}

//TestFunc ...
func TestFunc(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Testing api")
}
