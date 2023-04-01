package services

import (
	"context"
	"fmt"
	"library/src/models"

	"go.mongodb.org/mongo-driver/bson"
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
	query := createFilter(title, author)

	cursor, err := bookCollection.Find(context.TODO(), query)
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

func CreateBook(book models.Book) {
	bookCollection := models.GetBookCollection()

	_, insertError := bookCollection.InsertOne(context.TODO(), book)
	if insertError != nil {
		fmt.Println(insertError)
	}
}
