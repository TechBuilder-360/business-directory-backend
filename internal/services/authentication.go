package services

import (
	"errors"
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
	ActivateEmail(token string, email string, log *log.Entry) (string, error)
	Login(body *types.AuthRequest) (*types.JWTResponse, error)
	AuthEmail(body *types.EmailRequest) (string, *model.User, error)
	GenerateToken(userID string) (string, error)
	ValidateToken(encodedToken string) (*jwt.Token, error)
	ResendToken(body *types.EmailRequest) (string, *model.User, error)
}

type DefaultAuthService struct {
	repo     repository.AuthRepository
	userRepo repository.UserRepository
	activity repository.ActivityRepository
	redis    *redis.Client
}

func (d *DefaultAuthService) ActivateEmail(token string, email string, log *log.Entry) (string, error) {

	// Update with conditions
	body := &types.AuthRequest{
		EmailAddress: email,
		Token:        token,
	}
	ok, err := d.repo.IsTokenValid(d.redis, body)
	if err != nil {
		log.Error("An Error occurred when validating login token. %s", err.Error())
		return "", err
	}
	if ok {
		if err := d.repo.Activate(email); err != nil {
			log.Error("An Error occurred while Activating your account, Please try again. %s", err.Error())
			return "", err
		}
	}
	return " Your account has been activated, Please proceed to login", nil
}
func (u *DefaultAuthService) RegisterUser(body *types.Registration, log *log.Entry) error {

	email := utils.ToLower(body.EmailAddress)

	// Check if email address exist
	ok, err := u.repo.DoesUserEmailExist(email)
	if err != nil {
		return err
	}
	if ok {
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

	err = u.repo.Create(user)
	if err != nil {
		log.Error("error: occurred when saving new user. %s", err.Error())
		return errors.New("registration was not successful")
	}
	token := utils.GenerateNumericToken(20)
	err = u.redis.Set(email, token, time.Minute*20)
	if err != nil {
		log.Error("Error occurred when when token %s", err)
		return errors.New("registration was not successful")
	}

	// Send Activate email

	mailTemplate := &types.MailTemplate{
		Header:  "Activation Link",
		Subject: "Activate your Account",
		Token:   token,
		ToEmail: body.EmailAddress,
		ToName:  body.FirstName + " " + body.LastName,
	}
	err = sendgrid.SendMail(mailTemplate)
	if err != nil {
		log.Error("Error occurred when sending activation email. %s", err.Error())
		return err
	}
	return nil
}

// Login
// Handles authentication logic
func (d *DefaultAuthService) Login(body *types.AuthRequest) (*types.JWTResponse, error) {
	response := &types.JWTResponse{}
	// Validate user token
	ok, err := d.repo.IsTokenValid(d.redis, body)
	if err != nil {
		log.Error("An Error occurred when validating login token. %s", err.Error())
		return nil, err
	}

	if ok {

		user, err := d.repo.Get(strings.ToLower(body.EmailAddress))
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

	}

	return response, nil

}

func (d *DefaultAuthService) ResendToken(body *types.EmailRequest) (string, *model.User, error) {
	if !utils.ValidateEmail(body.EmailAddress) {
		return "", nil, errors.New("invalid email address")
	}
	email := strings.ToLower(body.EmailAddress)

	// Check if email address exist
	data, err := d.repo.GetByEmail(email)
	if err != nil {
		log.Error("An Error occurred while checking if user email exist. %s", err.Error())
		return "", nil, errors.New("email does not exist")
	}
	if data.ID == "" {
		log.Info("Email address does not exist. '%s'", email)
		return "", nil, errors.New("email not found")
	}

	token := utils.GenerateNumericToken(4)
	err = d.redis.Set(email, token, 2000)
	if err != nil {
		log.Error("Error occurred when sending token %s", err)
		return "", nil, err
	}

	// TODO: Send Token to user email

	// Activity log
	activity := &model.Activity{Message: "Requested for sign in token"}
	go func() {
		if err := d.activity.Create(activity); err != nil {
			log.Error("User activity failed to log")
		}
	}()
	dur, _ := time.ParseDuration("5mins")
	mailTemplate := &types.MailTemplate{
		Header:   "Activation Link",
		Subject:  "Activate your Account",
		Token:    token,
		ToEmail:  data.EmailAddress,
		ToName:   data.FirstName + " " + data.LastName,
		Duration: dur,
	}

	err = sendgrid.SendMail(mailTemplate)
	if err != nil {
		log.Error("Error occurred when sending verification email. %s", err.Error())
		return "", nil, err
	}
	return "sign in code has been re-sent ", data, nil
}
func (d *DefaultAuthService) AuthEmail(body *types.EmailRequest) (string, *model.User, error) {

	if !utils.ValidateEmail(body.EmailAddress) {
		return "", nil, errors.New("invalid email address")
	}
	email := strings.ToLower(body.EmailAddress)

	// Check if email address exist
	data, err := d.repo.GetByEmail(email)
	if err != nil {
		log.Error("An Error occurred while checking if user email exist. %s", err.Error())
		return "", nil, errors.New("email does not exist")
	}
	if data.ID == "" {
		log.Info("Email address does not exist. '%s'", email)
		return "", nil, errors.New("email not found")
	}

	token := utils.GenerateNumericToken(4)
	err = d.redis.Set(email, token, 2000)
	if err != nil {
		log.Error("Error occurred when when token %s", err)
		return "", nil, err
	}

	// TODO: Send Token to user email

	// Activity log
	activity := &model.Activity{Message: "Requested for sign in token"}
	go func() {
		if err := d.activity.Create(activity); err != nil {
			log.Error("User activity failed to log")
		}
	}()
	dur, _ := time.ParseDuration("5mins")
	mailTemplate := &types.MailTemplate{
		Header:   "Activation Link",
		Subject:  "Activate your Account",
		Token:    token,
		ToEmail:  data.EmailAddress,
		ToName:   data.FirstName + " " + data.LastName,
		Duration: dur,
	}

	err = sendgrid.SendMail(mailTemplate)
	if err != nil {
		log.Error("Error occurred when sending verification email. %s", err.Error())
		return "", nil, err
	}
	return "Please check your mail for sign on Code ", data, nil
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
		repo:  repository.NewAuthRepository(),
		redis: redis.RedisClient(),
	}
}
