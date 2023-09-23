package dataaccess

import (
	"github.com/sivayogasubramanian/ocv/src/models"
	"gorm.io/gorm"
)

func CreateTeacher(db *gorm.DB, teacher *models.Teacher) error {
	if err := db.Create(&teacher).Error; err != nil {
		return err
	}
	return nil
}

func DoesTeacherExists(db *gorm.DB, email string) (bool, error) {
	exists := false

	err := db.Model(&models.Teacher{}).
		Select("count(*) > 0").
		Where("email = ?", email).
		Find(&exists).
		Error
	if err != nil {
		return exists, err
	}

	return exists, nil
}

func FindAllTeachers(db *gorm.DB, emails []string) ([]models.Teacher, error) {
	var teachers []models.Teacher

	err := db.Where("email IN (?)", emails).Preload("Students").Find(&teachers).Error
	if err != nil {
		return teachers, err
	}

	return teachers, nil
}

func FindTeacher(db *gorm.DB, email string) (models.Teacher, error) {
	var teacher models.Teacher

	err := db.Where("email = ?", email).Preload("Students").Find(&teacher).Error
	if err != nil {
		return teacher, err
	}

	return teacher, nil
}
