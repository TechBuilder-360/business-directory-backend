package middlewares

import (
	"github.com/TechBuilder-360/business-directory-backend.git/configs"
	"github.com/TechBuilder-360/business-directory-backend.git/models"
	"github.com/TechBuilder-360/business-directory-backend.git/repository"
	"github.com/TechBuilder-360/business-directory-backend.git/utility"
	"github.com/Toflex/oris_log/logger"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type Middleware struct {
	Repo   *repository.DefaultRepo
	Logger logger.Logger
	Config *configs.Config
}

func (m *Middleware) ClientValidation() gin.HandlerFunc {
	m.Logger.Info("Client Validation middleware is active")
	response:=utility.NewResponse()
	return func(c *gin.Context) {

		if strings.HasPrefix( c.Request.RequestURI, "/api/v1" ) {
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
			m.Logger.Warning(body)

			if !client.ValidateClient(clientSecret, body) {
				m.Logger.Error("Client validation failed!")
				c.JSON(http.StatusOK, response.Error(utility.CLIENTERROR, utility.GetCodeMsg(utility.CLIENTERROR)))
				c.Abort()
				return
			}
		}

		c.Next()
	}
}
