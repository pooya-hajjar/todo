package permissions

import (
	"net/http"

	"github.com/gin-gonic/gin"
	authController "github.com/pooya-hajjar/todo/controllers/auth_controller"
)

func CheckUserAccess(ctx *gin.Context) {
	token := authController.GetAuthCookie(ctx)
	claim, err := authController.VerifyToken(token)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "unauthorized",
		})
	}

	ctx.Set("user_id", claim.Id)

	ctx.Next()
}
