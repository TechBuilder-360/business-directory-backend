package utils

// SuccessResponse ...
type SuccessResponse struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Meta    interface{} `json:"meta,omitempty"`
}

// ErrorResponse ...
type ErrorResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Error   string `json:"error,omitempty"`
}

type Response struct {
	ResponseCode string      `json:"response_code"`
	Message      string      `json:"message"`
	Data         interface{} `json:"data,omitempty"`
	Meta         interface{} `json:"meta,omitempty"`
	Error        string      `json:"error,omitempty"`
}

func Success(message string, data, meta interface{}) Response {
	return Response{
		ResponseCode: "00",
		Message:      message,
		Data:         data,
		Meta:         meta,
	}
}

func Error(data AppError) Response {
	return Response{
		ResponseCode: "40",
		Message:      data.Message,
		Error:        data.Error,
	}
}
