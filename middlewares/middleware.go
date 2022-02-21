package middlewares

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/TechBuilder-360/business-directory-backend.git/configs"
	"github.com/TechBuilder-360/business-directory-backend.git/models"
	"github.com/TechBuilder-360/business-directory-backend.git/repository"
	"github.com/TechBuilder-360/business-directory-backend.git/services"
	"github.com/TechBuilder-360/business-directory-backend.git/utility"
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
	response:=utility.NewResponse()
	return func(c *gin.Context) {

		if strings.HasPrefix( c.Request.RequestURI, "/api/v1" ) {
			if c.Request.Method != http.MethodGet && c.Request.Method != http.MethodDelete {
				client:=&models.Client{}
				client.ClientID = c.Request.Header.Get("CID")
				clientSecret:=c.Request.Header.Get("CS")
				client, err := m.Repo.GetClientByID(client.ClientID)
				if err!=nil{
					m.Logger.Error("Client not found. %s", err.Error())
					c.JSON(http.StatusOK, response.Error(utility.CLIENTERROR, utility.GetCodeMsg(utility.CLIENTERROR)))
					c.Abort()
					return
				}
				body, req := utility.ExtractRequestBody(c)
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

// AuthorizeJWT
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

