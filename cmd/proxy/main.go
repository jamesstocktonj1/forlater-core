package main

import (
	"log"

	"github.com/jamesstocktonj1/forlater-core/app"
)

func main() {
	s := app.NewServer()

	err := s.Run()
	if err != nil {
		log.Fatal(err)
	}
}
