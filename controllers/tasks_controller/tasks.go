package tasksController

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgtype/zeronull"
	"github.com/pooya-hajjar/todo/constants/query"
	"github.com/pooya-hajjar/todo/models"
	apiErrors "github.com/pooya-hajjar/todo/utils/api_errors"
	responseHelper "github.com/pooya-hajjar/todo/utils/response_helper"
	"net/http"
)

type AddTaskBody struct {
	Title     string `json:"title" binding:"required,min=3,max=50"`
	Priority  int    `json:"priority,omitempty" binding:"omitempty,gte=1,lte=100"`
	StartTime string `json:"start_time,omitempty" binding:"omitempty,datetime=2006-01-02 15:04"`
	EndTime   string `json:"end_time,omitempty" binding:"omitempty,datetime=2006-01-02 15:04"`
}

func AddTask(ctx *gin.Context) {
	if userId, exist := ctx.Get("user_id"); exist {
		var addTaskBody AddTaskBody

		validationErr := ctx.ShouldBindJSON(&addTaskBody)
		if validationErr != nil {
			apiErrors.HandleValidationError(ctx, validationErr)
			return
		}

		_, insertTaskErr := models.PostgresDB.Exec(context.Background(), query.AddTask, userId, addTaskBody.Title, zeronull.Int4(addTaskBody.Priority), zeronull.Text(addTaskBody.StartTime), zeronull.Text(addTaskBody.EndTime))

		if insertTaskErr != nil {
			var pgErr *pgconn.PgError
			if errors.As(insertTaskErr, &pgErr) {

				if pgErr.Code == "23514" {
					ctx.JSON(http.StatusUnprocessableEntity, gin.H{
						"message": "start time and end time must both exist or must both not exist",
					})
					return
				}

				ctx.JSON(http.StatusUnprocessableEntity, gin.H{
					"message": pgErr.Message,
				})
				return
			}
		}

		ctx.JSON(http.StatusCreated, gin.H{
			"message": "success",
		})
		return
	}

	ctx.JSON(http.StatusInternalServerError, gin.H{
		"message": "server error",
	})
}

func GetTasks(ctx *gin.Context) {
	if userId, exist := ctx.Get("user_id"); exist {
		getTasksQ, getTasksErr := models.PostgresDB.Query(context.Background(), query.GetTasks, userId)
		if getTasksErr != nil {
			var pgErr *pgconn.PgError
			if errors.As(getTasksErr, &pgErr) {

				ctx.JSON(http.StatusUnprocessableEntity, gin.H{
					"message": pgErr.Message,
				})
				return
			}
		}

		var tasks []map[string]interface{}
		for getTasksQ.Next() {
			var task models.Tasks
			scanErr := getTasksQ.Scan(&task.Title, &task.CreatedAt, &task.UpdatedAt, &task.Priority, &task.Status, &task.StartTime, &task.EndTime)
			if scanErr != nil {
				ctx.JSON(http.StatusUnprocessableEntity, gin.H{
					"message": scanErr.Error(),
				})
				return
			}

			taskMap := make(map[string]interface{})
			taskMap["title"] = task.Title
			taskMap["created_at"] = responseHelper.NilOrValue(task.CreatedAt)
			taskMap["updated_at"] = responseHelper.NilOrValue(task.UpdatedAt)
			taskMap["priority"] = responseHelper.NilOrValue(task.Priority)
			taskMap["status"] = task.Status
			taskMap["start_time"] = responseHelper.NilOrValue(task.StartTime)
			taskMap["end_time"] = responseHelper.NilOrValue(task.EndTime)

			tasks = append(tasks, taskMap)
		}

		ctx.JSON(http.StatusOK, gin.H{
			"tasks": tasks,
		})
		return
	}

	ctx.JSON(http.StatusInternalServerError, gin.H{
		"message": "server error",
	})
}
