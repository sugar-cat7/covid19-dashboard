package main

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type OpenAPIResponse struct {
	Data []struct {
		EditHistoryTweetIds []string  `json:"edit_history_tweet_ids"`
		CreatedAt           time.Time `json:"created_at"`
		ID                  string    `json:"id"`
		AuthorID            string    `json:"author_id"`
		Text                string    `json:"text"`
		Lang                string    `json:"lang"`
	} `json:"data"`
	Includes struct {
		Users []struct {
			ID       string `json:"id"`
			Name     string `json:"name"`
			Username string `json:"username"`
		} `json:"users"`
	} `json:"includes"`
	Meta struct {
		NewestID    string `json:"newest_id"`
		OldestID    string `json:"oldest_id"`
		ResultCount int    `json:"result_count"`
		NextToken   string `json:"next_token"`
	} `json:"meta"`
}

type Tweets struct {
	CreatedAt time.Time `db:"created_at" json:"createdAt"`
	TweetId   string    `db:"tweet_id" json:"tweetId"`
	Username  string    `db:"username" json:"username"`
	Content   string    `db:"content" json:"text"`
	LangCode  string    `db:"lang_code" json:"lang_code"`
}

func time2str(t time.Time) string {

	return t.Format("20060102")
}

var (
	url = "https://api.twitter.com/2/tweets/search/recent?query=covid-19+OR+covid19+OR+%E3%82%B3%E3%83%AD%E3%83%8A&expansions=author_id&tweet.fields=created_at,lang&media.fields=url&place.fields=country_code&max_results=100"
)

func handler(ctx context.Context, request events.APIGatewayProxyRequest) error {
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", "xxx")
	client := new(http.Client)
	resp, err := client.Do(req)
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
	tweets := make([]Tweets, len(OpenAPIResponse.Data))
	for i, t := range OpenAPIResponse.Data {
		var username string
		for _, u := range OpenAPIResponse.Includes.Users {
			if u.ID == t.AuthorID {
				username = u.Username
				break
			}
		}
		tweets[i] = Tweets{
			CreatedAt: t.CreatedAt,
			TweetId:   t.ID,
			Username:  username,
			Content:   t.Text,
			LangCode:  t.Lang,
		}
	}
	if len(tweets) == 0 {
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
	query := `INSERT INTO tweets (tweet_id, username, content, lang_code, created_at) VALUES (:tweet_id, :username, :content, :lang_code, :created_at);`
	_, err = tx.NamedExec(query, tweets)
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
