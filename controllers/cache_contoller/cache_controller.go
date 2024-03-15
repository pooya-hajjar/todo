package cacheContoller

import (
	"context"
	"fmt"
	"github.com/RediSearch/redisearch-go/redisearch"
	"github.com/pooya-hajjar/todo/models"
	"log"
)

func AddScoreBoardDocument(id, totalTasks int, username, avatar string) {
	DocName := fmt.Sprintf("user:%d", id)
	dc := redisearch.NewDocument(DocName, 1.0)

	dc.Set("username", username).
		Set("avatar", avatar).
		Set("total_tasks", totalTasks)

	err := models.ScoreBoardIndex.Index(dc)
	if err != nil {
		log.Fatal("create doc error")
	}
}

func UpdateScoreBoardDocumentTasks(id, totalTasks int) {
	DocName := fmt.Sprintf("user:%d", id)

	intCmd := models.RedisDb.HSet(context.Background(), DocName, "total_tasks", totalTasks)
	result, err := intCmd.Result()
	if err != nil {
		log.Fatal("redis error")
	}

}

func UpdateScoreBoardDocumentUsername(id int, username string) {
	DocName := fmt.Sprintf("user:%d", id)

	intCmd := models.RedisDb.HSet(context.Background(), DocName, "username", username)
	_, err := intCmd.Result()
	if err != nil {
		log.Fatal("redis error")
	}

}
