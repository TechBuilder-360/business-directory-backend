package apps

import (
	"github.com/TechBuilder-360/business-directory-backend.git/controllers"
	"github.com/TechBuilder-360/business-directory-backend.git/services"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"sync"
)

var once sync.Once

func (a *App) SetupRoutes() {
	once.Do(func() {

		controller := controllers.DefaultController(a.Serv, a.Logger)
		auth:= services.DefaultAuth(a.Repo)
		jwt:=services.DefultJWTAuth()
		authHandler := controllers.AuthHandler(auth, jwt, a.Logger)


		if a.Config.DEBUG {
			// use ginSwagger middlewares to serve the API docs
			a.Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		}

		a.Router.POST("/ping", controller.Ping)

		v1 := a.Router.Group("/api/v1")
		{
			v1.GET("/ping", controller.Ping)
			v1.POST("/getLoginToken", authHandler.Login)
		}

	})
	a.Logger.Info("Routes have been initialized")

}