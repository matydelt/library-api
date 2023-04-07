package routers

import (
	"library/src/controllers"

	"github.com/gin-gonic/gin"
)

func SetAuthRouter(router *gin.Engine) {
	group := router.Group("/auth")
	{
		group.PUT("/signIn", controllers.SignIn)
		group.POST("/signUp", controllers.SignUp)
	}
}
