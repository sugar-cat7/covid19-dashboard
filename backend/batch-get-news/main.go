package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"

	"github.com/jmoiron/sqlx"
)

type NewsResponse struct {
	Status       string `json:"status"`
	TotalResults int    `json:"totalResults"`
	Articles     []struct {
		Source struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"source"`
		Author      string    `json:"author"`
		Title       string    `json:"title"`
		Description string    `json:"description"`
		URL         string    `json:"url"`
		URLToImage  string    `json:"urlToImage"`
		PublishedAt time.Time `json:"publishedAt"`
		Content     string    `json:"content"`
	} `json:"articles"`
}

type News struct {
	SearchId    string    `db:"search_id"`
	CountryCode string    `db:"country_code"`
	Url         string    `db:"url"`
	UrlToImage  string    `db:"url_to_image"`
	Title       string    `db:"title"`
	PublishedAt time.Time `db:"published_at"`
	SourceId    string    `db:"source_id"`
}

type NewsTotal struct {
	Id         string    `db:"id"`
	TotalNews  string    `db:"total_news"`
	SearchFrom time.Time `db:"search_from"`
	SearchTo   time.Time `db:"search_to"`
}

func time2str(t time.Time) string {

	return t.Format("2006-01-02")
}

var (
	url = "https://newsapi.org/v2/everything?q=covid"
)

func handler(ctx context.Context, request events.APIGatewayProxyRequest) error {
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		panic(err)
	}
	nowJST := time2str(time.Now().In(jst).AddDate(0, 0, -1))
	prevJST := time2str(time.Now().In(jst).AddDate(0, 0, -7))
	resp, err := http.Get(url + "&from=" + prevJST + "&to=" + nowJST + "&sortBy=popularity&apiKey=xxx")
	if err != nil {
		panic(err.Error())
	}
	defer resp.Body.Close()
	byteArray, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err.Error())
	}
	var NewsResponse NewsResponse

	if err := json.Unmarshal(byteArray, &NewsResponse); err != nil {
		return err
	}
	uuidObj, _ := uuid.NewUUID()
	searchFrom, _ := time.Parse("2006-01-02", prevJST)
	SearchTo, _ := time.Parse("2006-01-02", nowJST)
	newsTotal := NewsTotal{
		Id:         uuidObj.String(),
		TotalNews:  strconv.Itoa(NewsResponse.TotalResults),
		SearchFrom: searchFrom,
		SearchTo:   SearchTo,
	}
	news := make([]News, len(NewsResponse.Articles))
	for i, n := range NewsResponse.Articles {
		news[i] = News{
			SearchId:    uuidObj.String(),
			CountryCode: "jp",
			Title:       n.Title,
			Url:         n.URL,
			UrlToImage:  n.URLToImage,
			PublishedAt: n.PublishedAt,
		}
	}
	if len(news) == 0 {
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
	query := `INSERT INTO covid_news_total (id, total_news, search_from, search_to) VALUES (:id, :total_news, :search_from, :search_to);`

	_, err = tx.NamedExec(query, newsTotal)
	if err != nil {
		return err
	}
	query = `INSERT INTO covid_news (search_id, country_code, url, url_to_image, title, published_at, source_id) VALUES (:search_id, :country_code, :url, :url_to_image, :title, :published_at, :source_id);`
	sliceSize := len(news)

	for i := 0; i < sliceSize; i += 10000 {
		end := i + 10000
		if sliceSize < end {
			end = sliceSize
		}
		_, err = tx.NamedExec(query, news[i:end])
		if err != nil {
			fmt.Println(err)
			return err
		}
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
