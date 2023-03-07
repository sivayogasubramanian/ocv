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
	for _, teacher := range ctx.Request.Form["teacher"] {
		teachers = append(teachers, teacher)
	}

	resp, err := handlers.GetCommonStudents(teachers)
	if err != nil {
		ctx.JSON(err.StatusCode(), NewErrorResponse(err.Message()))
	} else {
		ctx.JSON(http.StatusOK, resp)
	}
}
