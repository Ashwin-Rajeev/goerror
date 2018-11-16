package goerror

import (
	"fmt"
	"testing"
)

func TestGetErrorInfo(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "normal case",
			args: args{
				err: fmt.Errorf("New error"),
			},
			wantErr: true,
		},
		{
			name:    "failure case",
			args:    args{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := GetErrorInfo(tt.args.err)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetErrorInfo() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNew(t *testing.T) {
	type args struct {
		err string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "normal case",
			args: args{
				err: "New error",
			},
			wantErr: true,
		},
		{
			name:    "No arguments case",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := New(tt.args.err); (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
