package dataaccess

import (
	"github.com/sivayogasubramanian/ocv/src/models"
	"gorm.io/gorm"
)

func DoesStudentExists(db *gorm.DB, email string) (bool, error) {
	exists := false

	err := db.Model(&models.Student{}).
		Select("count(*) > 0").
		Where("email = ?", email).
		Find(&exists).
		Error
	if err != nil {
		return exists, err
	}

	return exists, nil
}

func SuspendStudent(db *gorm.DB, student *models.Student) error {
	student.IsSuspended = true
	err := db.Save(&student).Error
	if err != nil {
		return err
	}
	return nil
}

func IsStudentSuspended(db *gorm.DB, student *models.Student) (bool, error) {
	err := db.First(&student).Error
	if err != nil {
		return false, err
	}

	return student.IsSuspended, nil
}
