package views

import (
	"github.com/gin-gonic/gin"
	"github.com/sivayogasubramanian/ocv/handlers"
	"github.com/sivayogasubramanian/ocv/viewmodels"
	"net/http"
)

func Register(ctx *gin.Context) {
	req := viewmodels.RegisterRequest{}
	bindJson(ctx, &req)

	err := handlers.Register(&req)

	if err != nil {
		ctx.JSON(err.StatusCode(), NewErrorResponse(err.Message()))
		return
	}

	ctx.Status(http.StatusNoContent)
}
