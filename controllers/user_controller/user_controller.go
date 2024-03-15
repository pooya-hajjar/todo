package userController

import (
	"context"
	"errors"
	"fmt"
	cacheContoller "github.com/pooya-hajjar/todo/controllers/cache_contoller"
	apiErrors "github.com/pooya-hajjar/todo/utils/api_errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/pooya-hajjar/todo/constants/query"
	"github.com/pooya-hajjar/todo/models"
	responseHelper "github.com/pooya-hajjar/todo/utils/response_helper"
)

type updateUserBody struct {
	UserName string `json:"username" binding:"required,min=3,max=100"`
	Email    string `json:"email,omitempty" binding:"omitempty,email"`
	Avatar   string `json:"avatar,omitempty" binding:"omitempty,max=255"`
	Status   int    `json:"status" binding:"taskStatus"`
}

func UserInfo(ctx *gin.Context, userId int) {

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

func ThisUserInfo(ctx *gin.Context) {
	if userId, exist := ctx.Get("user_id"); exist {
		if userIdInt, ok := userId.(int); ok {
			UserInfo(ctx, userIdInt)
			return
		}
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "user id should be number",
		})
	}

	ctx.JSON(http.StatusInternalServerError, gin.H{
		"message": "server error",
	})
}

func SearchUser(ctx *gin.Context) {
	if userId := ctx.Param("id"); userId != "" {
		if userIdInt, err := strconv.Atoi(userId); err == nil {
			UserInfo(ctx, userIdInt)
			return
		}
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "user id should be number",
		})
		return
	}

	ctx.JSON(http.StatusInternalServerError, gin.H{
		"message": "server error",
	})
}

func UpdateUser(ctx *gin.Context) {
	if userId, exist := ctx.Get("user_id"); exist {
		if userIdInt, ok := userId.(int); ok {
			var updateBody updateUserBody
			validationErr := ctx.ShouldBindJSON(&updateBody)
			if validationErr != nil {
				apiErrors.HandleValidationError(ctx, validationErr)
				return
			}
			_, updateUserErr := models.PostgresDB.Exec(context.Background(), query.UpdateUser, userIdInt, updateBody.UserName, updateBody.Email, updateBody.Status, updateBody.Avatar)
			if updateUserErr != nil {
				var pgErr *pgconn.PgError
				if errors.As(updateUserErr, &pgErr) {
					ctx.JSON(http.StatusUnprocessableEntity, gin.H{
						"message": pgErr.Message,
					})
					return
				}

				ctx.JSON(http.StatusUnprocessableEntity, gin.H{
					"message": updateUserErr.Error(),
				})
				return
			}
			intUserId, ok := userId.(int)
			if !ok {
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"message": fmt.Sprintf("server error"),
				})
				return
			}

			cacheContoller.UpdateScoreBoardDocumentUsername(intUserId, updateBody.UserName)

			ctx.JSON(http.StatusOK, gin.H{
				"message": fmt.Sprintf("user updated"),
			})
			return
		}
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "user id should be number",
		})

	}
}
