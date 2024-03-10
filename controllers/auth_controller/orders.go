package authController

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/pooya-hajjar/todo/constants/query"
	"github.com/pooya-hajjar/todo/models"
	apiErrors "github.com/pooya-hajjar/todo/utils/api_errors"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type LoginBody struct {
	UserName string `form:"username" json:"username" binding:"required,min=3"`
	Password string `form:"password" json:"password" binding:"required,min=8"`
}

type SignUpBody struct {
	LoginBody
	Avatar string `json:"avatar"`
}

func Signup(ctx *gin.Context) {
	var signUpBody SignUpBody

	validationErr := ctx.BindJSON(&signUpBody)
	apiErrors.HandleValidationError(ctx, validationErr)

	hashedPass, hashPassErr := bcrypt.GenerateFromPassword([]byte(signUpBody.Password), bcrypt.DefaultCost)
	if hashPassErr != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": hashPassErr,
		})
		return
	}

	err2 := models.PostgresDB.QueryRow(context.Background(), query.AddNewUser, signUpBody.UserName, hashedPass, signUpBody.Avatar).Scan()
	if err2 != nil {
		var pgErr *pgconn.PgError
		if errors.As(err2, &pgErr) {
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

func SignIn(ctx *gin.Context) {

}
