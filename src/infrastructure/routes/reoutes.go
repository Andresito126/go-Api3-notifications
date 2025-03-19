package routes

import (
	"github.com/Andresito126/api3-notifications/src/infrastructure/controllers"

	"github.com/gin-gonic/gin"
)


func RegisterRoutes(router *gin.Engine) {
	
	routes := router.Group("/notification")
	{
		
		routes.POST("", controllers.NewSendEmailController().Run)
	}
}
