package routers

import (
	"library/src/controllers"

	"github.com/gin-gonic/gin"
)

func GetBookRouter(router *gin.Engine) {
	router.GET("/books", controllers.GetBooks)
	router.GET("/book/:id", controllers.GetBook)
	router.POST("/book", controllers.CreateBook)
}
