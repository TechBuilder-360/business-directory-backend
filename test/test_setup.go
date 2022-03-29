package test

import (
	"github.com/TechBuilder-360/business-directory-backend/configs"
	"github.com/TechBuilder-360/business-directory-backend/controllers"
	"github.com/TechBuilder-360/business-directory-backend/mocks/repository"
	"github.com/TechBuilder-360/business-directory-backend/mocks/services"
	log "github.com/Toflex/oris_log"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"os"
	"testing"
)

var Router *mux.Router
var Ch controllers.Controller
var Serv *services.MockService
var AuthServ *services.MockJWTService
var Repo *repository.MockRepository



func Setup(t *testing.T) func() {
	os.Setenv("Secret", "thisisasupersecret")
	os.Setenv("MongoURI", "mongodb://localhost:27017")
	config := &configs.Config{}
	config = configs.Configuration()
	var logger = log.New(log.Configuration{Output: log.CONSOLE, Name: "business-directory"})

	config.MongoDBName = "Test_Directory"
	config.ClientCol = "Test_Clients"
	config.OrganCol = "Test_Organisations"
	config.BranchCol = "Test_Branch"
	config.BranchCol = "Test_Users"
	config.ActivityCol = "Test_Activities"
	config.TokenCol = "Test_Token"

	crtl:= gomock.NewController(t)

	Serv = services.NewMockService(crtl)
	AuthServ = services.NewMockJWTService(crtl)
	Repo = repository.NewMockRepository(crtl)

	Ch = controllers.DefaultController(Serv, AuthServ, logger, Repo, config)

	Router = mux.NewRouter()

	return func() {
		Router = nil
		// drop test database
		defer crtl.Finish()
	}
}
