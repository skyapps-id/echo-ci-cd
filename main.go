package main

import (
	"context"
	"echo-ci-cd/driver"
	"echo-ci-cd/handler"
	"echo-ci-cd/repository"
	"echo-ci-cd/usecase"
	"echo-ci-cd/utils"
	"fmt"
	"os"

	"github.com/labstack/echo/v4"
)

// Dependency Injection
var ctx = context.Background()
var db = driver.NewGormDatabase()

var bookRepository = repository.NewBookRepository(db)
var bookUsecase = usecase.NewUsecase(bookRepository)
var bookHandler = handler.NewHandler(bookUsecase)

func helloHandler(ctx echo.Context) error {
	return utils.Response(ctx, "Hello, eFishery !!!", nil)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Create an Echo instance
	e := echo.New()

	// Define a route for handling GET requests to /hello
	e.GET("/hello", helloHandler)
	e.GET("/books/:uuid", bookHandler.GetBookByID)
	e.GET("/books", bookHandler.GetBooks)

	// Start the server on port 8080
	fmt.Println("Mode : " + os.Getenv("ENV"))
	e.Start(":" + port)

}
