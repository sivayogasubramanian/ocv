package handlers

import (
	"fmt"
	"github.com/sivayogasubramanian/ocv/dataaccess"
	ocverrs "github.com/sivayogasubramanian/ocv/errors"
	"github.com/sivayogasubramanian/ocv/models"
	"github.com/sivayogasubramanian/ocv/viewmodels"
	"net/http"
)

func Register(req *viewmodels.RegisterRequest) ocverrs.Error {
	newTeacher := models.Teacher{Email: req.Teacher}
	if err := newTeacher.Validate(); err != nil {
		return ocverrs.New(http.StatusBadRequest, err.Error())
	}

	for _, studentEmail := range req.Students {
		newStudent := models.Student{Email: studentEmail}
		if err := newStudent.Validate(); err != nil {
			return ocverrs.New(http.StatusBadRequest, err.Error())
		}
		newTeacher.Students = append(newTeacher.Students, &newStudent)
	}

	teacherExists, err := dataaccess.DoesTeacherExists(req.Teacher)
	if err != nil {
		return ocverrs.New(http.StatusInternalServerError, fmt.Sprintf("An error occurred while registering the teacher with email: %s.", req.Teacher))
	}
	if teacherExists {
		return ocverrs.New(http.StatusConflict, fmt.Sprintf("Teacher with email: %s already exists.", req.Teacher))
	}

	err = dataaccess.CreateTeacher(&newTeacher)
	if err != nil {
		return ocverrs.New(http.StatusInternalServerError, fmt.Sprintf("An error occurred while registering the teacher with email: %s.", req.Teacher))
	}

	return nil
}
