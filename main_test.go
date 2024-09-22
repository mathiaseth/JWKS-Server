package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAuthHandler(t *testing.T) {
	req, err := http.NewRequest("POST", "/auth?expire=false", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a response recorder to capture the response
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(jwksHandler) // Adjust to your handler name

	// Call the handler with the request and response recorder
	handler.ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check for a token in the response body
	// You can unmarshal and check specific fields if necessary
	// For now, I'm just checking if the response is not empty
	if rr.Body.String() == "" {
		t.Error("handler returned empty response body")
	}
}

func TestJWKSHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/.well-known/jwks.json", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(jwksHandler) // Adjust to your handler name

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check that the response body contains keys
	if rr.Body.String() == "{\"keys\":[]}" {
		t.Error("handler returned empty keys array")
	}
}
