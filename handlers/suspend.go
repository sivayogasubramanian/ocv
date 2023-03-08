package handlers

import (
	"fmt"
	"github.com/sivayogasubramanian/ocv/dataaccess"
	ocverrs "github.com/sivayogasubramanian/ocv/errors"
	"github.com/sivayogasubramanian/ocv/models"
	"github.com/sivayogasubramanian/ocv/viewmodels"
	"net/http"
)

func Suspend(req *viewmodels.SuspendRequest) ocverrs.Error {
	student := models.Student{Email: req.Student}
	if err := student.Validate(); err != nil {
		return ocverrs.New(http.StatusBadRequest, err.Error())
	}

	studentExists, err := dataaccess.DoesStudentExists(req.Student)
	if err != nil {
		return ocverrs.New(http.StatusInternalServerError, "An error occurred while suspending the student.")
	}
	if !studentExists {
		return ocverrs.New(http.StatusNotFound, fmt.Sprintf("Student with email: %s does not exist.", req.Student))
	}

	err = dataaccess.SuspendStudent(&student)
	if err != nil {
		return ocverrs.New(
			http.StatusInternalServerError,
			fmt.Sprintf("Failed to suspend student with email: %s.", req.Student),
		)
	}

	return nil
}
