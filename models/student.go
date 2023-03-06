package models

import Config "github.com/sivayogasubramanian/ocv/config"

type Student struct {
	Email    string     `json:"email" gorm:"primary_key"`
	Teachers []*Teacher `json:"teachers" gorm:"many2many:teacher_student;"`
}

func CreateStudent(student *Student) error {
	if err := Config.DB.Create(&student).Error; err != nil {
		return err
	}
	return nil
}
