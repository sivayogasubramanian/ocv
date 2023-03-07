package dataaccess

import (
	"github.com/sivayogasubramanian/ocv/config"
	"github.com/sivayogasubramanian/ocv/models"
)

func CreateTeacher(teacher *models.Teacher) error {
	if err := config.DB.Create(&teacher).Error; err != nil {
		return err
	}
	return nil
}

func CheckIfTeacherEmailTaken(email string) (bool, error) {
	exists := false

	err := config.DB.Model(&models.Teacher{}).
		Select("count(*) > 0").
		Where("email = ?", email).
		Find(&exists).
		Error
	if err != nil {
		return exists, err
	}

	return exists, nil
}
