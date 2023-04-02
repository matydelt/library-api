package services

import (
	"context"
	"fmt"
	"library/src/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func createFilter(title, author string) bson.D {
	query := bson.D{}
	if title != "" {
		query = append(query, bson.E{Key: "title", Value: title})
	}
	if author != "" {
		query = append(query, bson.E{Key: "author", Value: author})
	}
	return query
}

func GetBooks(title, author string) []models.Book {
	bookCollection := models.GetBookCollection()
	books := []models.Book{}
	filter := createFilter(title, author)

	cursor, err := bookCollection.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println(err)
	}
	for cursor.Next(context.TODO()) {
		var book models.Book
		err := cursor.Decode(&book)
		if err != nil {
			fmt.Println(err)
		}
		books = append(books, book)
	}
	return books
}

func GetBook(id string) models.Book {
	bookCollection := models.GetBookCollection()
	book := models.Book{}
	_id, err1 := primitive.ObjectIDFromHex(id)
	if err1 != nil {
		// manejar el error
		fmt.Println(err1)
	}
	cursor := bookCollection.FindOne(context.TODO(), bson.M{"_id": _id})
	err := cursor.Decode(&book)
	if err != nil {
		fmt.Println(err)
	}

	return book
}

func CreateBook(book models.Book) {
	bookCollection := models.GetBookCollection()

	_, insertError := bookCollection.InsertOne(context.TODO(), book)
	if insertError != nil {
		fmt.Println(insertError)
	}
}

func GetBookFromRequest(c *gin.Context) models.Book {
	var book models.Book
	err := c.BindJSON(&book)
	if err != nil {
		fmt.Println(err)
	}
	return book
}
