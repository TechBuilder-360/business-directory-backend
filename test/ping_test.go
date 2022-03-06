package test

import (
	"github.com/TechBuilder-360/business-directory-backend/controllers"
	"github.com/TechBuilder-360/business-directory-backend/mocks/service"
	"github.com/TechBuilder-360/business-directory-backend/services"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"net/http"
	"net/http/httptest"
	"testing"
)

var router *gin.Engine
var ch controllers.Controller
var serv *service.MockService


func setup(t *testing.T) func() {
	crtl:= gomock.NewController(t)

	serv = service.NewMockService(crtl)

	//ch = controllers.Controller{Service: serv}

	router = gin.Default()

	return func() {
		router = nil
		defer crtl.Finish()
	}
}

func Test_should_return_author(t *testing.T) {
	// Arrange



	person:=services.Person{
	Name:"Adegunwa Toluwalope",
	ID: 1200,
	}

	teardown:=setup(t)
	defer teardown()

	serv.EXPECT().GetAuthor().Return(person)

	router.GET("/ping", ch.Ping)

	request, _ := http.NewRequest(http.MethodGet, "/ping", nil)

	// Act
	recorder:= httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	// Asset
	if recorder.Code != http.StatusOK {
		t.Error("Returned invalid status code.")
	}
}
