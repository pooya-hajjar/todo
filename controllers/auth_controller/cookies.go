package authController

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func SetAuthCookie(ctx *gin.Context, token string) {
	expiration := time.Now().Add(7 * 24 * time.Hour) // Expires in 7 days
	cookie := http.Cookie{
		Name:     "auth",
		Value:    fmt.Sprintf("Bearer %s", token),
		Expires:  expiration,
		Path:     "/",
		Domain:   "localhost",
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteStrictMode, // Adjust as necessary
	}

	ctx.SetCookie(cookie.Name, cookie.Value, cookie.MaxAge, cookie.Path, cookie.Domain, cookie.Secure, cookie.HttpOnly)
}

func GetAuthCookie(ctx *gin.Context) string {
	token, err := ctx.Cookie("auth")
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "unauthorized",
		})
	}
	return token[len("Bearer "):]

}
