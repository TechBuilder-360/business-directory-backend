package middlewares

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/TechBuilder-360/business-directory-backend/configs"
	"github.com/TechBuilder-360/business-directory-backend/models"
	"github.com/TechBuilder-360/business-directory-backend/repository"
	"github.com/TechBuilder-360/business-directory-backend/services"
	"github.com/TechBuilder-360/business-directory-backend/utility"
	log "github.com/Toflex/oris_log"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type Middleware struct {
	Repo   repository.Repository
	Logger log.Logger
	Config *configs.Config
}

func (m *Middleware) ClientValidation() gin.HandlerFunc {
	m.Logger.Info("Client Validation middleware is active")
	response := utility.NewResponse()
	return func(c *gin.Context) {

		if strings.HasPrefix(c.Request.RequestURI, "/api/v1") {
			if c.Request.Method != http.MethodGet && c.Request.Method != http.MethodDelete {
				client := &models.Client{}
				client.ClientID = c.Request.Header.Get("CID")
				clientSecret := c.Request.Header.Get("CS")
				client, err := m.Repo.GetClientByID(client.ClientID)
				if err != nil {
					m.Logger.Error("Client not found. %s", err.Error())
					c.JSON(http.StatusOK, response.Error(utility.CLIENTERROR, utility.GetCodeMsg(utility.CLIENTERROR)))
					c.Abort()
					return
				}
				body, req := utility.ExtractRequestBody(c.Request.Body)
				c.Request.Body = req
				m.Logger.Debug(body)

				if !client.ValidateClient(clientSecret, body) {
					m.Logger.Error("Client validation failed!")
					c.JSON(http.StatusOK, response.Error(utility.CLIENTERROR, utility.GetCodeMsg(utility.CLIENTERROR)))
					c.Abort()
					return
				}
			}
		}

		c.Next()
	}
}

// AuthorizeJWT handles jwt validation
func AuthorizeJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		const BearerSchema = "Bearer"
		authHeader := c.GetHeader("Authorization")
		tokenString := authHeader[len(BearerSchema):]
		token, err := services.DefultJWTAuth().ValidateToken(tokenString)
		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			fmt.Println(claims)
		} else {
			fmt.Println(err)
			c.AbortWithStatus(http.StatusUnauthorized)
		}

	}
}

//SecurityMiddleware performs encryption of response and decrypt request
func (m *Middleware) SecurityMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		response := utility.NewResponse()
		ctx:=make(map[string]interface{})
		for k, v := range c.Request.Header {
			ctx[k] = v
		}
		log := m.Logger.NewContext(ctx)

		if strings.HasPrefix(c.Request.RequestURI, "/api/v1") {
			client := &models.Client{}
			client.ClientID = c.Request.Header.Get("CID")
			client, err := m.Repo.GetClientByID(client.ClientID)
			if err != nil {
				log.Error("Client not found. %s", err.Error())
				c.JSON(http.StatusOK, response.Error(utility.CLIENTERROR, utility.GetCodeMsg(utility.CLIENTERROR)))
				c.Abort()
				return
			}

			if c.Request.Method != http.MethodGet && c.Request.Method != http.MethodDelete {
				body, _ := ioutil.ReadAll(c.Request.Body)
				log.Debug("Encrypted request body %s", string(body))
				decrypt, err := utility.Decrypt(client.AESKey, string(body))
				if err != nil {
					log.Error("Request body could not be decrypted. %s", err.Error())
					resp,_ := json.Marshal(response.Error(utility.SECURITYDECRYPTERR, utility.GetCodeMsg(utility.SECURITYDECRYPTERR)))
					encrypt, err := utility.Encrypt(client.AESKey, string(resp))
					if err != nil {
						c.AbortWithStatusJSON(http.StatusOK, response.Error(utility.SECURITYDECRYPTERR, utility.GetCodeMsg(utility.SECURITYDECRYPTERR)))
						return
					}
					c.String(http.StatusOK, encrypt)
					c.Abort()
					return
				}
				c.Request.Body = ioutil.NopCloser(bytes.NewBuffer([]byte(decrypt)))
			}
			// Proceed to next middleware
			c.Next()
			// Encrypt response
			var obj interface{}
			json.NewDecoder(c.Request.Response.Body).Decode(&obj)
			log.Debug("%+v\n", obj)
			return
		}

		c.Next()
	}

}
