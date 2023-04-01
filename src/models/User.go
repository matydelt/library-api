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

var db *mongo.Database = configs.GetDB()
var userCollection mongo.Collection = *db.Collection("users")

func GetUserCollection() mongo.Collection {
	return userCollection
}
