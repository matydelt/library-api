package routers

import "github.com/gin-gonic/gin"

func InitServer() {
	router := gin.Default()
	SetBookRouter(router)
	SetAuthRouter(router)

	router.Run("localhost:3001")
}
