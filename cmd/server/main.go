package main

import (
	"fmt"
	"github.com/TechBuilder-360/business-directory-backend/internal/configs"
	"github.com/TechBuilder-360/business-directory-backend/internal/database"
	"github.com/TechBuilder-360/business-directory-backend/internal/database/redis"
	"github.com/TechBuilder-360/business-directory-backend/internal/routers"
	"github.com/TechBuilder-360/business-directory-backend/pkg/apm/newrelic"
	"github.com/TechBuilder-360/business-directory-backend/pkg/apm/sentry"
	log "github.com/sirupsen/logrus"
	"os"
	"time"
)

func initLog() *log.Logger {
	l := log.New()
	l.SetFormatter(&log.JSONFormatter{})
	l.SetOutput(os.Stdout)
	l.SetLevel(log.InfoLevel)

	return l
}

func main() {
	configs.Load()

	l := initLog()
	newrelic.InitialiseNewRelic(l)

	sentryHook, err := sentry.InitializeSentry(l)
	if err == nil {
		defer sentryHook.Flush(5 * time.Second)
	}

	// set up redis DB
	redis.NewClient()
	dbConnection := database.ConnectDB()
	sqlDB, err := dbConnection.DB()
	if err != nil {
		l.Fatalf("database connection failed %v", err.Error())
	}

	defer sqlDB.Close()

	// Set up the routes
	router := routers.SetupRoutes()

	// Start the server
	l.Info("Server started on port ", configs.Instance.Port)
	err = router.Listen(fmt.Sprintf("%s:%s", configs.Instance.BASEURL, configs.Instance.Port))
	if err != nil {
		l.Error("apiError when starting server ::: %s", err.Error())
		return
	}
}
