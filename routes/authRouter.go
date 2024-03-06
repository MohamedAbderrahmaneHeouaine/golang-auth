package routes

import (
	contoller "github.com/MohamedAbderrahmaneHeouaine/golang-auth/controllers"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/users/signup", contoller.Signup())
	incomingRoutes.POST("/users/login", contoller.Login())
}
