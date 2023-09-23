package views

import (
	"github.com/gin-gonic/gin"
	"github.com/sivayogasubramanian/ocv/src/handlers"
	"github.com/sivayogasubramanian/ocv/src/viewmodels"
	"gorm.io/gorm"
	"net/http"
)

func RetrieveNotifications(db *gorm.DB, ctx *gin.Context) {
	req := viewmodels.RetrieveNotificationsRequest{}
	bindJson(ctx, &req)

	resp, err := handlers.RetrieveNotifications(db, &req)

	if err != nil {
		ctx.JSON(err.StatusCode(), NewErrorResponse(err.Message()))
		return
	}

	ctx.JSON(http.StatusOK, resp)
}
