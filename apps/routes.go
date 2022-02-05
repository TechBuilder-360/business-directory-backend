package apps

import (
	"github.com/TechBuilder-360/business-directory-backend.git/controllers"
	"github.com/TechBuilder-360/business-directory-backend.git/repository"
	"github.com/TechBuilder-360/business-directory-backend.git/services"
	"github.com/gin-gonic/gin"
	"sync"
)

var once sync.Once

func (a *App) SetupRouter() {
	once.Do(func() {

		repo:= repository.NewRepository(nil)
		service:= services.NewService(repo)
		controller := controllers.Controller{
			Service: service,
			Logger:  a.Logger,
		}

		a.Router = gin.Default()
		a.Router.GET("/ping", controller.Ping)
	})
	a.Logger.Info("Routes have been initialized")

}