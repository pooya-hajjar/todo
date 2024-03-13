package main

import (
	"fmt"
	tasksController "github.com/pooya-hajjar/todo/controllers/tasks_controller"
	apiErrors "github.com/pooya-hajjar/todo/utils/api_errors"
	"log"
	"os"

	"github.com/pooya-hajjar/todo/models"
	"github.com/pooya-hajjar/todo/router"
	dotEnv "github.com/pooya-hajjar/todo/utils/dotenv"
)

func init() {
	dotEnv.Load()
}

func main() {
	app := router.Init()

	cvs := []apiErrors.CustomValidator{
		apiErrors.CustomValidator{Tag: "taskStatus", Handler: tasksController.StatusValidator},
	}

	apiErrors.RegisterCustomValidator(cvs...)

	models.ConnectToDatabases()

	appPort := fmt.Sprintf(":%s", os.Getenv("APP_PORT"))
	runErr := app.Run(appPort)
	if runErr != nil {
		log.Fatalf("error on running app : %s !", runErr)
	}

}
