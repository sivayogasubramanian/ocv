package main

import (
	"github.com/sivayogasubramanian/ocv/config"
	"github.com/sivayogasubramanian/ocv/routes"
	"log"
)

func main() {
	config.InitDB()
	r := routes.SetupRouter()

	if err := r.Run(); err != nil {
		log.Fatal("failed to run server")
	}
}
