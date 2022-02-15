package apps

import (
	"github.com/TechBuilder-360/business-directory-backend.git/controllers"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"sync"
)

var once sync.Once

func (a *App) SetupRoutes() {
	once.Do(func() {

		controller := controllers.Controller{
			Service: a.Serv,
			Repository: a.Repo,
			Logger: a.Logger,
			Config: a.Config,
		}
		if a.Config.DEBUG {
			// use ginSwagger middlewares to serve the API docs
			a.Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		}

		a.Router.GET("/ping", controller.Ping)

		v1 := a.Router.Group("/api/v1")
		{
			v1.POST("/ping", controller.Ping)
		}

	})
	a.Logger.Info("Routes have been initialized")

}