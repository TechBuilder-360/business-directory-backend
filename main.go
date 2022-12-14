package main

import (
	"fmt"
	"github.com/TechBuilder-360/business-directory-backend/docs"
	"github.com/TechBuilder-360/business-directory-backend/internal/common/utils"
	"github.com/TechBuilder-360/business-directory-backend/internal/configs"
	"github.com/TechBuilder-360/business-directory-backend/internal/database"
	"github.com/TechBuilder-360/business-directory-backend/internal/database/redis"
	"github.com/TechBuilder-360/business-directory-backend/internal/middlewares"
	"github.com/TechBuilder-360/business-directory-backend/routers"
	"github.com/TechBuilder-360/business-directory-backend/seeder"
	"github.com/gorilla/mux"
	logrus_papertrail "github.com/polds/logrus-papertrail-hook"
	log "github.com/sirupsen/logrus"
	_ "github.com/swaggo/files"
	"net/http"
	"os"
	"strconv"
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

func initLog() {

	log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.InfoLevel)

	// Log as JSON instead of the default ASCII formatter.
	port, err := strconv.Atoi(utils.AddToStr(configs.Instance.PaperTailPort))
	if err != nil {
		return
	}

	hook, err := logrus_papertrail.NewPapertrailHook(&logrus_papertrail.Hook{
		Host:    "logs.papertrailapp.com",
		Port:    port,
		Appname: utils.AddToStr(configs.Instance.PaperTailAppName),
	})

	hook.SetLevels([]log.Level{log.ErrorLevel, log.WarnLevel})

	if err == nil {
		log.AddHook(hook)
	}
}

func main() {

	configs.Load()

	initLog()

	// Generate swagger doc information
	documentation()

	// set up redis DB
	redis.NewClient(configs.Instance.RedisURL, configs.Instance.RedisPassword, configs.Instance.Namespace)
	dbConnection := database.ConnectDB()
	// migrate db models
	err := database.DBMigration(dbConnection)
	if err != nil {
		panic(fmt.Sprintf("Migration Failed: %s", err.Error()))
	}
	go seeder.Seed(dbConnection)

	// Setup cache
	middlewares.ResponseCache()

	// Set up the routes
	router := mux.NewRouter()
	routers.SetupRoutes(router)

	// Start the server
	log.Info("Server started on port ", configs.Instance.BASEURL)
	err = http.ListenAndServe(fmt.Sprintf("%s", configs.Instance.BASEURL), router)
	if err != nil {
		log.Error(err.Error())
		return
	}
}

func documentation() {
	// programmatically set swagger info
	docs.SwaggerInfo.Title = "Business directory API"
	docs.SwaggerInfo.Description = "This is the API for business directory api."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = fmt.Sprintf("%s", configs.Instance.BASEURL)
	docs.SwaggerInfo.BasePath = fmt.Sprintf("/%s/api/v1", configs.Instance.BASEURL)
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
}
