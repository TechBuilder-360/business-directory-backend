package apps

import (
	"fmt"
	"sync"

	httpSwagger "github.com/swaggo/http-swagger"

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

		if !a.Config.DEBUG{
			//a.Router.Use(m.ClientValidationMiddleware)
			a.Router.Use(m.SecurityMiddleware)
		}
		//--- End middlewares

		controller := controllers.DefaultController(a.Serv, a.Logger, a.Repo)
		//auth:= services.DefaultAuth(a.Repo)
		//jwt:=services.DefultJWTAuth(a.Config.Secret)
		//authHandler := controllers.AuthHandler(auth, jwt, a.Repo, a.Logger)

		baseURL := fmt.Sprintf("/%s/api/v1", a.Config.URLPrefix)
		apiRouter := a.Router.PathPrefix(baseURL).Subrouter()

		if a.Config.DEBUG {
			a.Router.PathPrefix(baseURL).Handler(httpSwagger.WrapHandler)
		}

		apiRouter.HandleFunc("/ping", controller.Ping)
		apiRouter.Handle("/a/ping", m.RoleWrapper(controller.Ping, "OWNER")) 
		//organisation
		apiRouter.HandleFunc("/organisation", controller.CreateOrganisation).Methods("POST")
		apiRouter.HandleFunc("/get_organisation", controller.GetOrganisation).Methods("GET")

		//branch api
		apiRouter.HandleFunc("/branch", controller.CreateBranch).Methods("POST")
		apiRouter.HandleFunc("/branch/{organisationId}", controller.GetBranch).Methods("GET")

		//activation or deactivaion of organisation
		apiRouter.HandleFunc("/de_activate_organisation/", controller.DeactivateOrganisation).Methods("POST")
		
		

	})
	a.Logger.Info("Routes have been initialized")

}
