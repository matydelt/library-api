package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"library/src/models"
	"library/src/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetBooks(c *gin.Context) {

	title := c.Query("title")
	author := c.Query("author")
	books := services.GetBooks(title, author)

	c.IndentedJSON(http.StatusOK, books)
}

func CreateBook(c *gin.Context) {

	req := c.Request
	body, bodyError := ioutil.ReadAll(req.Body)
	if bodyError != nil {
		fmt.Println(bodyError)
	}
	book := models.Book{}
	readError := json.Unmarshal(body, &book)
	if readError != nil {
		fmt.Println(readError)
	}
	fmt.Println(book)
	if book.Author == "" || book.Pdf == "" || book.Title == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
		return
	}
	services.CreateBook(book)
}
