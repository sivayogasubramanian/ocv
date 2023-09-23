package views

import (
	"github.com/gin-gonic/gin"
	"github.com/sivayogasubramanian/ocv/src/handlers"
	"github.com/sivayogasubramanian/ocv/src/viewmodels"
	"net/http"
)

func RetrieveNotifications(ctx *gin.Context) {
	req := viewmodels.RetrieveNotificationsRequest{}
	bindJson(ctx, &req)

	resp, err := handlers.RetrieveNotifications(&req)

	if err != nil {
		ctx.JSON(err.StatusCode(), NewErrorResponse(err.Message()))
		return
	}

	ctx.JSON(http.StatusOK, resp)
}
