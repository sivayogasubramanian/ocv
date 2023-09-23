package views

import (
	"github.com/gin-gonic/gin"
	"github.com/sivayogasubramanian/ocv/src/handlers"
	"github.com/sivayogasubramanian/ocv/src/viewmodels"
	"gorm.io/gorm"
	"net/http"
)

func Register(db *gorm.DB, ctx *gin.Context) {
	req := viewmodels.RegisterRequest{}
	bindJson(ctx, &req)

	err := handlers.Register(db, &req)

	if err != nil {
		ctx.JSON(err.StatusCode(), NewErrorResponse(err.Message()))
		return
	}

	ctx.Status(http.StatusNoContent)
}
