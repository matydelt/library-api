package models

import (
	"library/src/configs"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	Id       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Username string             `json:"username"`
	Password string             `json:"password"`
}

func GetUserCollection() mongo.Collection {
	var userCollection mongo.Collection = *configs.GetDB().Collection("users")
	return userCollection
}
