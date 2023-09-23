package config

import (
	"fmt"
	"github.com/sivayogasubramanian/ocv/src/models"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"os"
	"strconv"
)

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

func InitDB() (db *gorm.DB) {
	dsn := GetDatabaseUrl()
	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		log.Fatal("failed to connect database")
	}

	err = db.AutoMigrate(&models.Student{}, &models.Teacher{})
	if err != nil {
		log.Fatal("failed to migrate database")
	}

	return db
}

func InitMemoryDB() (db *gorm.DB) {
	db, err := gorm.Open(sqlite.Open(":memory:"))
	if err != nil {
		log.Fatal("failed to connect database")
	}

	err = db.AutoMigrate(&models.Student{}, &models.Teacher{})
	if err != nil {
		log.Fatal("failed to migrate database")
	}

	return db
}
