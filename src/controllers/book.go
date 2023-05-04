package controllers

import (
	"library/src/services"
	"library/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetBooks(c *gin.Context) {
	// panic("not implemented")
	title := c.Query("title")
	author := c.Query("author")
	books := services.GetBooks(title, author)

	c.IndentedJSON(http.StatusOK, books)
}

func GetBook(c *gin.Context) {

	id := c.Param("id")

	book := services.GetBook(id)
	if book.Id.Hex() == utils.NullId {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, book)
}

func CreateBook(c *gin.Context) {
	book := services.GetBookFromRequest(c)

	if book.Author == "" || book.Pdf == "" || book.Title == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
		return
	}
	services.CreateBook(book)
}
