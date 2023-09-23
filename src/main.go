package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sivayogasubramanian/ocv/src/config"
	"github.com/sivayogasubramanian/ocv/src/routes"
	"log"
	"os"
)

func main() {
	if os.Getenv("GIN_MODE") == "release" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}

	}

	db := config.InitDB()
	r := routes.SetupRouter(db)

	if err := r.Run(); err != nil {
		log.Fatal("failed to run server")
	}
}
