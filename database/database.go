package database

import (
	"fmt"
	log "github.com/Toflex/oris_log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	DB *gorm.DB
	Logger log.Logger
}

func (d *Database) ConnectToMongo() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", )
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	return
}