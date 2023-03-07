package views

import (
	"github.com/gin-gonic/gin"
	"github.com/sivayogasubramanian/ocv/handlers"
	"net/http"
)

func GetCommonStudents(ctx *gin.Context) {
	stderr := ctx.Request.ParseForm()
	if stderr != nil {
		ctx.JSON(http.StatusBadRequest, "Invalid request")
	}

	var teachers []string
	teachers = append(teachers, ctx.Request.Form["teacher"]...)

	resp, err := handlers.GetCommonStudents(teachers)
	if err != nil {
		ctx.JSON(err.StatusCode(), NewErrorResponse(err.Message()))
		return
	}
	
	ctx.JSON(http.StatusOK, resp)
}
