package permissions

import (
	"net/http"

	"github.com/gin-gonic/gin"
	authController "github.com/pooya-hajjar/todo/controllers/auth_controller"
)

func CheckUserAccess(ctx *gin.Context) {
	token := authController.GetAuthCookie(ctx)
	_, err := authController.VerifyToken(token)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "unauthorized",
		})
		return
	}

	ctx.Next()
}
