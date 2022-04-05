package repository

import (
	"context"
	"encoding/base64"
	"errors"
	"github.com/TechBuilder-360/business-directory-backend/dto"
	"github.com/TechBuilder-360/business-directory-backend/errs"
	"github.com/TechBuilder-360/business-directory-backend/models"
	"github.com/TechBuilder-360/business-directory-backend/utility"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
	"strings"
	"time"
)

// DoesUserEmailExist ...
func (r *DefaultRepo) DoesUserEmailExist(email string) (bool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	count, err := r.User.CountDocuments(ctx, bson.M{"email_address": strings.ToLower(email)})
	if err != nil {
		return false, err
	}

	if count == 0 {
		return false, nil
	}

	return true, nil
}

// RegisterUser ...
func (r *DefaultRepo) RegisterUser(data *dto.Registration) (string,error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	user := models.UserProfile{
		ID:           uuid.New().String(),
		EmailAddress: strings.ToLower(data.EmailAddress),
		FirstName:    data.FirstName,
		LastName:     data.LastName,
		DisplayName:  data.DisplayName,
		PhoneNumber:  data.PhoneNumber,
		CreatedAt: time.Now().Local(),
		UpdatedAt: time.Now().Local(),
	}

	id, err := r.User.InsertOne(ctx, &user)
	if err != nil {
		return "", err
	}

	return id.InsertedID.(string), nil
}

// CreateUserToken ...
func (r *DefaultRepo) CreateUserToken(email string) (string, error) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	token := utility.GenerateNumericToken(6)
	encodedToken := base64.StdEncoding.EncodeToString([]byte(token))
	expireAt:=time.Now().Add(time.Minute * time.Duration(r.Config.TOKENLIFESPAN)).UnixNano()
	tk := models.Token{
		ID:        uuid.New().String(),
		EmailAddress:    strings.ToLower(email),
		Token:     encodedToken,
		ExpireAt:  expireAt,
		CreatedAt: time.Now().Local(),
	}

	_, err := r.Token.InsertOne(ctx, &tk)
	if err != nil {
		return "123456", err
	}

	return token, nil
}

// IsTokenValid ...
func (r *DefaultRepo) IsTokenValid(data *dto.AuthRequest) (bool, *errs.AppError) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	encodedToken := base64.StdEncoding.EncodeToString([]byte(data.Token))
	tk:= &models.Token{}
	filter:=bson.M{"token": encodedToken, "email_address": data.Email}
	err := r.Token.FindOne(ctx, filter).Decode(&tk)
	if err != nil {
		return false, errs.CustomError(http.StatusFailedDependency, utility.SMMERROR, nil)
	}

	presentTime := time.Now().UnixNano()
	if tk.ExpireAt <= presentTime {
		go r.Token.DeleteOne(ctx, filter)
		return true, nil
	}

	return false, errs.CustomError(http.StatusBadRequest, "Token has expired, request new token", nil)
}

// GetUserInformation ...
func (r *DefaultRepo) GetUserInformation(email string) (dto.UserProfile, error) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	user:=dto.UserProfile{}

	filter:=bson.M{"email_address": strings.ToLower(email)}
	opts := options.FindOne().
		SetProjection(bson.D{
			{"created_at", 0},
			{"updated_at", 0},
		})
	err := r.User.FindOne(ctx, filter, opts).Decode(user)
	if err != nil {
		return user, errors.New("user profile could not be retrieved")
	}

	return user, nil
}
