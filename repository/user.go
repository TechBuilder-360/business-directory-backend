package repository

import (
	"context"
	"github.com/TechBuilder-360/business-directory-backend/dto"
	"github.com/TechBuilder-360/business-directory-backend/models"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

// DoesUserEmailExist ...
func (r *DefaultRepo) DoesUserEmailExist(email string) (bool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	count, err := r.User.CountDocuments(ctx, bson.M{"email_address": email})
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
		EmailAddress: data.EmailAddress,
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
