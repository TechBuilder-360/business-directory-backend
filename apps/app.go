package apps

import (
	"github.com/TechBuilder-360/business-directory-backend/configs"
	"github.com/TechBuilder-360/business-directory-backend/repository"
	"github.com/TechBuilder-360/business-directory-backend/services"
	log "github.com/Toflex/oris_log"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

type App struct {
	Router *mux.Router
	Config *configs.Config
	Logger log.Logger
	Mongo  *mongo.Client
	Repo   repository.Repository
	Serv   services.Service

	//repo:= repository.NewRepository(a.Mongo, a.Config)
	//service:= services.NewService(repo)
}
