package utils

import "github.com/TechBuilder-360/business-directory-backend/internal/common/apiError"

// SuccessResponse ...
//type SuccessResponse struct {
//	Status  bool        `json:"status"`
//	Message string      `json:"message"`
//	Data    interface{} `json:"data,omitempty"`
//	Meta    interface{} `json:"meta,omitempty"`
//}

// ErrorResponse ...
//type ErrorResponse struct {
//	Status  bool   `json:"status"`
//	Message string `json:"message"`
//	Error   string `json:"apiError,omitempty"`
//}

type Response struct {
	ResponseCode   string      `json:"response_code"`
	ResponseStatus bool        `json:"response_status"`
	Message        string      `json:"message"`
	Data           interface{} `json:"data,omitempty"`
	Meta           interface{} `json:"meta,omitempty"`
	Error          string      `json:"error,omitempty"`
}

func Success(message string, data, meta interface{}) Response {
	return Response{
		ResponseCode:   "00",
		ResponseStatus: true,
		Message:        message,
		Data:           data,
		Meta:           meta,
	}
}

func Error(data apiError.AppError) Response {
	return Response{
		ResponseCode: "44",
		Message:      data.Message,
		Error:        data.Error,
	}
}

func ValidationError(data apiError.AppError) Response {
	return Response{
		ResponseCode: "40",
		Message:      data.Message,
		Error:        data.Error,
	}
}
