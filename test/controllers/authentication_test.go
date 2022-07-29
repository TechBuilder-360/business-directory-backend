package test

//
//func Test_send_token_should_return_status_ok(t *testing.T) {
//	// Arrange
//
//	teardown := test.Setup(t)
//	defer teardown()
//
//	activity := &model.Activity{Message: "Requested for sign in token"}
//	requestData := &types.EmailRequest{}
//	requestData.EmailAddress = "janedoe@mail.com"
//
//	byteReq, _ := json.Marshal(requestData)
//
//	test.Repo.EXPECT().DoesUserEmailExist(requestData.EmailAddress).Return(true, nil)
//	test.Repo.EXPECT().CreateUserToken(requestData.EmailAddress).Return("123456", nil)
//	test.Repo.EXPECT().AddActivity(activity).Return(nil)
//
//	test.Router.HandleFunc("/request-login-token", test.Ch.AuthenticateEmail).Methods(http.MethodPost)
//	request, _ := http.NewRequest(http.MethodPost, "/request-login-token", bytes.NewBuffer(byteReq))
//
//	// Act
//	recorder := httptest.NewRecorder()
//	test.Router.ServeHTTP(recorder, request)
//
//	// Asset
//	if recorder.Code != http.StatusOK {
//		t.Errorf("Returned invalid status code. %d", recorder.Code)
//	}
//}
//
//func Test_login_should_return_status_ok(t *testing.T) {
//	// Arrange
//
//	teardown := test.Setup(t)
//	defer teardown()
//
//	requestData := &types.AuthRequest{}
//	requestData.Email = "janedoe@mail.com"
//	requestData.Token = "1234"
//
//	response := types.UserProfile{}
//	response.ID = uuid.New().String()
//	activity := &model.Activity{By: response.ID, Message: "Successful login"}
//
//	byteReq, _ := json.Marshal(requestData)
//
//	test.Repo.EXPECT().IsTokenValid(requestData).Return(true, nil)
//	test.Repo.EXPECT().GetUserInformation(requestData.Email).Return(response, nil)
//	test.AuthServ.EXPECT().GenerateToken(response.ID).Return(gomock.Any().String(), nil)
//	test.Repo.EXPECT().AddActivity(activity).Return(nil)
//
//	test.Router.HandleFunc("/login", test.Ch.Login).Methods(http.MethodPost)
//	request, _ := http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(byteReq))
//
//	// Act
//	recorder := httptest.NewRecorder()
//	test.Router.ServeHTTP(recorder, request)
//
//	// Asset
//	if recorder.Code != http.StatusOK {
//		t.Errorf("Returned invalid status code. %d", recorder.Code)
//	}
//}
