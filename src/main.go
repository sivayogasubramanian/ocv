package main

import (
	"github.com/joho/godotenv"
	"github.com/sivayogasubramanian/ocv/src/config"
	"github.com/sivayogasubramanian/ocv/src/routes"
	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config.InitDB()
	r := routes.SetupRouter()

	if err := r.Run(); err != nil {
		log.Fatal("failed to run server")
	}
}
