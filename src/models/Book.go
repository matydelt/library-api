package models

import (
	"library/src/configs"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Book struct {
	Id          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Title       string             `json:"title" bson:"title"`
	Author      string             `json:"author" bson:"author"`
	Year        int                `json:"year" bson:"year"`
	Description string             `json:"description" bson:"description"`
	Pdf         string             `json:"pdf" bson:"pdf"`
}

func GetBookCollection() mongo.Collection {
	var bookCollection mongo.Collection = *configs.GetDB().Collection("books")

	return bookCollection
}
