package router

import (
	"github.com/gin-gonic/gin"
	authController "github.com/pooya-hajjar/todo/controllers/auth_controller"
	scoreboardController "github.com/pooya-hajjar/todo/controllers/scoreboard_controller"
	tasksController "github.com/pooya-hajjar/todo/controllers/tasks_controller"
	"github.com/pooya-hajjar/todo/middlewares/permissions"
)

func Init() *gin.Engine {
	rt := gin.Default()

	rt.GET("/", func(context *gin.Context) {
		context.String(200, "I Have No Enemy :)")
	})

	authGroup := rt.Group("auth")
	tasksGroup := rt.Group("tasks", permissions.CheckUserAccess)
	userGroup := rt.Group("user", permissions.CheckUserAccess)

	authGroup.POST("signup", authController.Signup)
	authGroup.POST("signin", authController.SignIn)

	tasksGroup.POST("/add", tasksController.AddTask)
	tasksGroup.GET("/get", tasksController.GetTasks)
	tasksGroup.GET("/get/:id", tasksController.GetTask)
	tasksGroup.DELETE("/delete/:id", tasksController.DeleteTask)
	tasksGroup.PUT("/rename/:id", tasksController.RenameTask)
	tasksGroup.PUT("/update/:id", tasksController.UpdateTask)

	userGroup.GET("/:user-id")
	userGroup.GET("/search/:user-id")
	userGroup.PATCH("/update/:user-id")

	rt.GET("/top_ten", scoreboardController.Top10)

	return rt
}
