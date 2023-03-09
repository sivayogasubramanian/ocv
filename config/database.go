package config

import (
	"fmt"
	"github.com/sivayogasubramanian/ocv/models"
	"gorm.io/driver/mysql"
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
	port, _ := strconv.Atoi(os.Getenv("MYSQL_PORT"))

	dbConfig := DBConfig{
		Host:     os.Getenv("MYSQL_HOST"),
		Port:     port,
		User:     os.Getenv("MYSQL_USER"),
		Password: os.Getenv("MYSQL_PASSWORD"),
		DBName:   os.Getenv("MYSQL_DATABASE_NAME"),
	}

	return &dbConfig
}

func GetDatabaseUrl(dbConfig *DBConfig) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
	)
}

func InitDB() {
	var err error

	DB, err = gorm.Open(mysql.Open(GetDatabaseUrl(BuildDBConfig())))
	if err != nil {
		log.Fatal("failed to connect database")
	}

	err = DB.AutoMigrate(&models.Student{}, &models.Teacher{})
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

	err = DB.AutoMigrate(&models.Student{}, &models.Teacher{})
	if err != nil {
		log.Fatal("failed to migrate database")
	}
}
