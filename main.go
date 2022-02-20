package main

import (
	"fmt"

	"github.com/TechBuilder-360/business-directory-backend.git/apps"
	"github.com/TechBuilder-360/business-directory-backend.git/configs"
	"github.com/TechBuilder-360/business-directory-backend.git/database"
	"github.com/TechBuilder-360/business-directory-backend.git/docs"
	"github.com/TechBuilder-360/business-directory-backend.git/repository"
	"github.com/TechBuilder-360/business-directory-backend.git/services"
	log "github.com/Toflex/oris_log"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/swaggo/files"
	_ "github.com/swaggo/gin-swagger"
)

// @title           Business directory API
// @version         1.0
// @description     This is the API for business directory api.

// @contact.name   Techbuilder Support
// @contact.email  tech.builder.cirle@gmail.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth
func main()  {
	// APP config
	APP:= &apps.App{}
	APP.Config = configs.Configuration()
	APP.Logger = log.New()
	if !APP.Config.DEBUG {
		gin.SetMode(gin.ReleaseMode)
	}
	APP.Router = gin.New()

	// programmatically set swagger info
	docs.SwaggerInfo_swagger.Title = "Business directory API"
	docs.SwaggerInfo_swagger.Description = "This is the API for business directory api."
	docs.SwaggerInfo_swagger.Version = "1.0"
	docs.SwaggerInfo_swagger.Host = fmt.Sprintf("%s:%s", APP.Config.Host,APP.Config.Port)
	docs.SwaggerInfo_swagger.BasePath = "/api/v1"
	docs.SwaggerInfo_swagger.Schemes = []string{"http", "https"}

	// Database config
	Database:=&database.Database{}
	Database.Config = APP.Config
	Database.Logger = APP.Logger
	Database.ConnectToMongo()

	APP.Mongo = Database.Mongo
	APP.Repo=repository.NewRepository(APP.Mongo, APP.Config)
	APP.Serv=services.NewService(APP.Repo, APP.Logger)

	// middlewares ...
	APP.Router.SetTrustedProxies(APP.Config.TrustedProxies)
	config := cors.DefaultConfig()
	config.AllowOrigins = APP.Config.AllowedOrigin
	APP.Router.Use(cors.New(config))
	APP.SetupMiddlewares()

	// Set up the routes
	APP.SetupRoutes()

	// Start the server
	APP.Logger.Info("Server started on port %s", APP.Config.Port)
	APP.Router.Run(fmt.Sprintf(":%s",APP.Config.Port))

}
