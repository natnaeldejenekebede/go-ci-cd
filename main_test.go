package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestHelloHandler checks if the helloHandler responds with the correct message.
func TestHelloHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(helloHandler)
	handler.ServeHTTP(rr, req)

	// Check the status code.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check the response body.
	expected := "Hello, World!\n"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}
