package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/sivayogasubramanian/ocv/models"
	"net/http"
)

type RegisterRequest struct {
	Teacher  string   `json:"teacher"`
	Students []string `json:"students"`
}

func Register(ctx *gin.Context) {
	req := RegisterRequest{}
	err := ctx.BindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	}

	teacher := models.Teacher{Email: req.Teacher}
	for _, studentEmail := range req.Students {
		student := models.Student{Email: studentEmail}
		teacher.Students = append(teacher.Students, &student)
	}

	err = models.CreateTeacher(&teacher)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	} else {
		ctx.Status(http.StatusNoContent)
	}
}
