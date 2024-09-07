package sms

import (
	"testing"
)

func TestSend(t *testing.T) {
	type args struct {
		phoneNum []string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "normal",
			args:    args{[]string{"13164661907"}},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := Send(tt.args.phoneNum...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Send() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
