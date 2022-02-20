package controllers

import (
	"github.com/TechBuilder-360/business-directory-backend.git/dto"
	"github.com/TechBuilder-360/business-directory-backend.git/services"
	"github.com/TechBuilder-360/business-directory-backend.git/utility"
	"github.com/Toflex/oris_log/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Authentication interface {
	Login(ctx *gin.Context)
}

type NewAuthHandler struct {
	auth   services.AuthService
	jwtService services.JWTService
	logger  logger.Logger
}

func AuthHandler(auth services.AuthService, jwt services.JWTService, log logger.Logger) Authentication {
	return &NewAuthHandler{
		auth:      auth,
		jwtService: jwt,
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
