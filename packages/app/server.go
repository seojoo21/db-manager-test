package main

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/seojoo21/dbconnector-test/packages/app/repository"
	"github.com/seojoo21/dbconnector-test/packages/app/service"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/books", getBookList)
	e.GET("/books/:id", getBook)
	e.POST("/books", saveBook)
	e.PUT("/books/:id", updateBook)
	e.DELETE("/books/:id", deleteBook)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

// Handler
func getBookList(c echo.Context) error {
	book := service.NewBookService(repository.InitBookRepository())
	res, err := book.GetBookList()

	if err != nil {
		return c.String(http.StatusInternalServerError, "internal server error")
	}

	jsonRes, _ := json.Marshal(res)

	return c.String(http.StatusOK, string(jsonRes))
}

func getBook(c echo.Context) error {
	id := c.Param("id")

	book := service.NewBookService(repository.InitBookRepository())
	res, err := book.GetBook(id)

	if err != nil {
		return c.String(http.StatusInternalServerError, "internal server error")
	}

	jsonRes, _ := json.Marshal(res)

	return c.String(http.StatusOK, string(jsonRes))
}

type bookRequestParams struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

func saveBook(c echo.Context) error {
	requestParams := &bookRequestParams{}

	if err := c.Bind(requestParams); err != nil {
		return c.String(http.StatusInternalServerError, "internal server error")
	}

	book := service.NewBookService(repository.InitBookRepository())
	_, err := book.SaveBook(service.BookRequest{Title: requestParams.Title, Author: requestParams.Author})

	if err != nil {
		return c.String(http.StatusInternalServerError, "internal server error")
	}

	return c.String(http.StatusCreated, "created")
}

func updateBook(c echo.Context) error {
	id := c.Param("id")

	requestParams := &bookRequestParams{}

	if err := c.Bind(requestParams); err != nil {
		return c.String(http.StatusInternalServerError, "internal server error")
	}

	book := service.NewBookService(repository.InitBookRepository())
	_, err := book.UpdateBook(service.BookRequest{Id: id, Title: requestParams.Title, Author: requestParams.Author})

	if err != nil {
		return c.String(http.StatusInternalServerError, "internal server error")
	}

	return c.String(http.StatusNoContent, "updated")
}

func deleteBook(c echo.Context) error {
	id := c.Param("id")

	book := service.NewBookService(repository.InitBookRepository())
	_, err := book.DeleteBook(id)

	if err != nil {
		return c.String(http.StatusInternalServerError, "internal server error")
	}

	return c.String(http.StatusNoContent, "deleted")
}
