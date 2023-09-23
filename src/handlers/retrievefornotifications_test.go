package handlers

import (
	"github.com/sivayogasubramanian/ocv/src/config"
	models2 "github.com/sivayogasubramanian/ocv/src/models"
	"github.com/sivayogasubramanian/ocv/src/viewmodels"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRetrieveNotifications(t *testing.T) {
	type args struct {
		req *viewmodels.RetrieveNotificationsRequest
	}
	tests := []struct {
		name    string
		args    args
		want    *viewmodels.RetrieveNotificationsResponse
		wantErr ocverrs.Error
	}{
		{
			name: "Invalid teacher email",
			args: args{
				req: &viewmodels.RetrieveNotificationsRequest{
					Teacher:      "t1",
					Notification: "Hello students!",
				},
			},
			want:    nil,
			wantErr: ocverrs.New(http.StatusBadRequest, "invalid email: t1"),
		},
		{
			name: "Invalid student email",
			args: args{
				req: &viewmodels.RetrieveNotificationsRequest{
					Teacher:      "t1@gmail.com",
					Notification: "Hello students! @s1",
				},
			},
			want:    nil,
			wantErr: ocverrs.New(http.StatusBadRequest, "invalid email: s1"),
		},
		{
			name: "Teacher does not exist",
			args: args{
				req: &viewmodels.RetrieveNotificationsRequest{
					Teacher:      "t1@gmail.com",
					Notification: "Hello students!",
				},
			},
			want:    nil,
			wantErr: ocverrs.New(http.StatusNotFound, "Teacher with email: t1@gmail.com does not exist."),
		},
		{
			name: "Student does not exist",
			args: args{
				req: &viewmodels.RetrieveNotificationsRequest{
					Teacher:      "t2@gmail.com",
					Notification: "Hello students! @s1@gmail.com",
				},
			},
			want:    nil,
			wantErr: ocverrs.New(http.StatusNotFound, "Student with email: s1@gmail.com does not exist."),
		},
		{
			name: "Suspended students are not included in the response",
			args: args{
				req: &viewmodels.RetrieveNotificationsRequest{
					Teacher:      "t2@gmail.com",
					Notification: "Hello students!",
				},
			},
			want: &viewmodels.RetrieveNotificationsResponse{
				Recipients: []string{"s2@gmail.com"},
			},
			wantErr: nil,
		},
		{
			name: "Students in mentions are included in the response",
			args: args{
				req: &viewmodels.RetrieveNotificationsRequest{
					Teacher:      "t3@gmail.com",
					Notification: "Hello students! @s2@gmail.com",
				},
			},
			want: &viewmodels.RetrieveNotificationsResponse{
				Recipients: []string{"s2@gmail.com"},
			},
			wantErr: nil,
		},
		{
			name: "Suspended students in mentions are not included in the response",
			args: args{
				req: &viewmodels.RetrieveNotificationsRequest{
					Teacher:      "t3@gmail.com",
					Notification: "Hello students! @s3@gmail.com",
				},
			},
			want: &viewmodels.RetrieveNotificationsResponse{
				Recipients: []string{},
			},
			wantErr: nil,
		},
		{
			name: "Students from teacher and students in mentions are included in the response",
			args: args{
				req: &viewmodels.RetrieveNotificationsRequest{
					Teacher:      "t4@gmail.com",
					Notification: "Hello students! @s2@gmail.com",
				},
			},
			want: &viewmodels.RetrieveNotificationsResponse{
				Recipients: []string{"s2@gmail.com", "s4@gmail.com"},
			},
			wantErr: nil,
		},
	}

	config.InitMemoryDB()

	t2 := models2.Teacher{
		Email: "t2@gmail.com",
		Students: []*models2.Student{
			{
				Email: "s2@gmail.com",
			},
			{
				Email:       "s3@gmail.com",
				IsSuspended: true,
			},
		},
	}

	t3 := models2.Teacher{
		Email: "t3@gmail.com",
	}

	t4 := models2.Teacher{
		Email: "t4@gmail.com",
		Students: []*models2.Student{
			{
				Email: "s4@gmail.com",
			},
		},
	}

	var err error
	err = config.DB.Create(&t2).Error
	assert.Nil(t, err)
	err = config.DB.Create(&t3).Error
	assert.Nil(t, err)
	err = config.DB.Create(&t4).Error
	assert.Nil(t, err)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := RetrieveNotifications(tt.args.req)
			assert.Equalf(t, tt.want, got, "RetrieveNotifications(%v)", tt.args.req)
			assert.Equalf(t, tt.wantErr, got1, "RetrieveNotifications(%v)", tt.args.req)
		})
	}
}

func Test_getEmailsFromNotificationText(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Empty text",
			args: args{
				text: "",
			},
			want: []string{},
		},
		{
			name: "No emails",
			args: args{
				text: "Hello world",
			},
			want: []string{},
		},
		{
			name: "With email",
			args: args{
				text: "Hello world @email@gmail.com",
			},
			want: []string{"email@gmail.com"},
		},
		{
			name: "With multiple emails",
			args: args{
				text: "Hello world @email@gmail.com @email2@gmail.com",
			},
			want: []string{"email@gmail.com", "email2@gmail.com"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, getEmailsFromNotificationText(tt.args.text), "getEmailsFromNotificationText(%v)", tt.args.text)
		})
	}
}
