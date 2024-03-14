package authController

import (
	"context"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/pooya-hajjar/todo/constants/query"
	"github.com/pooya-hajjar/todo/models"
	apiErrors "github.com/pooya-hajjar/todo/utils/api_errors"
	"golang.org/x/crypto/bcrypt"
)

type LoginBody struct {
	UserName string `json:"username"  binding:"required,min=3"`
	Password string `json:"password" binding:"required,min=8"`
}

type SignUpBody struct {
	LoginBody
	Avatar string `json:"avatar"`
}

type SignInBody struct {
	LoginBody
}

func Signup(ctx *gin.Context) {
	var signUpBody SignUpBody

	validationErr := ctx.ShouldBindJSON(&signUpBody)
	if validationErr != nil {
		apiErrors.HandleValidationError(ctx, validationErr)
		return
	}

	hashedPass, hashPassErr := bcrypt.GenerateFromPassword([]byte(signUpBody.Password), bcrypt.DefaultCost)
	if hashPassErr != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "server error",
		})
		return
	}

	var newUserid int
	insertUserError := models.PostgresDB.QueryRow(context.Background(), query.AddNewUser, signUpBody.UserName, hashedPass, signUpBody.Avatar).Scan(&newUserid)
	if insertUserError != nil {
		var pgErr *pgconn.PgError
		if errors.As(insertUserError, &pgErr) {
			if pgErr.Code == "23505" {
				ctx.JSON(http.StatusUnprocessableEntity, gin.H{
					"message": "this username has already been selected",
				})
				return
			}

			ctx.JSON(http.StatusUnprocessableEntity, gin.H{
				"message": pgErr.Message,
			})
			return
		}
	}

	token, err := CreateToken(newUserid, AppIssuer)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "unauthorized",
		})
		return
	}

	SetAuthCookie(ctx, token)

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "success",
	})
}

func SignIn(ctx *gin.Context) {

	var signInBody SignInBody

	validationErr := ctx.ShouldBindJSON(&signInBody)
	if validationErr != nil {
		apiErrors.HandleValidationError(ctx, validationErr)
		return
	}

	var dbId int
	var dbUsername string
	var dbPass string

	getUserError := models.PostgresDB.QueryRow(context.Background(), query.GetUser, signInBody.UserName).Scan(&dbId, &dbUsername, &dbPass)

	if getUserError != nil {

		var pgErr *pgconn.PgError
		if errors.As(getUserError, &pgErr) {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{
				"message": pgErr.Message,
			})
			return
		}
	}

	if dbUsername == "" {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "the name or password is incorrect",
		})
		return
	}

	comparePassErr := bcrypt.CompareHashAndPassword([]byte(dbPass), []byte(signInBody.Password))

	if comparePassErr != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "the name or password is incorrect",
		})
		return
	}

	token, err := CreateToken(dbId, AppIssuer)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "unauthorized",
		})
		return
	}

	SetAuthCookie(ctx, token)

	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
	})

}
