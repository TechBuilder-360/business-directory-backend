package main

import (
	"github.com/TechBuilder-360/business-directory-backend.git/apps"
	"github.com/TechBuilder-360/business-directory-backend.git/docs"
	"github.com/Toflex/oris_log/logger"
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

	// programmatically set swagger info
	docs.SwaggerInfo_swagger.Title = "Business directory API"
	docs.SwaggerInfo_swagger.Description = "This is the API for business directory api."
	docs.SwaggerInfo_swagger.Version = "1.0"
	docs.SwaggerInfo_swagger.Host = "localhost:8080"
	docs.SwaggerInfo_swagger.BasePath = "/api/v1"
	docs.SwaggerInfo_swagger.Schemes = []string{"http", "https"}

	APP:= apps.App{}

	APP.Logger = logger.New()

	// Set up the routes
	APP.SetupRouter()

	// Start the server
	APP.Router.Run(":8080")

}