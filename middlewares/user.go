package middlewares

import (
	"encoding/json"
	"github.com/TechBuilder-360/business-directory-backend/internal/common/utils"
	"github.com/TechBuilder-360/business-directory-backend/internal/services"
	"github.com/dgrijalva/jwt-go"
	log "github.com/sirupsen/logrus"
	"net/http"
)

// AuthorizeUserJWT authorise user JWT
func AuthorizeUserJWT(m string) Middleware {

	// Create a new Middleware
	return func(f http.HandlerFunc) http.HandlerFunc {

		// Define the http.HandlerFunc
		return func(w http.ResponseWriter, r *http.Request) {

			const BearerSchema = "Bearer"
			authHeader := r.Header.Get("Authorization")
			tokenString := authHeader[len(BearerSchema):]
			token, err := services.NewAuthService().ValidateToken(tokenString)
			if token.Valid {
				claims := token.Claims.(jwt.MapClaims)
				log.Error("%+v", claims)
			} else {
				log.Error(err)
				w.WriteHeader(http.StatusUnauthorized)
				json.NewEncoder(w).Encode(utils.ErrorResponse{
					Status:  false,
					Message: "unauthorized",
				})
				return
			}

			// Call the next middleware/handler in chain
			f(w, r)
		}
	}
}
