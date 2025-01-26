package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestHelloHandler checks if the helloHandler responds with the correct message and headers.
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

	// Verify security headers
	headers := map[string]string{
		"X-Content-Type-Options":        "nosniff",
		"Cross-Origin-Opener-Policy":    "same-origin",
		"Cross-Origin-Embedder-Policy":  "require-corp",
	}
	for key, value := range headers {
		if rr.Header().Get(key) != value {
			t.Errorf("header %s is missing or incorrect: got %v want %v", key, rr.Header().Get(key), value)
		}
	}
}

// TestHiddenFileMiddleware checks if the middleware blocks access to hidden files.
func TestHiddenFileMiddleware(t *testing.T) {
	hiddenPaths := []string{
		"/.git",
		"/.env",
		"/folder/.hidden",
	}

	for _, path := range hiddenPaths {
		req, err := http.NewRequest("GET", path, nil)
		if err != nil {
			t.Fatal(err)
		}

		// Record the response.
		rr := httptest.NewRecorder()

		// Create a dummy handler to test the middleware
		dummyHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		})

		// Wrap the dummy handler with the hidden file middleware
		handler := hiddenFileMiddleware(dummyHandler)
		handler.ServeHTTP(rr, req)

		// Check the status code.
		if status := rr.Code; status != http.StatusForbidden {
			t.Errorf("middleware did not block hidden file access for path %s: got %v want %v", path, status, http.StatusForbidden)
		}
	}
}
