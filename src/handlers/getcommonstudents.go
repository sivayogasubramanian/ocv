package handlers

import (
	"github.com/sivayogasubramanian/ocv/src/dataaccess"
	"github.com/sivayogasubramanian/ocv/src/models"
	"github.com/sivayogasubramanian/ocv/src/viewmodels"
	"net/http"
)

func GetCommonStudents(teacherEmails []string) (*viewmodels.CommonStudentsResponse, ocverrs.Error) {
	teachers, stderr := dataaccess.FindAllTeachers(teacherEmails)
	if stderr != nil {
		return nil, ocverrs.New(http.StatusInternalServerError, "An error occurred while getting the common students.")
	}
	if len(teachers) == 0 {
		return nil, ocverrs.New(http.StatusNotFound, "No teacher was found.")
	}

	commonStudents := teachers[0].Students
	for _, teacher := range teachers[1:] {
		commonStudents = getCommonStudents(commonStudents, teacher.Students)
	}

	resp := viewmodels.CommonStudentsResponse{Students: []string{}}
	for _, student := range commonStudents {
		resp.Students = append(resp.Students, student.Email)
	}

	return &resp, nil
}

func getCommonStudents(students []*models.Student, students2 []*models.Student) []*models.Student {
	var commonStudents []*models.Student

	for _, student := range students {
		for _, student2 := range students2 {
			if student.Email == student2.Email {
				commonStudents = append(commonStudents, student)
			}
		}
	}

	return commonStudents
}
