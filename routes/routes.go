package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sivayogasubramanian/ocv/handlers"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")
	api.POST("/register", handlers.Register)

	return r
}
