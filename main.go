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
	port := os.Getenv("PORT")
	r := router.New()

	log.Fatal(http.ListenAndServe(":"+port, r))
}

//TestFunc ...
func TestFunc(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Testing api")
}
