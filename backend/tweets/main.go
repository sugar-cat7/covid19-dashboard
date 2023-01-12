package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Tweets struct {
	CreatedAt time.Time `db:"created_at" json:"createdAt"`
	TweetId   string    `db:"tweet_id" json:"tweetId"`
	Username  string    `db:"username" json:"username"`
	Content   string    `db:"content" json:"text"`
	LangCode  string    `db:"lang_code" json:"lang_code"`
}

func time2str(t time.Time) string {
	return t.Format("2006-01-02")
}

func parseQueryParams(request events.APIGatewayProxyRequest) (to, from, limit string, countryCode []string) {
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		panic(err.Error())
	}
	to = time2str(time.Now().In(jst))
	from = time2str(time.Now().In(jst).AddDate(0, -1, 0))

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
	if v, ok := request.QueryStringParameters["limit"]; ok {
		limit = v
	}
	return to, from, limit, countryCode
}

func fetchTweets(to, from, limit string, countryCode []string) ([]Tweets, error) {
	db, err := sqlx.Open("mysql", "xxx")
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	tweets := []Tweets{}
	query := `
			SELECT
			tweets.tweet_id, tweets.username, tweets.content, tweets.lang_code, tweets.created_at
			FROM
				tweets
			LEFT JOIN languages on languages.language_code = tweets.lang_code
			LEFT JOIN country_language on languages.id = country_language.language_id
			LEFT JOIN country on country_language.country_id = country.id
			WHERE tweets.created_at BETWEEN :from AND :to
			`
	if len(countryCode) > 0 {
		query += " AND country.country_code IN (:codeList)"
	}
	if limit != "" {
		query += "ORDER BY tweets.created_at LIMIT :limit"
	} else {
		query += "ORDER BY tweets.created_at LIMIT 100"
	}
	input := map[string]interface{}{
		"from":     from,
		"to":       to,
		"codeList": countryCode,
		"limit":    limit,
	}
	query, args, err := sqlx.Named(query, input)
	if err != nil {
		return nil, err
	}
	if len(countryCode) > 0 {
		query, args, err = sqlx.In(query, args...)
		if err != nil {
			panic(err)
		}
	}
	query = db.Rebind(query)
	err = db.Select(&tweets, query, args...)
	if err != nil {
		return nil, err
	}
	fmt.Println(tweets)
	return tweets, nil
}

var headers = map[string]string{
	"Content-Type":                 "application/json",
	"Access-Control-Allow-Origin":  "*",
	"Access-Control-Allow-Methods": "GET",
	"Access-Control-Allow-Headers": "Origin,Authorization,Accept,X-Requested-With",
}

func buildResponse(tweets []Tweets) events.APIGatewayProxyResponse {

	b, err := json.Marshal(tweets)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Headers:    headers,
			Body:       err.Error(),
		}
	}
	return events.APIGatewayProxyResponse{
		StatusCode:      http.StatusOK,
		Headers:         headers,
		Body:            string(b),
		IsBase64Encoded: false,
	}
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	to, from, limit, countryCode := parseQueryParams(request)
	tweets, err := fetchTweets(to, from, limit, countryCode)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Headers:    headers,
			Body:       err.Error(),
		}, nil
	}
	return buildResponse(tweets), nil
}

func main() {
	lambda.Start(handler)
}
