package apps

import (
	"github.com/TechBuilder-360/business-directory-backend/controllers"
	"github.com/TechBuilder-360/business-directory-backend/middlewares"
	"github.com/TechBuilder-360/business-directory-backend/services"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"sync"
)

var once sync.Once

func (a *App) SetupRoutes() {
	once.Do(func() {

		// middlewares ...
		a.Router.SetTrustedProxies(a.Config.TrustedProxies)
		//middlewares.New(a.Router.)
		m := &middlewares.Middleware{}
		m.Repo = a.Repo
		m.Logger = a.Logger
		m.Config = a.Config

		a.Router.Use(cors.Default(), gin.Recovery(), m.TestMiddleware())
		//a.Router.Use(m.ClientValidation())
		//a.Router.Use(m.SecurityMiddleware())
		//--- End middlewares

		controller := controllers.DefaultController(a.Serv, a.Logger)
		auth:= services.DefaultAuth(a.Repo)
		jwt:=services.DefultJWTAuth()
		authHandler := controllers.AuthHandler(auth, jwt, a.Repo, a.Logger)


		if a.Config.DEBUG {
			// use ginSwagger middlewares to serve the API docs
			a.Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		}

		a.Router.GET("/ping", controller.Ping)

		// Groups routes that does not require authentication
		preAuthentication := a.Router.Group("/api/v1")
		{
			preAuthentication.POST("/get-login-token", authHandler.SendLoginToken)
		}

		// Group routes that requires authentication
		authUrl := a.Router.Group("/api/v1")
		{
			authUrl.Use(m.AuthorizeJWT())
			authUrl.POST("/getLoginToken", authHandler.Login)
		}

	})
	a.Logger.Info("Routes have been initialized")

}