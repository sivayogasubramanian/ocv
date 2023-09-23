package config

import (
	"fmt"
	models2 "github.com/sivayogasubramanian/ocv/src/models"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"os"
	"strconv"
)

var DB *gorm.DB

type DBConfig struct {
	Host     string
	Port     int
	User     string
	DBName   string
	Password string
}

func BuildDBConfig() *DBConfig {
	port, _ := strconv.Atoi(os.Getenv("POSTGRES_PORT"))

	dbConfig := DBConfig{
		Host:     os.Getenv("POSTGRES_HOST"),
		Port:     port,
		User:     os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		DBName:   os.Getenv("POSTGRES_DATABASE_NAME"),
	}

	return &dbConfig
}

func GetDatabaseUrl() string {
	dbConfig := BuildDBConfig()

	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		strconv.Itoa(dbConfig.Port),
		dbConfig.DBName,
	)
}

func InitDB() {
	var err error

	dsn := GetDatabaseUrl()
	DB, err = gorm.Open(postgres.Open(dsn))
	if err != nil {
		log.Fatal("failed to connect database")
	}

	err = DB.AutoMigrate(&models2.Student{}, &models2.Teacher{})
	if err != nil {
		log.Fatal("failed to migrate database")
	}
}

func InitMemoryDB() {
	var err error

	DB, err = gorm.Open(sqlite.Open(":memory:"))
	if err != nil {
		log.Fatal("failed to connect database")
	}

	err = DB.AutoMigrate(&models2.Student{}, &models2.Teacher{})
	if err != nil {
		log.Fatal("failed to migrate database")
	}
}