package views

import (
	"github.com/gin-gonic/gin"
	"github.com/sivayogasubramanian/ocv/src/handlers"
	"github.com/sivayogasubramanian/ocv/src/viewmodels"
	"net/http"
)

func Suspend(ctx *gin.Context) {
	req := viewmodels.SuspendRequest{}
	bindJson(ctx, &req)

	err := handlers.Suspend(&req)

	if err != nil {
		ctx.JSON(err.StatusCode(), NewErrorResponse(err.Message()))
		return
	}

	ctx.Status(http.StatusNoContent)
}
