package apps

import (
	"github.com/TechBuilder-360/business-directory-backend.git/controllers"
	"github.com/TechBuilder-360/business-directory-backend.git/repository"
	"github.com/TechBuilder-360/business-directory-backend.git/services"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"sync"
)

var once sync.Once

func (a *App) SetupRouter() {
	once.Do(func() {

		repo:= repository.NewRepository(nil)
		service:= services.NewService(repo)
		controller := controllers.Controller{
			Service: service,
		}

		a.Router = gin.Default()
		// use ginSwagger middleware to serve the API docs
		a.Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		a.Router.GET("/ping", controller.Ping)
	})
	a.Logger.Info("Routes have been initialized")

}