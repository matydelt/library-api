package routers

import (
	"library/src/controllers"

	"github.com/gin-gonic/gin"
)

func SetUserRouter(router *gin.Engine) {
	router.PUT("/signIn", controllers.SignIn)
	router.POST("/signUp", controllers.SignUp)
}
