package router

import (
	"github.com/gin-gonic/gin"
	authController "github.com/pooya-hajjar/todo/controllers/auth_controller"
	"github.com/pooya-hajjar/todo/middlewares/permissions"
)

func Init() *gin.Engine {
	rt := gin.Default()

	rt.GET("/", func(context *gin.Context) {
		context.String(200, "yo friend whats going on")
	})

	authGroup := rt.Group("auth")
	tasksGroup := rt.Group("tasks", permissions.CheckUserAccess)
	userGroup := rt.Group("user", permissions.CheckUserAccess)

	authGroup.POST("signup", authController.Signup)
	authGroup.POST("signin", authController.SignIn)

	tasksGroup.GET("/:user-id")
	tasksGroup.PATCH("/:task-id")

	userGroup.GET("/:user-id")
	userGroup.GET("/search/:user-id")
	userGroup.PATCH("/update/:user-id")

	return rt
}
