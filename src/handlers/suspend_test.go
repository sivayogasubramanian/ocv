package handlers

import (
	"github.com/sivayogasubramanian/ocv/src/config"
	ocverrs "github.com/sivayogasubramanian/ocv/src/errors"
	"github.com/sivayogasubramanian/ocv/src/models"
	"github.com/sivayogasubramanian/ocv/src/viewmodels"
	"github.com/stretchr/testify/assert"
	"net/http"
	"reflect"
	"testing"
)

func TestSuspend(t *testing.T) {
	type args struct {
		req *viewmodels.SuspendRequest
	}
	tests := []struct {
		name string
		args args
		want ocverrs.Error
	}{
		{
			name: "Invalid email",
			args: args{
				req: &viewmodels.SuspendRequest{
					Student: "student",
				},
			},
			want: ocverrs.New(http.StatusBadRequest, "invalid email: student"),
		},
		{
			name: "Student does not exists",
			args: args{
				req: &viewmodels.SuspendRequest{
					Student: "student@gmail.com",
				},
			},
			want: ocverrs.New(http.StatusNotFound, "Student with email: student@gmail.com does not exist."),
		},
		{
			name: "Student with valid email",
			args: args{
				req: &viewmodels.SuspendRequest{
					Student: "s1@gmail.com",
				},
			},
			want: nil,
		},
	}

	db := config.InitMemoryDB()
	s1 := models.Student{
		Email:       "s1@gmail.com",
		IsSuspended: false,
		Teachers:    nil,
	}
	err := db.Save(&s1).Error
	assert.Nil(t, err)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Suspend(db, tt.args.req); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Suspend() = %v, want %v", got, tt.want)
			}
		})
	}
}
