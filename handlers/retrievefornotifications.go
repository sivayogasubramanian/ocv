package handlers

import (
	"fmt"
	"github.com/sivayogasubramanian/ocv/dataaccess"
	ocverrs "github.com/sivayogasubramanian/ocv/errors"
	"github.com/sivayogasubramanian/ocv/models"
	"github.com/sivayogasubramanian/ocv/viewmodels"
	"net/http"
	"sort"
	"strings"
)

func RetrieveNotifications(req *viewmodels.RetrieveNotificationsRequest) (*viewmodels.RetrieveNotificationsResponse, ocverrs.Error) {
	err := verifyRequestParams(req)
	if err != nil {
		return nil, err
	}

	teacherExists, stderr := dataaccess.DoesTeacherExists(req.Teacher)
	if stderr != nil {
		return nil, ocverrs.New(http.StatusInternalServerError, "An error occurred while retrieving notification recipients.")
	}
	if !teacherExists {
		return nil, ocverrs.New(http.StatusNotFound, fmt.Sprintf("Teacher with email: %s does not exist.", req.Teacher))
	}

	// Set is used to avoid duplicate student emails
	studentRecipients := make(map[string]bool)

	studentEmails := getEmailsFromNotificationText(req.Notification)
	for _, studentEmail := range studentEmails {
		studentExists, err := dataaccess.DoesStudentExists(studentEmail)
		if err != nil {
			return nil, ocverrs.New(http.StatusInternalServerError, "An error occurred while retrieving notification recipients.")
		}
		if !studentExists {
			return nil, ocverrs.New(http.StatusNotFound, fmt.Sprintf("Student with email: %s does not exist.", studentEmail))
		}

		studentRecipients[studentEmail] = true
	}

	teacher, stderr := dataaccess.FindTeacher(req.Teacher)
	if stderr != nil {
		return nil, ocverrs.New(http.StatusInternalServerError, "An error occurred while retrieving notification recipients.")
	}
	for _, student := range teacher.Students {
		studentRecipients[student.Email] = true
	}

	var recipients []string

	for studentEmail := range studentRecipients {
		student := models.Student{Email: studentEmail}

		isSuspended, err := dataaccess.IsStudentSuspended(&student)
		if err != nil {
			return nil, ocverrs.New(http.StatusInternalServerError, "An error occurred while retrieving notification recipients.")
		}
		if isSuspended {
			continue
		}

		recipients = append(recipients, studentEmail)
	}

	if len(recipients) == 0 {
		recipients = []string{}
	}

	sort.Strings(recipients)

	return &viewmodels.RetrieveNotificationsResponse{
		Recipients: recipients,
	}, nil
}

func verifyRequestParams(req *viewmodels.RetrieveNotificationsRequest) ocverrs.Error {
	teacher := models.Teacher{Email: req.Teacher}
	if err := teacher.Validate(); err != nil {
		return ocverrs.New(http.StatusBadRequest, err.Error())
	}

	studentEmails := getEmailsFromNotificationText(req.Notification)
	for _, studentEmail := range studentEmails {
		student := models.Student{Email: studentEmail}
		if err := student.Validate(); err != nil {
			return ocverrs.New(http.StatusBadRequest, err.Error())
		}
	}

	return nil
}

func getEmailsFromNotificationText(text string) []string {
	if len(text) == 0 {
		return []string{}
	}

	words := strings.Split(text, " ")

	var emails []string
	for _, word := range words {
		if word[0] == '@' {
			emails = append(emails, word[1:])
		}
	}

	if len(emails) == 0 {
		return []string{}
	}

	return emails
}
