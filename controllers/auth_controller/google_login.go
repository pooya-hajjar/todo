package authController

import (
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/pooya-hajjar/todo/config"
	"github.com/pooya-hajjar/todo/constants/query"
	"github.com/pooya-hajjar/todo/models"
	"golang.org/x/crypto/bcrypt"
)

type UserData struct {
	ID            string `json:"id"`
	Email         string `json:"email"`
	VerifiedEmail bool   `json:"verified_email"`
	Name          string `json:"name"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	Picture       string `json:"picture"`
	Locale        string `json:"locale"`
}

func GoogleLogin(ctx *gin.Context) {
	if !config.AppOath2Config.Ok {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "google api not implemented",
		})
		return
	}

	url := config.AppOath2Config.Client.AuthCodeURL("randomstate")

	ctx.Redirect(http.StatusSeeOther, url)
	ctx.JSON(http.StatusSeeOther, url)
}

func GoogleCallback(ctx *gin.Context) {
	if !config.AppOath2Config.Ok {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "google api not implemented",
		})
		return
	}

	state, ok := ctx.GetQuery("state")
	if !ok || state != "randomstate" {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "unauthorized", // states dont match
		})
		return
	}

	code, ok := ctx.GetQuery("code")
	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "unauthorized",
		})
		return
	}

	token, exchangeErr := config.AppOath2Config.Client.Exchange(context.Background(), code)
	if exchangeErr != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "unauthorized", // code token exchange failed
		})
		return
	}

	resp, getUserDataErr := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if getUserDataErr != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "unauthorized", // user data fetch failed
		})
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "server error",
		})
	}
	defer resp.Body.Close()
	handleUserData(ctx, body)

}

func handleUserData(ctx *gin.Context, body []byte) {
	var userData UserData

	// Unmarshal the JSON data into the struct
	if err := json.Unmarshal(body, &userData); err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "unauthorized",
		})
	}

	defaultPass := os.Getenv("DEFAULT_USER_PASSWORD")

	hashedPass, hashPassErr := bcrypt.GenerateFromPassword([]byte(defaultPass), bcrypt.DefaultCost)
	if hashPassErr != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "server error",
		})
		return
	}
	userName := generateUsername(userData.Email)

	if userId, exist := userExists(ctx, userName); exist {
		token, err := CreateToken(userId, Oauth2Issuer)
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
		return
	}

	var newUserid int
	insertUserError := models.PostgresDB.QueryRow(context.Background(), query.AddNewUser, userName, hashedPass, userData.Picture).Scan(&newUserid)
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

	token, err := CreateToken(newUserid, Oauth2Issuer)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "unauthorized",
		})
		return
	}

	SetAuthCookie(ctx, token)

	ctx.JSON(http.StatusCreated, gin.H{
		"message": token,
	})
	return
}

func generateUsername(email string) string {
	// Generate a username based on the user's email address
	// Example: user@example.com -> user
	// You can adjust this logic as needed
	// Ensure the generated username is unique within your system
	// You might need to sanitize or transform the email address
	// to meet your username requirements (e.g., lowercase, remove special characters)
	// For simplicity, let's use the part before the "@" symbol
	// and ensure it's lowercase
	return strings.Split(email, "@")[0]
}

func userExists(ctx *gin.Context, username string) (int, bool) {
	var userId int
	err := models.PostgresDB.QueryRow(context.Background(), query.CheckUserExist, username).Scan(&userId)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == "23505" {
				ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
					"message": "this username has already been selected",
				})
			}
			ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
				"message": pgErr.Message,
			})
		}
	}
	return userId, userId != 0
}
