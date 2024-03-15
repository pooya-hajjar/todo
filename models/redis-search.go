package models

import (
	"fmt"
	"github.com/RediSearch/redisearch-go/redisearch"
	"log"
	"net/url"
	"os"
	"slices"
)

var UsersIndex redisearch.Client

func connectToRedisSearch() {
	redisURL := os.Getenv("REDIS_URL")

	// Parse the Redis URL
	parsedURL, err := url.Parse(redisURL)
	if err != nil {
		log.Fatal("Error parsing Redis URL")
	}

	redisAddr := fmt.Sprintf("%s:%s", parsedURL.Hostname(), parsedURL.Port())

	client := redisearch.NewClient(redisAddr, "idx:users")

	schema := redisearch.NewSchema(redisearch.DefaultOptions).
		AddField(redisearch.NewTextFieldOptions("username", redisearch.TextFieldOptions{Weight: 3.0, Sortable: true})).
		AddField(redisearch.NewTextFieldOptions("avatar", redisearch.TextFieldOptions{Weight: 1.0})).
		AddField(redisearch.NewNumericFieldOptions("status", redisearch.NumericFieldOptions{Sortable: true})).
		AddField(redisearch.NewNumericFieldOptions("total_tasks", redisearch.NumericFieldOptions{Sortable: true})).
		AddField(redisearch.NewNumericFieldOptions("today_tasks", redisearch.NumericFieldOptions{Sortable: true}))

	list, listErr := client.List()
	if listErr != nil {
		log.Fatal("create index error")
	}

	if !slices.Contains(list, "idx:users") {
		if err := client.CreateIndex(schema); err != nil {
			log.Fatal("create index error")
		}
	}

	UsersIndex = *client
}
