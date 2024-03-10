package main

import (
	"context"
	"fmt"
	"github.com/pooya-hajjar/todo/models"
	"github.com/pooya-hajjar/todo/router"
	dotEnv "github.com/pooya-hajjar/todo/utils/dotenv"
	"log"
	"os"
)

var CTX = context.Background()

func init() {
	dotEnv.Load()
}

func main() {
	app := router.Init()

	models.ConnectToDatabases()

	appPort := fmt.Sprintf(":%s", os.Getenv("APP_PORT"))
	runErr := app.Run(appPort)
	if runErr != nil {
		log.Fatalf("error on running app : %s !", runErr)
	}

}
