package middlewares

import (
	"encoding/json"
	"github.com/TechBuilder-360/business-directory-backend/internal/common/constant"
	"github.com/TechBuilder-360/business-directory-backend/internal/common/utils"
	log "github.com/sirupsen/logrus"
	"net/http"
)

// Response send encrypted response
type Response struct {
	Data string `json:"data"`
}

type ContextKey string

const (
	AuthUserContextKey         ContextKey = "user"
	AuthOrganisationContextKey ContextKey = "organisation"
)

func Recovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		defer func() {
			err := recover()
			if err != nil {
				log.Error(err)

				jsonBody, _ := json.Marshal(utils.ErrorResponse{
					Status:  false,
					Message: constant.InternalServerError,
				})

				w.WriteHeader(http.StatusInternalServerError)
				w.Write(jsonBody)
			}

		}()

		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)

	})
}

// Adapter is an alias, so I don't have to type so much.
type Adapter func(http.Handler) http.Handler

// Adapt takes Handler functions and chains them to the main handler.
func Adapt(handler http.Handler, adapters ...Adapter) http.Handler {
	// The loop is reversed so the adapters/middleware gets executed in the same
	// order as provided in the array.
	for i := len(adapters); i > 0; i-- {
		handler = adapters[i-1](handler)
	}
	return handler
}
