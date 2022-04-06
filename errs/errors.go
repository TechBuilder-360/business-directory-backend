package errs

import (
	"github.com/TechBuilder-360/business-directory-backend/utility"
	"net/http"
)

type AppError struct {
	StatusCode    int    `json:",omitempty"`
	ResponseCode string `json:"response_code,omitempty"`
	ResponseMessage string `json:"message,omitempty"`
}

func (e AppError) AsMessage() string {
	return e.ResponseMessage
}

func NotFoundError(code string) *AppError {
	message := utility.GetCodeMsg(code)
	return &AppError{StatusCode: http.StatusNotFound, ResponseCode: code, ResponseMessage: message}
}

func UnexpectedError(code string) *AppError {
	message := utility.GetCodeMsg(code)
	return &AppError{StatusCode: http.StatusInternalServerError, ResponseCode: code, ResponseMessage: message}
}

func NewValidationError(code string) *AppError {
	message := utility.GetCodeMsg(code)
	return &AppError{StatusCode: http.StatusUnprocessableEntity, ResponseCode: code, ResponseMessage: message}
}

func CustomError(statusCode int, responseCode string, message *string) *AppError {
	resMsg:=utility.StringPtrToString(message)
	if resMsg == "" {
		resMsg = utility.GetCodeMsg(responseCode)
	}
	return &AppError{StatusCode: statusCode, ResponseCode: responseCode, ResponseMessage: resMsg}
}
