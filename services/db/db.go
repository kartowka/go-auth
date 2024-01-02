package db

import (
	"fmt"
	"log"

	"github.com/antfley/go-auth/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBConfig struct {
	DBHost string
	DBUser string
	DBName string
	DBPort string
	DbPass string
}

func InitDB(config DBConfig) *gorm.DB {
	var err error
	dbURI := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jerusalem",
		config.DBHost, config.DBUser, config.DbPass, config.DBName, config.DBPort,
	)
	db, err := gorm.Open(postgres.Open(dbURI), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("!! Database Connection Loaded !!")
	return db
}
func DBMigrate(db *gorm.DB) {
	var err error
	for _, model := range models.RegisterModels() {
		err = db.Debug().AutoMigrate(model.Model)
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("!! Database Migration Succeed !!")
}
