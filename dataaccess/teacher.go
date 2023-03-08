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

func DoesTeacherExists(email string) (bool, error) {
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

func FindAllTeachers(emails []string) ([]models.Teacher, error) {
	var teachers []models.Teacher

	err := config.DB.Where("email IN (?)", emails).Preload("Students").Find(&teachers).Error
	if err != nil {
		return teachers, err
	}

	return teachers, nil
}

func FindTeacher(email string) (models.Teacher, error) {
	var teacher models.Teacher

	err := config.DB.Where("email = ?", email).Preload("Students").Find(&teacher).Error
	if err != nil {
		return teacher, err
	}

	return teacher, nil
}
