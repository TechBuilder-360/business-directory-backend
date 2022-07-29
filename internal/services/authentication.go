package services

import (
	"errors"
	"github.com/TechBuilder-360/business-directory-backend/internal/common/types"
	"github.com/TechBuilder-360/business-directory-backend/internal/common/utils"
	"github.com/TechBuilder-360/business-directory-backend/internal/configs"
	"github.com/TechBuilder-360/business-directory-backend/internal/model"
	"github.com/TechBuilder-360/business-directory-backend/internal/repository"
	"github.com/dgrijalva/jwt-go"
	log "github.com/sirupsen/logrus"
	"time"
)

//go:generate mockgen -destination=../mocks/services/mockService.go -package=services github.com/TechBuilder-360/business-directory-backend/services UserService
type AuthService interface {
	RegisterUser(body *types.Registration, log *log.Entry) (*types.UserProfile, error)
	Login(body *types.AuthRequest) (*types.JWTResponse, error)
	AuthEmail(body *types.EmailRequest) error
	GenerateToken(userID string) (string, error)
	ValidateToken(encodedToken string) (*jwt.Token, error)
}

type DefaultAuthService struct {
	repo     repository.AuthRepository
	userRepo repository.UserRepository
	activity repository.ActivityRepository
}

func (d *DefaultAuthService) RegisterUser(body *types.Registration, log *log.Entry) (*types.UserProfile, error) {
	panic("implement me")
}

// Login
// Handles authentication logic
func (d *DefaultAuthService) Login(body *types.AuthRequest) (*types.JWTResponse, error) {
	response := &types.JWTResponse{}
	// Validate user token
	err := d.repo.IsTokenValid(body)
	if err != nil {
		log.Error("An Error occurred when validating login token. %s", err.Error())
		return nil, err
	}

	user := &model.User{}
	user.ID = body.UserId
	err = d.userRepo.Get(user)
	if err != nil {
		log.Error("An error occurred when fetching user profile. %s", err.Error())
		return nil, err
	}

	// Generate JWT for user
	token, err := d.GenerateToken(user.ID)
	if err != nil {
		log.Error("An error occurred when generating jwt token. %s", err.Error())
		return nil, errors.New("authentication failed")
	}

	profile := types.UserProfile{
		ID:            user.ID,
		FirstName:     user.FirstName,
		LastName:      user.LastName,
		DisplayName:   user.DisplayName,
		EmailAddress:  user.EmailAddress,
		PhoneNumber:   user.PhoneNumber,
		EmailVerified: user.EmailVerified,
	}

	response.Profile = profile
	response.AccessToken = token

	// Activity log
	activity := &model.Activity{By: response.Profile.ID, Message: "Successful login"}
	go func() {
		if err := d.activity.Create(activity); err != nil {
			log.Error("User activity failed to log")
		}
	}()

	return response, nil
}

func (d *DefaultAuthService) AuthEmail(body *types.EmailRequest) error {

	if !utils.ValidateEmail(body.EmailAddress) {
		return errors.New("invalid email address")
	}

	User := &model.User{}
	User.EmailAddress = body.EmailAddress

	// Check if email address exist
	err := d.userRepo.GetByEmail(User)
	if err != nil {
		log.Error("An Error occurred while checking if user email exist. %s", err.Error())
		return errors.New("email could not be retrieved")
	}
	if User.ID == "" {
		log.Info("Email address does not exist. '%s'", body.EmailAddress)
		return errors.New("email not found")
	}

	tk := &model.Token{}
	tk.UserID = User.ID
	tk.Token = utils.GenerateNumericToken(6)
	err = d.repo.CreateToken(tk)
	if err != nil {
		log.Error("Error occurred when saving sign-in token. %s", err.Error())
		return errors.New("token could not be generated")
	}

	// TODO: Send Token to user email

	// Activity log
	activity := &model.Activity{Message: "Requested for sign in token"}
	go func() {
		if err := d.activity.Create(activity); err != nil {
			log.Error("User activity failed to log")
		}
	}()

	return nil
}

type authCustomClaims struct {
	UserId string `json:"user_id"`
	jwt.StandardClaims
}

func (d *DefaultAuthService) GenerateToken(userId string) (string, error) {
	claims := &authCustomClaims{
		userId,
		jwt.StandardClaims{
			//ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
			Issuer:   configs.Instance.Issuer,
			IssuedAt: time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//encoded string
	t, err := token.SignedString([]byte(configs.Instance.Secret))
	if err != nil {
		log.Error(err.Error())
		return "", errors.New("token could not be generated")
	}
	return t, nil
}

func (d *DefaultAuthService) ValidateToken(encodedToken string) (*jwt.Token, error) {
	return jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if _, isvalid := token.Method.(*jwt.SigningMethodHMAC); !isvalid {
			return nil, errors.New("invalid token")
		}
		return []byte(configs.Instance.Secret), nil
	})
}

func NewAuthService() AuthService {
	return &DefaultAuthService{
		repo: repository.NewAuthRepository(),
	}
}
