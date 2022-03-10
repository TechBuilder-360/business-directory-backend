package controllers

import (
	"github.com/TechBuilder-360/business-directory-backend/dto"
	"github.com/TechBuilder-360/business-directory-backend/repository"
	"github.com/TechBuilder-360/business-directory-backend/services"
	"github.com/TechBuilder-360/business-directory-backend/utility"
	log "github.com/Toflex/oris_log"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Authentication interface {
	SendLoginToken(ctx *gin.Context)
	Login(ctx *gin.Context)
}

type NewAuthHandler struct {
	auth   services.AuthService
	jwtService services.JWTService
	repo repository.Repository
	logger  log.Logger
}

func AuthHandler(auth services.AuthService, jwt services.JWTService, repo repository.Repository, log log.Logger) Authentication {
	return &NewAuthHandler{
		auth:      auth,
		jwtService: jwt,
		repo: repo,
		logger:     log,
	}
}

func (l *NewAuthHandler) Login(ctx *gin.Context) {
	response := utility.NewResponse()
	var credential dto.AuthRequest
	err := ctx.ShouldBind(&credential)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, response.Error(utility.AUTHERROR004, utility.GetCodeMsg(utility.AUTHERROR004)))
		return
	}
	isUserAuthenticated := l.auth.LoginUser(credential.Email, credential.Token)
	if isUserAuthenticated {
		token:= l.jwtService.GenerateToken(credential.Email, true)
		ctx.JSON(http.StatusOK, response.Success( utility.SYSTEM001, utility.GetCodeMsg(utility.SYSTEM001), gin.H{
			"token": token,
		}))
		return
	}

	ctx.JSON(http.StatusUnauthorized, response.Error(utility.AUTHERROR004, utility.GetCodeMsg(utility.AUTHERROR004)))
	return
}

func (l *NewAuthHandler) SendLoginToken(c *gin.Context)  {
	response := utility.NewResponse()
	log:= l.logger.NewContext()
	log.SetLogID(c.Request.Header.Get("LogID"))

	requestData:=&dto.EmailRequest{}
	log.Info("Request data: %+v", requestData)

	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, response.ValidationError(utility.VALIDATIONERR, utility.GetCodeMsg(utility.VALIDATIONERR), err.Error()))
		return
	}



	c.JSON(http.StatusOK, response.PlainSuccess(utility.SYSTEM001, utility.GetCodeMsg(utility.SYSTEM001)))

}
