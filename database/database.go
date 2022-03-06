package database

import (
	"context"
	"github.com/TechBuilder-360/business-directory-backend/configs"
	log "github.com/Toflex/oris_log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type Database struct {
	Mongo *mongo.Database
	Logger log.Logger
	Config *configs.Config
}

func (d *Database) ConnectToMongo() {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(d.Config.MongoURI))
	if err != nil {
		d.Logger.Fatal("An error occurred when connection to mongo DB %s", err.Error())
	}

	d.Logger.Info("Connected to mongodb successfully")
	d.Mongo = client.Database(d.Config.MongoDBName)
	return
}