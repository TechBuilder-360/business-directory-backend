package apiError

type AppError struct {
	Error   string `json:"apiError"`
	Message string `json:"message"`
}
