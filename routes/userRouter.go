package routes

import (
	contoller "github.com/MohamedAbderrahmaneHeouaine/golang-auth/controllers"
	"github.com/MohamedAbderrahmaneHeouaine/golang-auth/middleware"
	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/users", contoller.GetUsers())
	incomingRoutes.GET("/users/:id", contoller.GetUser())
}
