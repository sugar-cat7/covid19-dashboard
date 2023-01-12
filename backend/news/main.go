package main

import (
	"encoding/json"
	"strings"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type News struct {
	SearchId    string     `db:"search_id" json:"searchId"`
	Url         string     `db:"url" json:"url"`
	UrlToImage  string     `db:"url_to_image" json:"urlToImage"`
	Title       string     `db:"title" json:"title"`
	PublishedAt time.Time  `db:"published_at" json:"publishedAt"`
	NewsSource  NewsSource `db:"covid_news_source" json:"source"`
}

type NewsSource struct {
	SourceName *string `db:"source_name"`
	Language   *string `db:"language"`
	Country    *string `db:"country"`
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
	countryCode := []string{}
	if v, ok := request.QueryStringParameters["to"]; ok {
		t, _ := time.Parse("2006-01-02", v)
		to = time2str(t)
	}
	if v, ok := request.QueryStringParameters["from"]; ok {
		t, _ := time.Parse("2006-01-02", v)
		from = time2str(t)
	}
	if v, ok := request.QueryStringParameters["country_code"]; ok {
		countryCode = strings.Split(v, ",")
	}

	db, err := sqlx.Open("mysql", "xxx")
	if err != nil {
		panic(err.Error())
	}
	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	news := []News{}
	query := `
		SELECT
			covid_news.search_id,
			covid_news.url,
			covid_news.url_to_image,
			covid_news.title,
			covid_news.published_at,
			covid_news_source.source_name "covid_news_source.source_name",
			covid_news_source.country "covid_news_source.country",
			covid_news_source.language "covid_news_source.language"
		FROM
			covid_news
		LEFT JOIN covid_news_source on
			covid_news.source_id = covid_news_source.source_id
		LEFT JOIN languages on languages.language_code = covid_news_source.language
		LEFT JOIN country_language on languages.id = country_language.language_id
		LEFT JOIN country on country_language.country_id = country.id
		WHERE covid_news.published_at BETWEEN :from AND :to
		`
	if len(countryCode) > 0 {
		query += " AND country.country_code IN (:codeList)"
	}
	query += "ORDER BY covid_news.published_at LIMIT 20"
	input := map[string]interface{}{
		"from":     from,
		"to":       to,
		"codeList": countryCode,
	}
	query, args, err := sqlx.Named(query, input)
	if err != nil {
		panic(err)
	}
	if len(countryCode) > 0 {
		query, args, err = sqlx.In(query, args...)
		if err != nil {
			panic(err)
		}
	}
	query = db.Rebind(query)
	err = db.Select(&news, query, args...)

	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	jsonResult, err := json.Marshal(news)
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
