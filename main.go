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
// @description     This is the API for business directory api.

// @contact.name   Techbuilder Support
// @contact.email  tech.builder.cirle@gmail.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth
func main() {
	// APP config
	APP := &apps.App{}
	APP.Config = configs.Configuration()
	APP.Logger = log.New(log.Configuration{Output: log.CONSOLE, Name: "Business_Directory"})
	if !APP.Config.DEBUG {

	}

	// Server
	APP.Router = mux.NewRouter()

	// programmatically set swagger info
	docs.SwaggerInfo_swagger.Title = "Business directory API"
	docs.SwaggerInfo_swagger.Description = "This is the API for business directory api."
	docs.SwaggerInfo_swagger.Version = "1.0"
	docs.SwaggerInfo_swagger.Host = fmt.Sprintf("%s:%s", APP.Config.Host, APP.Config.Port)
	docs.SwaggerInfo_swagger.BasePath = "/api/v1"
	docs.SwaggerInfo_swagger.Schemes = []string{"http", "https"}

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
