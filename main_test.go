package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
)

func TestHelloHandler(t *testing.T) {
	// Create a new Echo instance
	e := echo.New()

	// Set up a request
	req := httptest.NewRequest(http.MethodGet, "/hello", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Call the handler
	if err := helloHandler(c); err != nil {
		t.Fatalf("helloHandler failed: %s", err)
	}

	// Check the response status code
	if rec.Code != http.StatusOK {
		t.Errorf("expected status %d; got %d", http.StatusOK, rec.Code)
	}

	// Check the response body
	expectedResponse := "Hello, World!"
	if rec.Body.String() != expectedResponse {
		t.Errorf("expected body %q; got %q", expectedResponse, rec.Body.String())
	}
}
