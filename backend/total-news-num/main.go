package main

import (
	"encoding/json"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type NewsTotal struct {
	Id         string    `db:"id" json:"id"`
	TotalNews  string    `db:"total_news" json:"totalNews"`
	SearchFrom time.Time `db:"search_from" json:"searchFrom"`
	SearchTo   time.Time `db:"search_to" json:"searchTo"`
}

func time2str(t time.Time) string {
	return t.Format("2006-01-02")
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	headers := map[string]string{
		"Content-Type":                 "application/json",
		"Access-Control-Allow-Origin":  "*",
		"Access-Control-Allow-Methods": "GET",
		"Access-Control-Allow-Headers": "Origin,Authorization,Accept,X-Requested-With",
	}
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		panic(err.Error())
	}
	to := time2str(time.Now().In(jst))
	from := time2str(time.Now().In(jst).AddDate(0, -1, 0))
	if v, ok := request.QueryStringParameters["to"]; ok {
		t, _ := time.Parse("2006-01-02", v)
		to = time2str(t)
	}
	if v, ok := request.QueryStringParameters["from"]; ok {
		t, _ := time.Parse("2006-01-02", v)
		from = time2str(t)
	}

	db, err := sqlx.Open("mysql", "xxx")
	if err != nil {
		panic(err.Error())
	}
	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	newsTotal := []NewsTotal{}
	err = db.Select(&newsTotal, "SELECT id, total_news, search_from, search_to FROM covid_news_total WHERE search_from >= ? AND search_to <= ?", from, to)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	jsonResult, err := json.Marshal(newsTotal)
	if err != nil {
		panic(err.Error())
	}
	return events.APIGatewayProxyResponse{
		Headers:    headers,
		Body:       string(jsonResult),
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handler)
}
