package services

import (
	"errors"
	"fmt"
	"github.com/TechBuilder-360/business-directory-backend/internal/common/constant"
	"github.com/TechBuilder-360/business-directory-backend/internal/common/types"
	"github.com/TechBuilder-360/business-directory-backend/internal/common/utils"
	"github.com/TechBuilder-360/business-directory-backend/internal/configs"
	"github.com/TechBuilder-360/business-directory-backend/internal/database/redis"
	"github.com/TechBuilder-360/business-directory-backend/internal/infrastructure/sendgrid"
	"github.com/TechBuilder-360/business-directory-backend/internal/model"
	"github.com/TechBuilder-360/business-directory-backend/internal/repository"
	"github.com/dgrijalva/jwt-go"
	log "github.com/sirupsen/logrus"
	"strings"
	"time"
)

//go:generate mockgen -destination=../mocks/services/mockService.go -package=services github.com/TechBuilder-360/business-directory-backend/services UserService
type AuthService interface {
	RegisterUser(body *types.Registration, log *log.Entry) error
	ActivateEmail(token string, uid string, log *log.Entry) error
	Login(body *types.AuthRequest) (*types.LoginResponse, error)
	GenerateJWT(userID string) (*string, error)
	ValidateToken(encodedToken string) (*jwt.Token, error)
	RequestToken(body *types.EmailRequest, logger *log.Entry) error
}

type DefaultAuthService struct {
	repo     repository.AuthRepository
	userRepo repository.UserRepository
	activity repository.ActivityRepository
	redis    *redis.Client
}

func NewAuthService() AuthService {
	return &DefaultAuthService{
		repo:     repository.NewAuthRepository(),
		userRepo: repository.NewUserRepository(),
		activity: repository.NewActivityRepository(),
		redis:    redis.RedisClient(),
	}
}

func (d *DefaultAuthService) ActivateEmail(token string, uid string, logger *log.Entry) error {

	user, err := d.userRepo.GetUserByID(uid)
	if err != nil || user == nil {
		return errors.New("user not found")
	}

	if user.EmailVerified == true {
		return errors.New("account is already active")
	}

	valid, err := d.repo.IsTokenValid(user.ID, token)
	if err != nil {
		logger.Error("An Error occurred when validating login token. %s", err.Error())
		return errors.New(constant.InternalServerError)
	}
	if valid == false {
		logger.Error(err.Error())
		return errors.New("invalid activation link")
	}

	user.EmailVerified = true
	if err = d.userRepo.Update(user); err != nil {
		logger.Error("An Error occurred while Activating your account, Please try again. %s", err.Error())
		return errors.New("account activation failed")
	}

	err = d.redis.Delete(user.ID)
	if err != nil {
		logger.Error(err.Error())
	}

	return nil
}

func (d *DefaultAuthService) RegisterUser(body *types.Registration, log *log.Entry) error {

	email := utils.ToLower(body.EmailAddress)
	// Check if email address exist
	existingUser, err := d.userRepo.GetByEmail(email)
	if err != nil {
		log.Error(err.Error())
		return errors.New(constant.InternalServerError)
	}
	if existingUser != nil {
		log.Info("Email address already exist. '%s'", body.EmailAddress)
		return errors.New("email address is already registered")
	}

	// Save user details
	user := &model.User{
		FirstName:    body.FirstName,
		LastName:     body.LastName,
		DisplayName:  body.DisplayName,
		EmailAddress: email,
		PhoneNumber:  body.PhoneNumber,
	}

	err = d.userRepo.Create(user)
	if err != nil {
		log.Error("error: occurred when saving new user. %s", err.Error())
		return errors.New("registration was not successful")
	}

	var token string
	if configs.Instance.GetEnv() != configs.SANDBOX {
		token = utils.GenerateRandomString(20)
	} else {
		token = "sandbox"
	}

	err = d.redis.Set(user.ID, token, time.Hour*24)
	if err != nil {
		log.Error("Error occurred when when token %s", err)
	}

	if configs.Instance.GetEnv() != configs.SANDBOX {
		// Send Activate email
		mailTemplate := &sendgrid.ActivationMailRequest{
			Token:    token,
			ToMail:   body.EmailAddress,
			ToName:   fmt.Sprintf("%s %s", body.LastName, body.FirstName),
			FullName: fmt.Sprintf("%s %s", body.LastName, body.FirstName),
			UID:      user.ID,
		}
		err = sendgrid.SendActivateMail(mailTemplate)
		if err != nil {
			log.Error("Error occurred when sending activation email. %s", err.Error())
		}
	}

	return nil
}

// Login
// Handles authentication logic
func (d *DefaultAuthService) Login(body *types.AuthRequest) (*types.LoginResponse, error) {
	response := &types.LoginResponse{}

	user, err := d.userRepo.GetByEmail(strings.ToLower(body.EmailAddress))
	if err != nil {
		log.Error("An error occurred when fetching user profile. %s", err.Error())
		return nil, errors.New(constant.InternalServerError)
	}
	if user == nil {
		return nil, errors.New("user not found")
	}

	// Validate user token
	token, err := d.redis.Get(user.ID)
	if err != nil {
		log.Error("An Error occurred when validating login token. %s", err.Error())
		return nil, errors.New("token validation failed")
	}

	if token == nil || utils.AddToStr(token) != body.Otp {
		return nil, errors.New("invalid otp")
	}

	// Generate JWT for user
	jwToken, err := d.GenerateJWT(user.ID)
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
		LastLogin:     user.LastLogin,
	}

	response.Profile = profile
	response.Token = utils.AddToStr(jwToken)
	// Activity log
	activity := &model.Activity{By: user.ID, Message: "Successful login"}
	go func() {
		user.LastLogin = time.Now()
		if err = d.userRepo.Update(user); err != nil {
			log.Error("User last login failed to be updated. %s", err.Error())
		}

		if err = d.redis.Delete(user.ID); err != nil {
			log.Error("User token failed to be deleted from cache. %s", err.Error())
		}

		if err = d.activity.Create(activity); err != nil {
			log.Error("User activity failed to log. %s", err.Error())
		}
	}()

	return response, nil

}

func (d *DefaultAuthService) RequestToken(body *types.EmailRequest, logger *log.Entry) error {
	if !utils.ValidateEmail(body.EmailAddress) {
		return errors.New("invalid email address")
	}
	email := strings.ToLower(body.EmailAddress)

	// Check if email address exist
	user, err := d.userRepo.GetByEmail(email)
	if err != nil {
		logger.Error(err.Error())
		return errors.New(constant.InternalServerError)
	}

	if user == nil {
		logger.Error(err.Error())
		return errors.New("user not found")
	}

	var token string
	var duration uint = 5
	if configs.Instance.GetEnv() != configs.SANDBOX {
		token = utils.GenerateNumericToken(4)
		var duration uint = 5
		err = d.redis.Set(user.ID, token, time.Minute*time.Duration(duration))
		if err != nil {
			logger.Error("Error occurred when sending token %s", err)
			return errors.New("request failed please try again")
		}
	} else {
		token = "1234"
		err = d.redis.Set(user.ID, token, time.Minute*time.Duration(duration))
		if err != nil {
			logger.Error("Error occurred when sending token %s", err)
			return errors.New("request failed please try again")
		}
	}

	// Activity log
	activity := &model.Activity{
		By:      user.ID,
		Message: "Requested for sign-in token",
	}
	go func() {
		if err = d.activity.Create(activity); err != nil {
			logger.Error("User activity failed to log")
		}
	}()

	if configs.Instance.GetEnv() != configs.SANDBOX {
		mailTemplate := &sendgrid.OTPMailRequest{
			Code:     token,
			ToMail:   user.EmailAddress,
			ToName:   user.LastName + " " + user.FirstName,
			Name:     user.DisplayName,
			Duration: duration,
		}
		err = sendgrid.SendOTPMail(mailTemplate)
		if err != nil {
			logger.Error("Error occurred when sending otp email. %s", err.Error())
		}
	}

	return nil
}

type authCustomClaims struct {
	UserId string `json:"user_id"`
	jwt.StandardClaims
}

func (d *DefaultAuthService) GenerateJWT(userId string) (*string, error) {
	claims := &authCustomClaims{
		userId,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
			Issuer:    configs.Instance.Issuer,
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//encoded string
	t, err := token.SignedString([]byte(configs.Instance.Secret))
	if err != nil {
		log.Error(err.Error())
		return nil, errors.New("token could not be generated")
	}

	return &t, nil
}

func (d *DefaultAuthService) ValidateToken(encodedToken string) (*jwt.Token, error) {
	return jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if _, isvalid := token.Method.(*jwt.SigningMethodHMAC); !isvalid {
			return nil, errors.New("invalid token")
		}
		return []byte(configs.Instance.Secret), nil
	})
}
