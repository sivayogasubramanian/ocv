package config

import (
	"fmt"
	"github.com/sivayogasubramanian/ocv/models"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
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
	dbConfig := DBConfig{
		Host:     "localhost",
		Port:     3306,
		User:     "root",
		Password: "password",
		DBName:   "ocv",
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
