package apps

import (
	"github.com/TechBuilder-360/business-directory-backend/configs"
	"github.com/TechBuilder-360/business-directory-backend/repository"
	"github.com/TechBuilder-360/business-directory-backend/services"
	log "github.com/Toflex/oris_log"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type App struct {
	Router *gin.Engine
	Config *configs.Config
	Logger log.Logger
	Mongo  *mongo.Database
	Repo   repository.Repository
	Serv   services.Service

	//repo:= repository.NewRepository(a.Mongo, a.Config)
	//service:= services.NewService(repo)
}
