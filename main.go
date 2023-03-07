package main

import (
	"github.com/sivayogasubramanian/ocv/config"
	"github.com/sivayogasubramanian/ocv/models"
	"github.com/sivayogasubramanian/ocv/routes"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func main() {
	var err error

	config.DB, err = gorm.Open(mysql.Open(config.GetDatabaseUrl(config.BuildDBConfig())))
	if err != nil {
		log.Fatal("failed to connect database")
	}

	err = config.DB.AutoMigrate(&models.Student{}, &models.Teacher{})
	if err != nil {
		log.Fatal("failed to migrate database")
	}

	r := routes.SetupRouter()

	err = r.Run()
	if err != nil {
		log.Fatal("failed to run server")
	}
}
