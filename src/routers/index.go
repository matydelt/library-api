package routers

import (
	"library/src/middlewares"

	"github.com/gin-gonic/gin"
)

func InitServer() {
	router := gin.Default()
	router.Use(middlewares.Logger())
	SetBookRouter(router)
	SetAuthRouter(router)

	router.Run("localhost:3001")
}
