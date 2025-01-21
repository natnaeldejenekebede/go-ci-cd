package main

import (
    "net/http"
    "net/http/httptest"
    "testing"
)

func TestHelloHandler(t *testing.T) {
    req, err := http.NewRequest("GET", "/", nil)
    if err != nil {
        t.Fatal(err)
    }

    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(helloHandler)
    handler.ServeHTTP(rr, req)

    if rr.Code != http.StatusOK {
        t.Errorf("expected status code 200, got %d", rr.Code)
    }

    expected := "Hello, world!"
    if rr.Body.String() != expected {
        t.Errorf("expected body %s, got %s", expected, rr.Body.String())
    }
}
