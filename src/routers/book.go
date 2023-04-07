package routers

import (
	"library/src/controllers"

	"github.com/gin-gonic/gin"
)

func SetBookRouter(router *gin.Engine) {
	group := router.Group("/book")
	{
		group.GET("/all", controllers.GetBooks)
		group.GET("/:id", controllers.GetBook)
		group.POST("", controllers.CreateBook)
	}
}
