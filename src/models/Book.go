package models

import (
	"library/src/configs"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Book struct {
	Id     primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Title  string             `json:"title"`
	Author string             `json:"author"`
	Pdf    string             `json:"pdf"`
}

func GetBookCollection() mongo.Collection {
	var bookCollection mongo.Collection = *configs.GetDB().Collection("books")

	return bookCollection
}
