package main

import (
	"fmt"
	"github.com/TechBuilder-360/business-directory-backend/docs"
	"github.com/TechBuilder-360/business-directory-backend/internal/configs"
	"github.com/TechBuilder-360/business-directory-backend/internal/database"
	"github.com/TechBuilder-360/business-directory-backend/internal/database/redis"
	"github.com/TechBuilder-360/business-directory-backend/routers"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	_ "github.com/swaggo/files"
	"net/http"
	"os"
	"github.com/TechBuilder-360/business-directory-backend/internal/common/utils"
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

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.WarnLevel)
}

func main() {


	res,err:= utils.SendMail("Activate your account","folayanshola@gmail.com","<h2 style='color:red;'> hello there</h2>","folayan adesola")
	if err!=nil{
		  log.Error("Error occurred when sending activation email. %s", err.Error())
		  return 
	}else{
		log.Println(res.Body)
  return 
	}
	 configs.Load()

	// // Generate swagger doc information
	 documentation()

	// // set up redis DB
	 redis.NewClient(configs.Instance.RedisURL, configs.Instance.RedisPassword, configs.Instance.Namespace)

	// // Set up the routes
	router := mux.NewRouter()
	routers.SetupRoutes(router)

	// // migrate db models
	 database.DBMigration()

	// // Start the server
	 log.Info("Server started on port %s", configs.Instance.Host)
	 err:= http.ListenAndServe(fmt.Sprintf("%s", configs.Instance.Host), router)
	 if err != nil {
	 	return
	 }
}

func documentation() {
	// programmatically set swagger info
	docs.SwaggerInfo_swagger.Title = "Business directory API"
	docs.SwaggerInfo_swagger.Description = "This is the API for business directory api."
	docs.SwaggerInfo_swagger.Version = "1.0"
	docs.SwaggerInfo_swagger.Host = fmt.Sprintf("%s", configs.Instance.Host)
	docs.SwaggerInfo_swagger.BasePath = fmt.Sprintf("/%s/api/v1", configs.Instance.URLPrefix)
	docs.SwaggerInfo_swagger.Schemes = []string{"http", "https"}
}
