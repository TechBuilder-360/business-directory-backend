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

		
		//--- End middlewares

		controller := controllers.DefaultController(a.Serv, a.Logger, a.Repo)
		
		baseURL := fmt.Sprintf("/%s", a.Config.URLPrefix)
		apiRouter := a.Router.PathPrefix(baseURL).Subrouter()
		apiRouter2 := a.Router.PathPrefix(baseURL).Subrouter()

		if a.Config.DEBUG {
			a.Router.PathPrefix(baseURL).Handler(httpSwagger.WrapHandler)
		}

		apiRouter.HandleFunc("/ping", controller.Ping)
		apiRouter2.Handle("/a/ping", m.RoleWrapper(controller.Ping, "OWNER")) 
		//organisation
		apiRouter2.HandleFunc("/organisation", controller.CreateOrganisation).Methods("POST")
		apiRouter2.HandleFunc("/get_organisation", controller.GetOrganisation).Methods("GET")

		//branch api
		apiRouter2.HandleFunc("/branch", controller.CreateBranch).Methods("POST")
		apiRouter2.HandleFunc("/branch/{organisationId}", controller.GetBranch).Methods("GET")

		//activation or deactivaion of organisation
		apiRouter2.HandleFunc("/de_activate_organisation/", controller.DeactivateOrganisation).Methods("POST")
		
		

	})
	a.Logger.Info("Routes have been initialized")

}
