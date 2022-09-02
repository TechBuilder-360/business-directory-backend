package middlewares

import (
	"encoding/json"
	"github.com/TechBuilder-360/business-directory-backend/internal/common/utils"
	"github.com/TechBuilder-360/business-directory-backend/internal/services"
	"github.com/dgrijalva/jwt-go"
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

// Response send encrypted response
type Response struct {
	Data string `json:"data"`
}

type Middleware func(http.HandlerFunc) http.HandlerFunc

// Logging logs all requests with its path and the time it took to process
func Logging() Middleware {

	// Create a new Middleware
	return func(f http.HandlerFunc) http.HandlerFunc {

		// Define the http.HandlerFunc
		return func(w http.ResponseWriter, r *http.Request) {

			// Do middleware things
			start := time.Now()
			defer func() { log.Println(r.URL.Path, time.Since(start)) }()

			// Call the next middleware/handler in chain
			f(w, r)
		}
	}
}

// Method ensures that url can only be requested with a specific method, else returns a 400 Bad Request
func Method(m string) Middleware {

	// Create a new Middleware
	return func(f http.HandlerFunc) http.HandlerFunc {

		// Define the http.HandlerFunc
		return func(w http.ResponseWriter, r *http.Request) {

			// Do middleware things
			if r.Method != m {
				http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
				return
			}

			// Call the next middleware/handler in chain
			f(w, r)
		}
	}
}

// Chain applies middlewares to a http.HandlerFunc
func Chain(f http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for _, m := range middlewares {
		f = m(f)
	}
	return f
}

//AuthorizeOrganisationJWT handles organisation jwt validation
func AuthorizeOrganisationJWT(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		const BearerSchema = "bearer"
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
				Error:   nil,
			})
			return
		}
		next.ServeHTTP(w, r)
	})
}

//func extractClaims(tokenStr string) (jwt.MapClaims, bool) {
//	hmacSecretString := "secureSecretText"
//	hmacSecret := []byte(hmacSecretString)
//	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
//		// check token signing method etc
//		return hmacSecret, nil
//	})
//
//	if err != nil {
//		return nil, false
//	}
//
//	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
//		return claims, true
//	} else {
//		fmt.Printf("Invalid JWT Token")
//		return nil, false
//	}
//}

//func (m *Middleware) AuthorizationMiddleware(role ...string) http.HandlerFunc {
//	return func(w http.ResponseWriter, r *http.Request) {
//		response := utils.NewResponse()
//		consts BearerSchema = "Bearer"
//		authHeader := r.Header.Get("Authorization")
//		tokenString := authHeader[len(BearerSchema):]
//		log.Debug(tokenString)
//		token, err := services.DefultJWTAuth(m.Config.Secret).ValidateToken(tokenString)
//		if token.Valid {
//			claims, _ := extractClaims(tokenString)
//			log.Debug(claims)
//			if !utils.IsContain(role, claims.Role) {
//				log.Error(utils.UNAUTHORISE)
//				w.WriteHeader(http.StatusUnauthorized)
//				json.NewEncoder(w).Encode(response.Error(utils.UNAUTHORISE, utils.GetCodeMsg(utils.UNAUTHORISE)))
//				return
//			}
//			log.Debug(claims)
//		} else {
//			fmt.Println(err)
//			//c.AbortWithStatus(http.StatusUnauthorized)
//		}
//
//		//c.Next()
//	}
//}

// RoleWrapper used to restrict access to a controller by the user role
//func RoleWrapper(controller http.HandlerFunc, roles ...string) http.Handler {
//	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		response := utils.NewResponse()
//		var userRole []string
//		if utils.UserHasRole(userRole, roles) {
//			w.WriteHeader(http.StatusUnauthorized)
//			json.NewEncoder(w).Encode(response.Error(utils.AUTHERROR004))
//			return
//		}
//
//		controller(w, r)
//	})
//}

func Recovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		//defer func() {
		//	err := recover()
		//	if err != nil {
		//		log.Error(err) // May be log this error? Send to sentry?
		//
		//		jsonBody, _ := json.Marshal(utils.ErrorResponse{
		//			Status:  false,
		//			Message: "internal server error",
		//		})
		//
		//		w.Header().Set("Content-Type", "application/json")
		//		w.WriteHeader(http.StatusInternalServerError)
		//		w.Write(jsonBody)
		//	}
		//
		//}()

		next.ServeHTTP(w, r)

	})
}
