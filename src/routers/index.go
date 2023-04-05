package routers

import "github.com/gin-gonic/gin"

func InitServer() {
	router := gin.Default()
	SetBookRouter(router)
	SetUserRouter(router)

	router.Run("localhost:3001")
}
