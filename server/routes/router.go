package routes

import (
	"api-temp_conv/controllers"

	"github.com/gin-gonic/gin"
)



func ConfigRoutes(router *gin.Engine) *gin.Engine {

	// Criação das rotas
	main := router.Group("")
	{
		main.GET("/:scale/:val", controllers.Converte)
	}	
	return router

}