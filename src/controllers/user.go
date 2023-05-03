package controllers

import (
	"library/src/services"
	"library/utils"

	"github.com/gin-gonic/gin"
)

func SignIn(c *gin.Context) {
	user := services.GetCredentialsFromRequest(c)
	user = services.SignIn(user)
	if user.Id.Hex() == utils.NullId {
		c.JSON(400, gin.H{"error": "El usuario no existe"})
	} else {
		c.JSON(200, user)
	}
}

func SignUp(c *gin.Context) {
	credentials := services.GetCredentialsFromRequest(c)
	validated := services.ValidateRequest(c, credentials)
	if !validated {
		return
	}
	user := services.SignUp(credentials)
	if user.Id.Hex() == utils.NullId {
		c.JSON(400, gin.H{"error": "Error al crear el usuario"})
	} else {
		c.JSON(200, user)
	}
}
