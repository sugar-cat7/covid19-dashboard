package main

import (
	"context"
	"testing"
	"time"

	"github.com/aws/aws-lambda-go/events"
	_ "github.com/go-sql-driver/mysql"
)

func Test_time2str(t *testing.T) {
	type args struct {
		t time.Time
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := time2str(tt.args.t); got != tt.want {
				t.Errorf("time2str() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_handler(t *testing.T) {
	type args struct {
		ctx     context.Context
		request events.APIGatewayProxyRequest
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := handler(tt.args.ctx, tt.args.request); (err != nil) != tt.wantErr {
				t.Errorf("handler() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}
