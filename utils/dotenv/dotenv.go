package dotEnv

import (
	"log"

	"github.com/joho/godotenv"
)

func Load() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error Loading env file")
	}
}
