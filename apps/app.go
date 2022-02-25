package apps

import (
	"github.com/TechBuilder-360/business-directory-backend.git/configs"
	"github.com/TechBuilder-360/business-directory-backend.git/middlewares"
	"github.com/TechBuilder-360/business-directory-backend.git/repository"
	"github.com/TechBuilder-360/business-directory-backend.git/services"
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

// SetupMiddlewares sets up middlewares
func (a *App) SetupMiddlewares() {
	m := middlewares.Middleware{}
	m.Repo = a.Repo
	m.Logger = a.Logger
	m.Config = a.Config

	a.Router.Use(gin.Recovery())
	a.Router.Use(m.ClientValidation())
	a.Router.Use(m.SecurityMiddleware())
}
