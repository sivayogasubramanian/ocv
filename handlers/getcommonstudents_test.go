package handlers

import (
	"github.com/sivayogasubramanian/ocv/config"
	ocverrs "github.com/sivayogasubramanian/ocv/errors"
	"github.com/sivayogasubramanian/ocv/models"
	"github.com/sivayogasubramanian/ocv/viewmodels"
	"github.com/stretchr/testify/assert"
	"net/http"
	"reflect"
	"testing"
)

func TestGetCommonStudents(t *testing.T) {
	type args struct {
		teacherEmails []string
	}
	tests := []struct {
		name    string
		args    args
		want    *viewmodels.CommonStudentsResponse
		wantErr ocverrs.Error
	}{
		{
			name: "Teacher not found",
			args: args{
				teacherEmails: []string{"t@gmail.com"},
			},
			want:    nil,
			wantErr: ocverrs.New(http.StatusNotFound, "No teacher was found."),
		},
		{
			name: "t1 and t2",
			args: args{
				teacherEmails: []string{"t1@gmail.com", "t2@gmail.com"},
			},
			want: &viewmodels.CommonStudentsResponse{
				Students: []string{"s1@gmail.com", "s2@gmail.com"},
			},
			wantErr: nil,
		},
		{
			name: "t1 and t3",
			args: args{
				teacherEmails: []string{"t1@gmail.com", "t3@gmail.com"},
			},
			want: &viewmodels.CommonStudentsResponse{
				Students: []string{"s1@gmail.com"},
			},
			wantErr: nil,
		},
		{
			name: "t2 and t3",
			args: args{
				teacherEmails: []string{"t2@gmail.com", "t3@gmail.com"},
			},
			want: &viewmodels.CommonStudentsResponse{
				Students: []string{"s1@gmail.com"},
			},
			wantErr: nil,
		},
	}

	config.InitMemoryDB()

	t1 := models.Teacher{
		Email: "t1@gmail.com",
		Students: []*models.Student{
			{
				Email: "s1@gmail.com",
			},
			{
				Email: "s2@gmail.com",
			},
			{
				Email: "s3@gmail.com",
			},
		},
	}

	t2 := models.Teacher{
		Email: "t2@gmail.com",
		Students: []*models.Student{
			{
				Email: "s1@gmail.com",
			},
			{
				Email: "s2@gmail.com",
			},
		},
	}

	t3 := models.Teacher{
		Email: "t3@gmail.com",
		Students: []*models.Student{
			{
				Email: "s1@gmail.com",
			},
		},
	}

	var err error
	err = config.DB.Create(&t1).Error
	assert.Nil(t, err)
	err = config.DB.Create(&t2).Error
	assert.Nil(t, err)
	err = config.DB.Create(&t3).Error
	assert.Nil(t, err)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := GetCommonStudents(tt.args.teacherEmails)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCommonStudents() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.wantErr) {
				t.Errorf("GetCommonStudents() got1 = %v, want %v", got1, tt.wantErr)
			}
		})
	}
}
