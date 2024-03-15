package models

import (
	"fmt"
	"github.com/RediSearch/redisearch-go/redisearch"
	"log"
	"net/url"
	"os"
	"slices"
)

const scoreBoardIndexName = "idx:scoreboard"

var ScoreBoardIndex redisearch.Client

func connectToRedisSearch() {
	redisURL := os.Getenv("REDIS_URL")

	// Parse the Redis URL
	parsedURL, err := url.Parse(redisURL)
	if err != nil {
		log.Fatal("Error parsing Redis URL")
	}

	redisAddr := fmt.Sprintf("%s:%s", parsedURL.Hostname(), parsedURL.Port())

	client := redisearch.NewClient(redisAddr, scoreBoardIndexName)

	schema := redisearch.NewSchema(redisearch.DefaultOptions).
		AddField(redisearch.NewTextFieldOptions("username", redisearch.TextFieldOptions{Weight: 3, Sortable: false})).
		AddField(redisearch.NewTextFieldOptions("avatar", redisearch.TextFieldOptions{Weight: 0.5, Sortable: false})).
		AddField(redisearch.NewNumericFieldOptions("total_tasks", redisearch.NumericFieldOptions{Sortable: true}))

	list, listErr := client.List()
	if listErr != nil {
		log.Fatal("create index error")
	}

	if !slices.Contains(list, scoreBoardIndexName) {
		if err := client.CreateIndex(schema); err != nil {
			log.Fatal("create index error")
		}
	}

	ScoreBoardIndex = *client
}
