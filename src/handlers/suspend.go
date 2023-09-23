package handlers

import (
	"fmt"
	"github.com/sivayogasubramanian/ocv/src/dataaccess"
	ocverrs "github.com/sivayogasubramanian/ocv/src/errors"
	"github.com/sivayogasubramanian/ocv/src/models"
	"github.com/sivayogasubramanian/ocv/src/viewmodels"
	"gorm.io/gorm"
	"net/http"
)

func Suspend(db *gorm.DB, req *viewmodels.SuspendRequest) ocverrs.Error {
	student := models.Student{Email: req.Student}
	if err := student.Validate(); err != nil {
		return ocverrs.New(http.StatusBadRequest, err.Error())
	}

	studentExists, err := dataaccess.DoesStudentExists(db, req.Student)
	if err != nil {
		return ocverrs.New(http.StatusInternalServerError, "An error occurred while suspending the student.")
	}
	if !studentExists {
		return ocverrs.New(http.StatusNotFound, fmt.Sprintf("Student with email: %s does not exist.", req.Student))
	}

	err = dataaccess.SuspendStudent(db, &student)
	if err != nil {
		return ocverrs.New(
			http.StatusInternalServerError,
			fmt.Sprintf("Failed to suspend student with email: %s.", req.Student),
		)
	}

	return nil
}
