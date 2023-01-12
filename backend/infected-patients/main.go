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

type CountryInfectedCountry struct {
	Date        time.Time `db:"date" json:"publishedAt"`
	InfectedNum int       `db:"infected_num" json:"infectedNum"`
	DeceasedNum int       `db:"deceased_num" json:"deceasedNum"`
	Country     Country   `db:"country" json:"country"`
}

type Country struct {
	CountryName string `db:"country_name" json:"countryName"`
	CountryCode string `db:"country_code" json:"countryCode"`
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

	countries := []CountryInfectedCountry{}
	stmt := `
			SELECT
			covid_infected_country.date,
			covid_infected_country.infected_num,
			covid_infected_country.deceased_num,
			CASE
			WHEN country.country_name is null THEN 'その他'
			ELSE country.country_name
			END "country.country_name",
			CASE
				WHEN country.country_code is null THEN 'ZZZ'
				ELSE country.country_code
			END "country.country_code"
			FROM
				covid_infected_country
			LEFT JOIN country on
				country.country_name = covid_infected_country.data_name
			WHERE
				covid_infected_country.date BETWEEN :from AND :to
			`
	if len(countryCode) > 0 {
		stmt += " AND country.country_code IN (:codeList)"
	}
	stmt += " ORDER BY covid_infected_country.date"
	input := map[string]interface{}{
		"from":     from,
		"to":       to,
		"codeList": countryCode,
	}
	stmt, args, err := sqlx.Named(stmt, input)
	if err != nil {
		panic(err)
	}
	if len(countryCode) > 0 {
		stmt, args, err = sqlx.In(stmt, args...)
		if err != nil {
			panic(err)
		}
	}
	stmt = db.Rebind(stmt)
	err = db.Select(&countries, stmt, args...)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	jsonResult, err := json.Marshal(countries)
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
