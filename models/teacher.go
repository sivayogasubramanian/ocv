package models

import (
	"errors"
	"fmt"
	ocvutils "github.com/sivayogasubramanian/ocv/utils"
)

type Teacher struct {
	Email    string     `json:"email" gorm:"primary_key"`
	Students []*Student `json:"students" gorm:"many2many:teacher_student;"`
}

func (t *Teacher) Validate() error {
	if !ocvutils.IsValidEmail(t.Email) {
		return errors.New(fmt.Sprintf("Invalid email: %s", t.Email))
	}
	return nil
}
