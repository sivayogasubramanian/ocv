package ocvutils

import "testing"

func TestIsValidEmail(t *testing.T) {
	type args struct {
		email string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "valid",
			args: args{email: "testing@gmail.com"},
			want: true,
		},
		{
			name: "valid",
			args: args{email: "a@a.com"},
			want: true,
		},
		{
			name: "invalid",
			args: args{email: "testing.gmail.com"},
			want: false,
		},
		{
			name: "invalid",
			args: args{email: "111"},
			want: false,
		},
		{
			name: "invalid",
			args: args{email: ""},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValidEmail(tt.args.email); got != tt.want {
				t.Errorf("IsValidEmail() = %v, want %v", got, tt.want)
			}
		})
	}
}
