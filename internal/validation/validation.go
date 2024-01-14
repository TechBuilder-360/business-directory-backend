package validation

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	log "github.com/sirupsen/logrus"
	"reflect"
	"strings"
)

func ValidateStruct(requestData interface{}, logger *log.Entry) (string, bool) {
	validationRes := validator.New()
	if err := validationRes.Struct(requestData); err != nil {
		var validationErrors validator.ValidationErrors
		errors.As(err, &validationErrors)
		errMsgs := make([]string, 0)
		logger.Error("Validation failed on some fields : %+v", validationErrors)
		for _, err := range validationErrors {
			fieldName := err.Field()
			field, _ := reflect.TypeOf(&requestData).Elem().FieldByName(fieldName)
			fieldJSONName, _ := field.Tag.Lookup("json")

			errMsgs = append(errMsgs, fmt.Sprintf(
				"[%s]: '%v' | validation failed '%s'",
				fieldJSONName,
				err.Value(),
				err.Tag(),
			))
		}

		return strings.Join(errMsgs, "\n"), false
	}
	return "", true
}
