package main

import (
	"log"
	"os"

	"github.com/jamesstocktonj1/forlater-core/app"
)

func main() {
	serverConfig, err := app.LoadConfig(os.Getenv("APP_CONFIG"))
	if err != nil {
		log.Fatal(err)
	}

	s := app.NewServer(serverConfig)

	err = s.Run()
	if err != nil {
		log.Fatal(err)
	}
}
