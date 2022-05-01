package main

import (
	"math/rand"
	"net/http"
	"time"
	 "os"

	"github.com/labstack/echo/v4"
)

type Book struct {
	Title  string
	Author string
}

var books []Book = []Book{
	{"Seva Sadan", "Premchand"},
	{"Rashmirathi", "Ramdhari Singh Dinkar"},
	{"Harry Potter", "J K Rowling"},
	{"Malgudi Days", "R K narayn"},
	{"Roads to Mussoorie", "Ruskin Bond"},
}

func main() {
	rand.Seed(time.Now().Unix())

	api := echo.New()

	api.GET("/books", getBook)
	api.POST("/books", getBook)
	api.PUT("/books", getBook)

	port := os.Getenv("PORT")

	if port == "" {
		port = "25255"
	}


	api.Start(":"+port)

}

func getBook(c echo.Context) error {
	index := rand.Intn(len(books))
	return c.JSON(http.StatusOK, books[index])
}
