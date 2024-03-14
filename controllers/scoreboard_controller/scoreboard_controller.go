package scoreboardController

import (
	"context"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/pooya-hajjar/todo/constants/query"
	"github.com/pooya-hajjar/todo/models"
)

type TopTenUsers struct {
	UserName       any `json:"username"`
	CompletedTasks any `json:"completed_tasks"`
}

func Top10(ctx *gin.Context) {
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
