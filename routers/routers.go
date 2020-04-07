package routers

import (
	"workshop-service/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	r.GET("/authors", controllers.ListAuthor)
	r.POST("/authors/create", controllers.AddNewAuthor)
	return r
}
