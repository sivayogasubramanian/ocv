package models

import (
	"errors"
	"fmt"
	ocvutils "github.com/sivayogasubramanian/ocv/utils"
)

type Student struct {
	Email    string     `json:"email" gorm:"primary_key"`
	Teachers []*Teacher `json:"teachers" gorm:"many2many:teacher_student;"`
}

func (s *Student) Validate() error {
	if !ocvutils.IsValidEmail(s.Email) {
		return errors.New(fmt.Sprintf("Invalid email: %s", s.Email))
	}
	return nil
}
