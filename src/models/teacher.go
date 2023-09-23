package models

import (
	"fmt"
	ocvutils "github.com/sivayogasubramanian/ocv/src/utils"
)

type Teacher struct {
	Email    string     `json:"email" gorm:"primary_key"`
	Students []*Student `json:"students" gorm:"many2many:teacher_student;"`
}

func (t *Teacher) Validate() error {
	if !ocvutils.IsValidEmail(t.Email) {
		return fmt.Errorf("invalid email: %s", t.Email)
	}
	return nil
}
