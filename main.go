package main

import (
	"fmt"
	"github.com/TechBuilder-360/business-directory-backend/apps"
	"github.com/TechBuilder-360/business-directory-backend/configs"
	"github.com/TechBuilder-360/business-directory-backend/database"
	"github.com/TechBuilder-360/business-directory-backend/docs"
	"github.com/TechBuilder-360/business-directory-backend/repository"
	"github.com/TechBuilder-360/business-directory-backend/services"
	log "github.com/Toflex/oris_log"
	"github.com/gorilla/mux"
	_ "github.com/swaggo/files"
	"net/http"
)

// @title           Business directory API
// @version         1.0
// @description     This is the API for business directory api..

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8000
// @BasePath  /directory/api/v1

// @Security ApiKeyAuth
// @securityDefinitions.basic  ApiKeyAuth


func main() {
	// APP config
	APP := &apps.App{}
	APP.Config = configs.Configuration()
	APP.Logger = log.New(log.Configuration{Output: log.CONSOLE, Name: "Business_Directory"})

	// Server
	APP.Router = mux.NewRouter()

	// programmatically set swagger info
	docs.SwaggerInfo.Title = "Business directory API"
	docs.SwaggerInfo.Description = "This is the API for business directory api."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = fmt.Sprintf("%s:%s", APP.Config.Host, APP.Config.Port)
	docs.SwaggerInfo.BasePath = fmt.Sprintf("/%s/api/v1", APP.Config.URLPrefix)
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	// Database config
	Database := &database.Database{}
	Database.Config = APP.Config
	Database.Logger = APP.Logger
	Database.ConnectToMongo()

	APP.Mongo = Database.Mongo
	APP.Repo = repository.NewRepository(APP.Mongo, APP.Config)
	APP.Serv = services.NewService(APP.Repo)

	// Set up the routes
	APP.SetupRoutes()

	// Start the server
	APP.Logger.Info("Server started on port %s", APP.Config.Port)
	http.ListenAndServe(fmt.Sprintf(":%s", APP.Config.Port), APP.Router)
}
