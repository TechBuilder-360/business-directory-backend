package services

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

//go:generate mockgen -destination=../mocks/services/mockJWTService.go -package=services github.com/TechBuilder-360/business-directory-backend/services JWTService
type JWTService interface {
	GenerateToken(string) (string, error)
	ValidateToken(string) (*jwt.Token, error)
}

type authCustomClaims struct {
	UserId string `json:"user_id"`
	jwt.StandardClaims
}

type jwtServices struct {
	secretKey string
	issure    string
}

//DefaultJWTAuth ...
func DefaultJWTAuth(secret, issuer string) JWTService {
	return &jwtServices{
		secretKey: secret,
		issure: issuer,
	}
}

func (service *jwtServices) GenerateToken(userId string) (string, error) {
	claims := &authCustomClaims{
		userId,
		jwt.StandardClaims{
			//ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
			Issuer:    service.issure,
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//encoded string
	t, err := token.SignedString([]byte(service.secretKey))
	if err != nil {
		return "", err
	}
	return t, nil
}

func (service *jwtServices) ValidateToken(encodedToken string) (*jwt.Token, error) {
	return jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if _, isvalid := token.Method.(*jwt.SigningMethodHMAC); !isvalid {
			return nil, errors.New("invalid token")
		}
		return []byte(service.secretKey), nil
	})

}
