package dataaccess

import (
	"github.com/sivayogasubramanian/ocv/config"
	"github.com/sivayogasubramanian/ocv/models"
)

func DoesStudentExists(email string) (bool, error) {
	exists := false

	err := config.DB.Model(&models.Student{}).
		Select("count(*) > 0").
		Where("email = ?", email).
		Find(&exists).
		Error
	if err != nil {
		return exists, err
	}

	return exists, nil
}

func SuspendStudent(student *models.Student) error {
	student.IsSuspended = true
	err := config.DB.Save(&student).Error
	if err != nil {
		return err
	}
	return nil
}

func IsStudentSuspended(student *models.Student) (bool, error) {
	err := config.DB.First(&student).Error
	if err != nil {
		return false, err
	}

	return student.IsSuspended, nil
}
