package models

import Config "github.com/sivayogasubramanian/ocv/config"

type Teacher struct {
	Email    string     `json:"email" gorm:"primary_key"`
	Students []*Student `json:"students" gorm:"many2many:teacher_student;"`
}

func CreateTeacher(teacher *Teacher) error {
	if err := Config.DB.Create(&teacher).Error; err != nil {
		return err
	}
	return nil
}
