package main

import (
	"github.com/TechBuilder-360/business-directory-backend.git/apps"
	"github.com/Toflex/oris_log/logger"
)

func main()  {
	APP:= apps.App{}

	APP.Logger = logger.New()

	// Set up the routes
	APP.SetupRouter()

	// Start the server
	APP.Router.Run(":8080")

}