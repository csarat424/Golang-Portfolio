package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestFileServer tests the file server
func TestFileServer(t *testing.T) {
	// Create a request to pass to our handler
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fs := http.FileServer(http.Dir("static"))
		fs.ServeHTTP(w, r)
	})

	// Serve the HTTP request
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect
	expected := "<!DOCTYPE html>"
	if !contains(rr.Body.String(), expected) {
		t.Errorf("handler returned unexpected body: got %v want contains %v",
			rr.Body.String(), expected)
	}
}

// contains checks if the substring is present in the string
func contains(s, substr string) bool {
	return len(s) >= len(substr) && s[:len(substr)] == substr
}
