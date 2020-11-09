package test

import (
	"ecommerce/controller"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestWelcomeApi(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/v1/", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(controller.Welcome)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusForbidden {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	//fmt.Println(rr.Code)
	// Check the response body is what we expect.
	expected := `{"status":403,"message":"Testing","data":null}`
	//strings.TrimRight(rr.Body.String(), "\n")
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)

	}
}

func TestAddApi(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/v1/add", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(controller.Add)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	//fmt.Println(rr.Code)
	// Check the response body is what we expect.
	expected := `{"status":200,"message":"Sucessfull","data":null}`
	//st rings.TrimRight(rr.Body.String(), "\n")
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)

	}
}
