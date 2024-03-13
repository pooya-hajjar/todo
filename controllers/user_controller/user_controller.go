package userController

import (
	"context"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/pooya-hajjar/todo/constants/query"
	"github.com/pooya-hajjar/todo/models"
	responseHelper "github.com/pooya-hajjar/todo/utils/response_helper"
)

func UserInfo(ctx *gin.Context) {
	if userId, exist := ctx.Get("user_id"); exist {

		getUserInfoQ := models.PostgresDB.QueryRow(context.Background(), query.UserInfo, userId)

		user := struct {
			TotalTasks int `json:"total_tasks"`
			TodayTasks int `json:"today_tasks"`
			models.Users
		}{}

		scanErr := getUserInfoQ.Scan(&user.UserName, &user.Email, &user.Status, &user.Avatar, &user.TotalTasks, &user.TodayTasks)
		if scanErr != nil {
			var pgErr *pgconn.PgError
			if errors.As(scanErr, &pgErr) {
				ctx.JSON(http.StatusUnprocessableEntity, gin.H{
					"message": pgErr.Message,
				})
				return
			}

			ctx.JSON(http.StatusUnprocessableEntity, gin.H{
				"message": scanErr.Error(),
			})
			return
		}
		userInfoMap := make(map[string]interface{})

		userInfoMap["username"] = user.UserName
		userInfoMap["status"] = user.Status
		userInfoMap["total_tasks"] = user.TotalTasks
		userInfoMap["today_tasks"] = user.TodayTasks
		userInfoMap["email"] = responseHelper.NilOrValue(user.Email)
		userInfoMap["avatar"] = responseHelper.NilOrValue(user.Avatar)

		ctx.JSON(http.StatusOK, userInfoMap)
		return
	}

	ctx.JSON(http.StatusInternalServerError, gin.H{
		"message": "server error",
	})
}
