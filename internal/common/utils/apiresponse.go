package utils

// SuccessResponse ...
type SuccessResponse struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// ErrorResponse ...
type ErrorResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Error   string `json:"error,omitempty"`
}
