package main

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type OpenAPIResponse struct {
	ErrorInfo struct {
		ErrorFlag    string      `json:"errorFlag"`
		ErrorCode    interface{} `json:"errorCode"`
		ErrorMessage interface{} `json:"errorMessage"`
	} `json:"errorInfo"`
	ItemList []struct {
		Date        string `json:"date"`
		DataName    string `json:"dataName"`
		InfectedNum string `json:"infectedNum"`
		DeceasedNum string `json:"deceasedNum"`
	} `json:"itemList"`
}

type Country struct {
	Date        time.Time `db:"date"`
	DataName    string    `db:"data_name"`
	InfectedNum int       `db:"infected_num"`
	DeceasedNum int       `db:"deceased_num"`
}

func time2str(t time.Time) string {

	return t.Format("20060102")
}

var (
	url = "https://opendata.corona.go.jp/api/OccurrenceStatusOverseas"
)

func handler(ctx context.Context, request events.APIGatewayProxyRequest) error {
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		panic(err)
	}
	nowJST := time.Now().In(jst).AddDate(0, 0, -1)
	resp, err := http.Get(url + "?date=" + time2str(nowJST))
	if err != nil {
		panic(err.Error())
	}
	defer resp.Body.Close()
	byteArray, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err.Error())
	}
	var OpenAPIResponse OpenAPIResponse

	if err := json.Unmarshal(byteArray, &OpenAPIResponse); err != nil {
		return err
	}
	countries := make([]Country, len(OpenAPIResponse.ItemList))
	for i, c := range OpenAPIResponse.ItemList {
		t, _ := time.Parse("2006-01-02", strings.Replace(c.Date, "/", "-", -1))
		infectedNum, _ := strconv.Atoi(strings.Replace(strings.Replace(c.InfectedNum, ",", "", -1), " ", "", -1))
		deceasedNum, _ := strconv.Atoi(strings.Replace(strings.Replace(c.DeceasedNum, ",", "", -1), " ", "", -1))
		countries[i] = Country{
			Date:        t,
			DataName:    c.DataName,
			InfectedNum: infectedNum,
			DeceasedNum: deceasedNum,
		}
	}
	if len(countries) == 0 {
		return nil
	}
	db, err := sqlx.Open("mysql", "xxx")
	if err != nil {
		panic(err.Error())
	}
	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}
	tx := db.MustBegin()
	query := `INSERT INTO covid_infected_country (date, data_name, infected_num, deceased_num) VALUES (:date, :data_name, :infected_num, :deceased_num);`
	sliceSize := len(countries)

	for i := 0; i < sliceSize; i += 10000 {
		end := i + 10000
		if sliceSize < end {
			end = sliceSize
		}
		_, err = tx.NamedExec(query, countries[i:end])
	}
	if err != nil {
		return err
	}
	tx.Commit()
	defer db.Close()
	return nil
}

func main() {
	lambda.Start(handler)
}
