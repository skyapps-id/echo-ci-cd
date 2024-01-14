package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func helloHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, Efishery !!!")
}

func main() {
	// Create an Echo instance
	e := echo.New()

	// Define a route for handling GET requests to /hello
	e.GET("/hello", helloHandler)

	// Start the server on port 8080
	e.Start(":8080")
}
