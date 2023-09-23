package handlers

import (
	"github.com/sivayogasubramanian/ocv/src/config"
	"github.com/sivayogasubramanian/ocv/src/viewmodels"
	"net/http"
	"reflect"
	"testing"
)

func TestRegister(t *testing.T) {
	type args struct {
		req *viewmodels.RegisterRequest
	}
	tests := []struct {
		name string
		args args
		want ocverrs.Error
	}{
		{
			name: "Register a teacher with valid email and students",
			args: args{
				req: &viewmodels.RegisterRequest{
					Teacher:  "teacher1@gmail.com",
					Students: []string{"s1@gmail.com", "s2@gmail.com"},
				},
			},
			want: nil,
		},
		{
			name: "Register the same teacher with valid email and students",
			args: args{
				req: &viewmodels.RegisterRequest{
					Teacher:  "teacher1@gmail.com",
					Students: []string{"s1@gmail.com", "s2@gmail.com"},
				},
			},
			want: ocverrs.New(http.StatusConflict, "Teacher with email: teacher1@gmail.com already exists."),
		},
		{
			name: "Register new teacher with valid email and same students",
			args: args{
				req: &viewmodels.RegisterRequest{
					Teacher:  "teacher2@gmail.com",
					Students: []string{"s1@gmail.com", "s2@gmail.com"},
				},
			},
			want: nil,
		},
	}

	config.InitMemoryDB()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Register(tt.args.req); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Register() = %v, want %v", got, tt.want)
			}
		})
	}
}
