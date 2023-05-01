package routers

import (
	"library/src/controllers"

	"github.com/gin-gonic/gin"
)

func SetAuthRouter(router *gin.Engine) {
	group := router.Group("/auth")
	{
		group.PUT("/signin", controllers.SignIn)
		group.POST("/signup", controllers.SignUp)
	}

}
