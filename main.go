package main

import (
	"fmt"
	"log"
	"os"

	//"log"
	"net/http"

	"ecommerce/router"
)

func main() {
	r := router.New()

	log.Fatal(http.ListenAndServe(os.Getenv("PORT"), r))
}

//TestFunc ...
func TestFunc(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Testing api")
}
