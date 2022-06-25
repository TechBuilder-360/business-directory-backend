package database

import (
	"fmt"
	"github.com/TechBuilder-360/business-directory-backend/configs"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	DB *gorm.DB
}

func ConnectDB() *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", configs.Instance.DbHost, configs.Instance.DbUser, configs.Instance.DbPass, configs.Instance.DbName, configs.Instance.DbPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Could not connect to DB. %s", err.Error())
	}
	return db
}
