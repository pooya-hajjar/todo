package scoreboardController

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/RediSearch/redisearch-go/redisearch"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/pooya-hajjar/todo/constants/query"

	"github.com/gin-gonic/gin"
	"github.com/pooya-hajjar/todo/models"
)

type TopTenUsers struct {
	UserName       any `json:"username"`
	CompletedTasks any `json:"completed_tasks"`
}

func Top10(ctx *gin.Context) {
	err := readFromRedis(ctx)
	if err != nil {
		fmt.Println(err.Error())

		readFromMainDB(ctx)
	}
}

func readFromRedis(ctx *gin.Context) error {
	docs, _, err := models.ScoreBoardIndex.Search(redisearch.NewQuery("*").
		SetReturnFields("username", "avatar", "total_tasks").
		SetSortBy("total_tasks", true).
		Limit(0, 10))

	if err != nil {
		return err
	}

	if len(docs) < 1 {
		return errors.New("there is no documents")
	}

	var docsMapArr []map[string]interface{}

	for _, doc := range docs {
		docsMap := make(map[string]interface{})
		docsMap["username"] = doc.Properties["username"]
		docsMap["avatar"] = doc.Properties["avatar"]
		docsMap["total_tasks"] = doc.Properties["total_tasks"]

		docsMapArr = append(docsMapArr, docsMap)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"top_ten": docsMapArr,
	})
	return nil
}

func readFromMainDB(ctx *gin.Context) {
	topTenQ, topTenErr := models.PostgresDB.Query(context.Background(), query.GetTopTen)
	if topTenErr != nil {
		var pgErr *pgconn.PgError
		if errors.As(topTenErr, &pgErr) {

			ctx.JSON(http.StatusUnprocessableEntity, gin.H{
				"message": pgErr.Message,
			})
			return
		}
	}

	var topTenMap []map[string]interface{}

	for topTenQ.Next() {
		var ttu TopTenUsers

		user := make(map[string]interface{})

		scanErr := topTenQ.Scan(&ttu.UserName, &ttu.CompletedTasks)
		if scanErr != nil {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{
				"message": scanErr.Error(),
			})
			return
		}

		user["username"] = ttu.UserName
		user["completed_task"] = ttu.CompletedTasks
		topTenMap = append(topTenMap, user)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"top_ten": topTenMap,
	})
}
