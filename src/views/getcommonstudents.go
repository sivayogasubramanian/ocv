package views

import (
	"github.com/gin-gonic/gin"
	"github.com/sivayogasubramanian/ocv/src/handlers"
	"gorm.io/gorm"
	"net/http"
)

func GetCommonStudents(db *gorm.DB, ctx *gin.Context) {
	stderr := ctx.Request.ParseForm()
	if stderr != nil {
		ctx.JSON(http.StatusBadRequest, "Invalid request")
	}

	var teachers []string
	teachers = append(teachers, ctx.Request.Form["teacher"]...)

	resp, err := handlers.GetCommonStudents(db, teachers)
	if err != nil {
		ctx.JSON(err.StatusCode(), NewErrorResponse(err.Message()))
		return
	}

	ctx.JSON(http.StatusOK, resp)
}
