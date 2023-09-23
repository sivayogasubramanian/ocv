package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sivayogasubramanian/ocv/views"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")
	api.POST("/register", views.Register)
	api.GET("/commonstudents", views.GetCommonStudents)
	api.POST("/suspend", views.Suspend)
	api.POST("/retrievefornotifications", views.RetrieveNotifications)

	return r
}
