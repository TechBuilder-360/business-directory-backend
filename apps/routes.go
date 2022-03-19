package apps

import (
	"fmt"
	httpSwagger "github.com/swaggo/http-swagger"
	"sync"

	"github.com/TechBuilder-360/business-directory-backend/controllers"
	"github.com/TechBuilder-360/business-directory-backend/middlewares"
)

var once sync.Once

func (a *App) SetupRoutes() {
	once.Do(func() {

		m := middlewares.Middleware{}
		m.Repo = a.Repo
		m.Logger = a.Logger
		m.Config = a.Config

		//a.Router.Use(m.ClientValidation())
		//a.Router.Use(m.SecurityMiddleware())
		//--- End middlewares

		controller := controllers.DefaultController(a.Serv, a.Logger)
		//auth:= services.DefaultAuth(a.Repo)
		//jwt:=services.DefultJWTAuth(a.Config.Secret)
		//authHandler := controllers.AuthHandler(auth, jwt, a.Repo, a.Logger)

		baseURL := fmt.Sprintf("/%s", a.Config.URLPrefix)
		apiRouter := a.Router.PathPrefix(baseURL).Subrouter()
		apiRouter2 := a.Router.PathPrefix(baseURL).Subrouter()

		if a.Config.DEBUG {
			a.Router.PathPrefix(baseURL).Handler(httpSwagger.WrapHandler)
		}

		apiRouter.HandleFunc("/ping", controller.Ping)

		apiRouter2.Handle("/a/ping", m.RoleWrapper(controller.Ping, "OWNER"))

		// Groups routes that does not require authentication
		//preAuthentication := apiRouter ("/api/v1")
		//{
		//	preAuthentication.POST("/get-login-token", authHandler.SendLoginToken)
		//}
		//
		//// Group routes that requires authentication
		//authUrl := a.Router.Group("/api/v1")
		//{
		//	//authUrl.Use(m.AuthorizeJWT())
		//	authUrl.POST("/getLoginToken", authHandler.Login)
		//}

	})
	a.Logger.Info("Routes have been initialized")

}
