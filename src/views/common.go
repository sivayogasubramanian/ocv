package views

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Error struct {
	Message string `json:"message"`
}

func NewErrorResponse(message string) *Error {
	return &Error{Message: message}
}

func bindJson(ctx *gin.Context, req interface{}) {
	err := ctx.BindJSON(&req)
	if err == nil {
		return
	}
	ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request"})
}
