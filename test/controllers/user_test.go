package test

import (
	"bytes"
	"encoding/json"
	"github.com/TechBuilder-360/business-directory-backend/dto"
	"github.com/TechBuilder-360/business-directory-backend/models"
	"github.com/TechBuilder-360/business-directory-backend/test"
	"net/http"
	"net/http/httptest"
	"testing"
)


func Test_register_user_should_return_status_created(t *testing.T) {
	// Arrange

	teardown:= test.Setup(t)
	defer teardown()

	userId:="2425615151"
	activity := &models.Activity{By: userId, Message: "Registered"}
	requestData := &dto.Registration{}
	requestData.EmailAddress = "janedoe@mail.com"
	requestData.DisplayName = "Jane"
	requestData.PhoneNumber = "+2348190082722"
	requestData.LastName = "Doe"
	requestData.FirstName = "Jane"

	byteReq, _:= json.Marshal(requestData)

	//serv.EXPECT().GetAuthor().Return(person)
	test.Repo.EXPECT().DoesUserEmailExist(requestData.EmailAddress).Return(false,nil)
	test.Repo.EXPECT().RegisterUser(requestData).Return(userId,nil)
	test.Repo.EXPECT().AddActivity(activity).Return(nil)

	test.Router.HandleFunc("/user-registration", test.Ch.RegisterUser).Methods(http.MethodPost)

	request, _ := http.NewRequest(http.MethodPost, "/user-registration", bytes.NewBuffer(byteReq))

	// Act
	recorder:= httptest.NewRecorder()
	test.Router.ServeHTTP(recorder, request)

	// Asset
	if recorder.Code != http.StatusCreated {
		t.Errorf("Returned invalid status code. %d", recorder.Code)
	}
}
