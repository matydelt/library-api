package services

import (
	"context"
	"fmt"
	"library/src/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func SignIn(credentials models.User) models.User {
	userCollection := models.GetUserCollection()
	user := models.User{}
	cursor := userCollection.FindOne(context.TODO(), bson.M{"username": credentials.Username})
	if cursor.Err() != nil {
		fmt.Println(cursor.Err())
	}
	err := cursor.Decode(&user)
	if err != nil {
		fmt.Println(err)
	}
	if user.Username == "" {
		fmt.Println("El usuario no existe")
	}
	if user.Password != credentials.Password {
		fmt.Println("La contraseña es incorrecta")
	}
	return user
}

func SignUp(user models.User) models.User {
	userCollection := models.GetUserCollection()
	cursor, err := userCollection.InsertOne(context.TODO(), user)
	if err != nil {
		fmt.Println(err)
	}
	user.Id = cursor.InsertedID.(primitive.ObjectID)
	return user
}

func ValidateRequest(c *gin.Context, user models.User) bool {
	if user.Password == "" {
		c.JSON(400, gin.H{"error": "You need a password"})
		return false
	}
	if user.Username == "" {
		c.JSON(400, gin.H{"error": "You need a username"})
		return false
	}
	if GetUserByUsername(user.Username).Username != "" {
		c.JSON(400, gin.H{"error": "User already exists"})
		return false
	}
	return true
}

func GetUserByUsername(username string) models.User {
	userCollection := models.GetUserCollection()
	user := models.User{}
	cursor := userCollection.FindOne(context.TODO(), bson.M{"username": username})
	if cursor.Err() != nil {
		fmt.Println(cursor.Err())
	}
	err := cursor.Decode(&user)
	if err != nil {
		fmt.Println(err)
	}
	return user
}

func GetCredentialsFromRequest(c *gin.Context) models.User {
	var user models.User
	err := c.BindJSON(&user)
	if err != nil {
		fmt.Println(err)
	}
	return user
}
