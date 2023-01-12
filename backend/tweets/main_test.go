package main

import (
	"reflect"
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

func Test_parseQueryParams(t *testing.T) {
	type args struct {
		request events.APIGatewayProxyRequest
	}
	tests := []struct {
		name            string
		args            args
		wantTo          string
		wantFrom        string
		wantLimit       string
		wantCountryCode []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotTo, gotFrom, gotLimit, gotCountryCode := parseQueryParams(tt.args.request)
			if gotTo != tt.wantTo {
				t.Errorf("parseQueryParams() gotTo = %v, want %v", gotTo, tt.wantTo)
			}
			if gotFrom != tt.wantFrom {
				t.Errorf("parseQueryParams() gotFrom = %v, want %v", gotFrom, tt.wantFrom)
			}
			if gotLimit != tt.wantLimit {
				t.Errorf("parseQueryParams() gotLimit = %v, want %v", gotLimit, tt.wantLimit)
			}
			if !reflect.DeepEqual(gotCountryCode, tt.wantCountryCode) {
				t.Errorf("parseQueryParams() gotCountryCode = %v, want %v", gotCountryCode, tt.wantCountryCode)
			}
		})
	}
}

func Test_fetchTweets(t *testing.T) {
	tt := time.Date(2023, 1, 7, 0, 13, 14, 0, time.UTC)
	type args struct {
		to          string
		from        string
		limit       string
		countryCode []string
	}
	tests := []struct {
		name    string
		args    args
		want    []Tweets
		wantErr bool
	}{
		{name: "case1", args: args{
			from:        "2023-01-07",
			to:          "2023-01-08",
			countryCode: []string{"JPN"},
			limit:       "1",
		}, want: []Tweets{{
			CreatedAt: tt,
			TweetId:   "1611516272193449991",
			Username:  "waizuestate",
			Content:   "食品関連事業者の倒産は577件～新型コロナ関連倒産4845件～(帝国データバンク)\n#Yahooニュース\nhttps://t.co/FuSDVch6KS",
			LangCode:  "ja",
		}}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := fetchTweets(tt.args.to, tt.args.from, tt.args.limit, tt.args.countryCode)
			if (err != nil) != tt.wantErr {
				t.Errorf("fetchTweets() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fetchTweets() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_buildResponse(t *testing.T) {
	type args struct {
		tweets []Tweets
	}
	tests := []struct {
		name string
		args args
		want events.APIGatewayProxyResponse
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buildResponse(tt.args.tweets); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("buildResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_handler(t *testing.T) {
	type args struct {
		request events.APIGatewayProxyRequest
	}
	tests := []struct {
		name    string
		args    args
		want    events.APIGatewayProxyResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := handler(tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("handler() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("handler() = %v, want %v", got, tt.want)
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
