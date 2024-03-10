package main

import (
	"context"
	"github.com/pooya-hajjar/todo/models"
	"github.com/pooya-hajjar/todo/router"
	dotEnv "github.com/pooya-hajjar/todo/utils/dotenv"
	"log"
)

var CTX = context.Background()

func init() {
	dotEnv.Load()
}

func main() {
	app := router.Init()

	models.ConnectToDatabases()

	//col, err := pgx.CollectRows(q, pgx.RowToStructByName[PGClass])

	//trustProxiesErr := app.SetTrustedProxies(nil)
	//if trustProxiesErr != nil {
	//	log.Fatalf("error on running app : %s !", trustProxiesErr)
	//}

	runErr := app.Run(":3000")
	if runErr != nil {
		log.Fatalf("error on running app : %s !", runErr)
	}

}
