package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sivayogasubramanian/ocv/src/views"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	api := r.Group("/api")

	api.POST("/register", func(ctx *gin.Context) {
		views.Register(db, ctx)
	})

	api.GET("/commonstudents", func(ctx *gin.Context) {
		views.GetCommonStudents(db, ctx)
	})

	api.POST("/suspend", func(ctx *gin.Context) {
		views.Suspend(db, ctx)
	})

	api.POST("/retrievefornotifications", func(ctx *gin.Context) {
		views.RetrieveNotifications(db, ctx)
	})

	return r
}
