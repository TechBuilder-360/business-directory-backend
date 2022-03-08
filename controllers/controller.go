package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/TechBuilder-360/business-directory-backend/services"
	logger "github.com/Toflex/oris_log"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type Controller interface {
	Ping(ctx *gin.Context)
}
type customClaims struct {
	Username string `json:"username"`
	Role     string ` json:"role"`
	jwt.StandardClaims
}
type NewController struct {
	Service services.Service
	Logger  logger.Logger
}
type CreateBookInput struct {
	Title  string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
}
type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	Token  string `json: "token"`
}

func DefaultController(serv services.Service, log logger.Logger) Controller {
	return &NewController{
		Service: serv,
		Logger:  log,
	}
}

func (c *NewController) Ping(ct *gin.Context) {
	var body interface{}
	err := json.NewDecoder(ct.Request.Body).Decode(&body)
	if err != nil {
		c.Logger.Error(err.Error())
	}
	c.Logger.Info("%+v", body)
	ct.JSON(http.StatusOK, "Pong ...")
}

func CreateBook(c *gin.Context) {

	claims := customClaims{
		Username: "shola",
		Role:     "admin",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: 15000,
			Issuer:    "createbook",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, _ := token.SignedString([]byte("secureSecretText"))
	// Validate input
	var input CreateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create book
	book := Book{Title: input.Title, Author: input.Author, Token: signedToken}
	c.JSON(http.StatusOK, gin.H{"data": book})
}

func GetBook(c *gin.Context) {

	c.JSON(http.StatusOK, "get ...")
}
