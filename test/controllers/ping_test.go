package test

//func Test_ping_should_return_status_ok(t *testing.T) {
//	// Arrange
//
//	teardown:= test.Setup(t)
//	defer teardown()
//
//	//serv.EXPECT().GetAuthor().Return(person)
//
//	test.Router.HandleFunc("/ping", test.Ch.Ping).Methods(http.MethodGet)
//
//	request, _ := http.NewRequest(http.MethodGet, "/ping", nil)
//
//	// Act
//	recorder:= httptest.NewRecorder()
//	test.Router.ServeHTTP(recorder, request)
//
//	// Asset
//	if recorder.Code != http.StatusOK {
//		t.Errorf("Returned invalid status code. %d", recorder.Code)
//	}
//}
