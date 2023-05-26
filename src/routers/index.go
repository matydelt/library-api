package routers

import (
	"library/src/middlewares"

	"github.com/gin-gonic/gin"
)

func InitServer() {
	router := gin.Default()
	SetBookRouter(router)
	router.Use(middlewares.Logger())
	SetAuthRouter(router)

	router.Run("localhost:3001")
}
