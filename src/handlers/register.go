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

func Register(db *gorm.DB, req *viewmodels.RegisterRequest) ocverrs.Error {
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

	teacherExists, err := dataaccess.DoesTeacherExists(db, req.Teacher)
	if err != nil {
		return ocverrs.New(http.StatusInternalServerError, fmt.Sprintf("An error occurred while registering the teacher with email: %s.", req.Teacher))
	}
	if teacherExists {
		return ocverrs.New(http.StatusConflict, fmt.Sprintf("Teacher with email: %s already exists.", req.Teacher))
	}

	err = dataaccess.CreateTeacher(db, &newTeacher)
	if err != nil {
		return ocverrs.New(http.StatusInternalServerError, fmt.Sprintf("An error occurred while registering the teacher with email: %s.", req.Teacher))
	}

	return nil
}
