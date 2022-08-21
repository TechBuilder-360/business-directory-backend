package middlewares

import (
	"encoding/json"
	"github.com/TechBuilder-360/business-directory-backend/internal/common/constant"
	"github.com/TechBuilder-360/business-directory-backend/internal/common/utils"
	"github.com/TechBuilder-360/business-directory-backend/internal/services"
	"github.com/dgrijalva/jwt-go"
	log "github.com/sirupsen/logrus"
	"net/http"
)

// Response send encrypted response
type Response struct {
	Data string `json:"data"`
}

//func ClientValidationMiddleware(next http.Handler) http.Handler {
//	log.Info("ClientValidationMiddleware successfully registered")
//
//	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		log.Info("Client Validation middleware is active")
//		response := utils.NewResponse()
//
//		if strings.Contains(r.RequestURI, "/api/v1") {
//			if r.Method != http.MethodGet && r.Method != http.MethodDelete {
//				client := &model.Client{}
//				client.ClientID = r.Header.Get("CID")
//				clientSecret := r.Header.Get("CS")
//				client, err := m.Repo.GetClientByID(client.ClientID)
//				if err != nil {
//					log.Error("Client not found. %s", err.Error())
//					json.NewEncoder(w).Encode(response.Error(utils.CLIENTERROR))
//
//					return
//				}
//				body, req := utils.ExtractRequestBody(r.Body)
//				r.Body = req
//				log.Debug(body)
//
//				if !client.ValidateClient(clientSecret, body) {
//					log.Error("Client validation failed!")
//					json.NewEncoder(w).Encode(response.Error(utils.CLIENTERROR))
//					w.WriteHeader(http.StatusUnauthorized)
//					return
//				}
//			}
//		}
//
//		next.ServeHTTP(w, r)
//	})
//}

//AuthorizeUserJWT handles users jwt validation
func AuthorizeUserJWT(next http.Handler) http.Handler {
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
			})
			return
		}
		next.ServeHTTP(w, r)
	})
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
			})
			return
		}
		next.ServeHTTP(w, r)
	})
}

//SecurityMiddleware performs encryption of response and decrypt request
//func SecurityMiddleware(next http.Handler) http.Handler {
//	log.Info("SecurityMiddleware successfully registered")
//
//	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		response := utils.NewResponse()
//
//		if strings.Contains(r.RequestURI, "/api/v1") {
//			client := &model.Client{}
//			client.ClientID = r.Header.Get("CID")
//			client, err := m.Repo.GetClientByID(client.ClientID)
//			if err != nil {
//				log.Error("Client not found. %s", err.Error())
//				json.NewEncoder(w).Encode(response.Error(utils.CLIENTERROR))
//				w.WriteHeader(http.StatusOK)
//				return
//			}
//
//			if r.Method != http.MethodGet && r.Method != http.MethodDelete {
//				body, _ := ioutil.ReadAll(r.Body)
//				log.Debug("Encrypted request body %s", string(body))
//				decrypt, err := utils.Decrypt(client.AESKey, string(body))
//				if err != nil {
//					log.Error("Request body could not be decrypted. %s", err.Error())
//					resp, _ := json.Marshal(response.Error(utils.SECURITYDECRYPTERR))
//					encrypt, err := utils.Encrypt(client.AESKey, string(resp))
//					if err != nil {
//						json.NewEncoder(w).Encode(response.Error(utils.SECURITYDECRYPTERR))
//						w.WriteHeader(http.StatusOK)
//						return
//					}
//					json.NewEncoder(w).Encode(Response{Data: encrypt})
//					w.WriteHeader(http.StatusOK)
//					return
//				}
//				r.Body = ioutil.NopCloser(bytes.NewBuffer([]byte(decrypt)))
//			}
//			rec := httptest.NewRecorder()
//			// Proceed to next middleware
//			next.ServeHTTP(rec, r)
//			for k, v := range rec.Header() {
//				w.Header()[k] = v
//			}
//
//			switch w.Header().Get("Content-Encoding") {
//			case "gzip":
//				if reader, err := gzip.NewReader(rec.Result().Body); err != nil {
//					log.Error("%s: An error occurred while reading gzip > %s", r.RequestURI, err.Error())
//				} else {
//					json.NewDecoder(reader).Decode(&response)
//				}
//			default:
//				if err := json.Unmarshal(rec.Body.Bytes(), &response); err != nil {
//					log.Error("%s: An error unmarshalling response body > %s", r.RequestURI, err.Error())
//				}
//			}
//			// Response to JSON format
//			js, _ := json.Marshal(response)
//			jsonResponse := string(js)
//			// Encrypt response
//			if encoded, err := utils.Encrypt(client.AESKey, jsonResponse); err != nil {
//				log.Error("%s: An error occurred while encrypting response > %s", r.RequestURI, err.Error())
//				w.WriteHeader(http.StatusInternalServerError)
//				json.NewEncoder(w).Encode(response.Error(utils.SMMERROR))
//				return
//			} else {
//				resp, _ := json.Marshal(Response{Data: encoded})
//				w.WriteHeader(rec.Code)
//				var b bytes.Buffer
//				if w.Header().Get("Content-Encoding") == "gzip" {
//					gz := gzip.NewWriter(&b)
//					if _, err := gz.Write(resp); err != nil {
//						log.Error("%s: An error occurred while gzip response > %s", r.RequestURI, err.Error())
//						w.WriteHeader(http.StatusInternalServerError)
//						json.NewEncoder(w).Encode(response.Error(utils.SMMERROR))
//					}
//					if err := gz.Close(); err != nil {
//						log.Error("%s: An error occurred while closing gzip response > %s", r.RequestURI, err.Error())
//					}
//					w.Write(b.Bytes())
//					return
//				}
//				w.Write(resp)
//				return
//			}
//		}
//
//		next.ServeHTTP(w, r)
//
//	})
//}

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
//		constant BearerSchema = "Bearer"
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
