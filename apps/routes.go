package apps

import (
	"sync"

	"github.com/TechBuilder-360/business-directory-backend/controllers"
	"github.com/TechBuilder-360/business-directory-backend/middlewares"
	"github.com/TechBuilder-360/business-directory-backend/services"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

var once sync.Once

func (a *App) SetupRoutes() {
	once.Do(func() {
		m := middlewares.Middleware{}
		m.Repo = a.Repo
		m.Logger = a.Logger
		m.Config = a.Config

		a.Router.SetTrustedProxies(a.Config.TrustedProxies)
		controller := controllers.DefaultController(a.Serv, a.Logger)
		auth := services.DefaultAuth(a.Repo)
		jwt := services.DefultJWTAuth()
		authHandler := controllers.AuthHandler(auth, jwt, a.Logger)
		a.Router.Use(cors.Default())
		a.Router.Use(gin.Recovery())
		//a.Router.Use(m.ClientValidation())
		//a.Router.Use(m.SecurityMiddleware())
		if a.Config.DEBUG {
			// use ginSwagger middlewares to serve the API docs
			a.Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		}

		a.Router.GET("/ping", controller.Ping)

		v1 := a.Router.Group("/api/v1")
		{
			v1.POST("/ping", controller.Ping)
			v1.POST("/getLoginToken", authHandler.Login)
			v1.POST("/create", controllers.CreateBook)
			v1.GET("/get", m.AuthorizationMiddleware("admin", "staff"), controllers.GetBook)
		}

	})
	a.Logger.Info("Routes have been initialized")

}
