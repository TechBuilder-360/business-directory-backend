package utility

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func ValidateStruct(w http.ResponseWriter, requestData interface{}, logger *log.Entry) bool {
	validationRes := validator.New()
	if err := validationRes.Struct(requestData); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		logger.Error("Validation failed on some fields : %+v", validationErrors)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{
			Status:  false,
			Message: "Invalid request",
			Error:   validationErrors.Error(),
		})
		return true
	}
	return false
}
