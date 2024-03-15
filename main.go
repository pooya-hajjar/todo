package main

import (
	"fmt"
	"log"
	"os"

	"github.com/pooya-hajjar/todo/config"

	tasksController "github.com/pooya-hajjar/todo/controllers/tasks_controller"
	apiErrors "github.com/pooya-hajjar/todo/utils/api_errors"

	"github.com/pooya-hajjar/todo/models"
	"github.com/pooya-hajjar/todo/router"
	dotEnv "github.com/pooya-hajjar/todo/utils/dotenv"
)

func init() {
	dotEnv.Load()
}

func main() {
	app := router.Init()

	config.GoogleConfig()

	cvs := []apiErrors.CustomValidator{
		{Tag: "taskStatus", Handler: tasksController.StatusValidator},
	}

	apiErrors.RegisterCustomValidator(cvs...)

	models.ConnectToDatabases()

	appPort := fmt.Sprintf(":%s", os.Getenv("APP_PORT"))
	runErr := app.Run(appPort)
	if runErr != nil {
		log.Fatalf("error on running app : %s !", runErr)
	}

}
