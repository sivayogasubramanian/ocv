package routes

import (
	"github.com/gin-gonic/gin"
	views2 "github.com/sivayogasubramanian/ocv/src/views"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")
	api.POST("/register", views2.Register)
	api.GET("/commonstudents", views2.GetCommonStudents)
	api.POST("/suspend", views2.Suspend)
	api.POST("/retrievefornotifications", views2.RetrieveNotifications)

	return r
}
