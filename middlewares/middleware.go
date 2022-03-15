package middlewares

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/TechBuilder-360/business-directory-backend/configs"
	"github.com/TechBuilder-360/business-directory-backend/models"
	"github.com/TechBuilder-360/business-directory-backend/repository"
	"github.com/TechBuilder-360/business-directory-backend/utility"
	log "github.com/Toflex/oris_log"
	"github.com/dgrijalva/jwt-go"
)

type Middleware struct {
	Repo   repository.Repository
	Logger log.Logger
	Config *configs.Config
}

func (m *Middleware) ClientValidationMiddleware (next http.Handler) http.Handler {
	m.Logger.Info("ClientValidationMiddleware successfully registered")

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		m.Logger.Info("Client Validation middleware is active")
		response := utility.NewResponse()

		if strings.HasPrefix(r.RequestURI, "/api/v1") {
			if r.Method != http.MethodGet && r.Method != http.MethodDelete {
				client := &models.Client{}
				client.ClientID = r.Header.Get("CID")
				clientSecret := r.Header.Get("CS")
				client, err := m.Repo.GetClientByID(client.ClientID)
				if err != nil {
					m.Logger.Error("Client not found. %s", err.Error())
					json.NewEncoder(w).Encode(response.Error(utility.CLIENTERROR, utility.GetCodeMsg(utility.CLIENTERROR)))

					return
				}
				body, req := utility.ExtractRequestBody(r.Body)
				r.Body = req
				m.Logger.Debug(body)

				if !client.ValidateClient(clientSecret, body) {
					m.Logger.Error("Client validation failed!")
					json.NewEncoder(w).Encode(response.Error(utility.CLIENTERROR, utility.GetCodeMsg(utility.CLIENTERROR)))
					w.WriteHeader(http.StatusUnauthorized)
					return
				}
			}
		}

		next.ServeHTTP(w, r)
	})
}

// AuthorizeJWT handles jwt validation
//func AuthorizeJWT() gin.HandlerFunc {
//	return func(c *gin.Context) {
//		const BearerSchema = "Bearer"
//		authHeader := c.GetHeader("Authorization")
//		tokenString := authHeader[len(BearerSchema):]
//		token, err := services.DefultJWTAuth().ValidateToken(tokenString)
//		if token.Valid {
//			claims := token.Claims.(jwt.MapClaims)
//			fmt.Println(claims)
//		} else {
//			fmt.Println(err)
//			c.AbortWithStatus(http.StatusUnauthorized)
//		}
//	}
//}

// Response send encrypted response
type Response struct {
	Data string `json:"data"`
}

//SecurityMiddleware performs encryption of response and decrypt request
func (m *Middleware) SecurityMiddleware (next http.Handler) http.Handler {
	m.Logger.Info("SecurityMiddleware successfully registered")

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := utility.NewResponse()
		log := m.Logger.NewContext()
		log.AddContext("Header", r.Header)

		if strings.HasPrefix(r.RequestURI, "/api/v1") {
			client := &models.Client{}
			client.ClientID = r.Header.Get("CID")
			client, err := m.Repo.GetClientByID(client.ClientID)
			if err != nil {
				log.Error("Client not found. %s", err.Error())
				json.NewEncoder(w).Encode(response.Error(utility.CLIENTERROR, utility.GetCodeMsg(utility.CLIENTERROR)))
				w.WriteHeader(403)
				return
			}

			if r.Method != http.MethodGet && r.Method != http.MethodDelete {
				body, _ := ioutil.ReadAll(r.Body)
				log.Debug("Encrypted request body %s", string(body))
				decrypt, err := utility.Decrypt(client.AESKey, string(body))
				if err != nil {
					log.Error("Request body could not be decrypted. %s", err.Error())
					resp, _ := json.Marshal(response.Error(utility.SECURITYDECRYPTERR, utility.GetCodeMsg(utility.SECURITYDECRYPTERR)))
					encrypt, err := utility.Encrypt(client.AESKey, string(resp))
					if err != nil {
						json.NewEncoder(w).Encode(response.Error(utility.SECURITYDECRYPTERR, utility.GetCodeMsg(utility.SECURITYDECRYPTERR)))
						w.WriteHeader(200)
						return
					}
					json.NewEncoder(w).Encode(Response{Data: encrypt})
					w.WriteHeader(200)
					return
				}
				r.Body = ioutil.NopCloser(bytes.NewBuffer([]byte(decrypt)))
			}
			rr := httptest.NewRecorder()
			// Proceed to next middleware
			next.ServeHTTP(rr, r)
			// Encrypt response
			var obj interface{}
			json.NewDecoder(r.Response.Body).Decode(&obj)
			log.Debug("%+v\n", obj)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func extractClaims(tokenStr string) (jwt.MapClaims, bool) {
	hmacSecretString := "secureSecretText"
	hmacSecret := []byte(hmacSecretString)
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		// check token signing method etc
		return hmacSecret, nil
	})

	if err != nil {
		return nil, false
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, true
	} else {
		fmt.Printf("Invalid JWT Token")
		return nil, false
	}
}

//func (m *Middleware) AuthorizationMiddleware(role ...string) http.HandlerFunc {
//	return func(w http.ResponseWriter, r *http.Request) {
//		response := utility.NewResponse()
//		const BearerSchema = "Bearer"
//		authHeader := r.Header.Get("Authorization")
//		tokenString := authHeader[len(BearerSchema):]
//		m.Logger.Debug(tokenString)
//		token, err := services.DefultJWTAuth(m.Config.Secret).ValidateToken(tokenString)
//		if token.Valid {
//			claims, _ := extractClaims(tokenString)
//			m.Logger.Debug(claims)
//			if !utility.IsContain(role, claims.Role) {
//				m.Logger.Error(utility.UNAUTHORISE)
//				w.WriteHeader(http.StatusUnauthorized)
//				json.NewEncoder(w).Encode(response.Error(utility.UNAUTHORISE, utility.GetCodeMsg(utility.UNAUTHORISE)))
//				return
//			}
//			m.Logger.Debug(claims)
//		} else {
//			fmt.Println(err)
//			//c.AbortWithStatus(http.StatusUnauthorized)
//		}
//
//		//c.Next()
//	}
//}


// RoleWrapper used to restrict access to a controller by the user role
func (m *Middleware) RoleWrapper(controller http.HandlerFunc, roles ...string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := utility.NewResponse()
		var userRole []string
		if utility.UserHasRole(userRole, roles) {
			json.NewEncoder(w).Encode(response.Error(utility.AUTHERROR004, utility.GetCodeMsg(utility.AUTHERROR004)))
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		controller(w, r)
	})
}
