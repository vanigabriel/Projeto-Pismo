package main

import (
	"log"
	"os"

	"github.com/evalphobia/go-timber/timber"
	"github.com/joho/godotenv"

	"github.com/vanigabriel/Projeto-Pismo/api"
	"github.com/vanigabriel/Projeto-Pismo/usecases"
)

//PORT port to be used
const PORT = "9292"

func main() {
	// Load env variables
	handleError(godotenv.Load())

	conf := timber.Config{
		APIKey:         os.Getenv("TIMBER_API_KEY"),
		SourceID:       os.Getenv("TIMBER_SOURCE_ID"),
		CustomEndpoint: "https://logs.timber.io",
		Environment:    "production",
		MinimumLevel:   timber.LogLevelInfo,
		Sync:           true,
		Debug:          true,
	}

	log, err := timber.New(conf)
	handleError(err)

	s := usecases.NewService(usecases.InitStatements())

	r := api.SetupRouter(s, log)

	r.Run(":" + PORT)

}

func handleError(err error) {
	if err != nil {
		log.Println(err)
	}
}
