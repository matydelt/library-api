package routers

import "github.com/gin-gonic/gin"

func InitServer() {
	router := gin.Default()
	GetBookRouter(router)

	router.Run("localhost:3001")
}
