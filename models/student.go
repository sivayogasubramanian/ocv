package models

import (
	"fmt"
	ocvutils "github.com/sivayogasubramanian/ocv/utils"
)

type Student struct {
	Email       string     `json:"email" gorm:"primary_key"`
	IsSuspended bool       `json:"isSuspended" gorm:"default:false"`
	Teachers    []*Teacher `json:"teachers" gorm:"many2many:teacher_student;"`
}

func (s *Student) Validate() error {
	if !ocvutils.IsValidEmail(s.Email) {
		return fmt.Errorf("invalid email: %s", s.Email)
	}
	return nil
}
