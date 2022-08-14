package repository

import (
	"context"
	"errors"
	"github.com/TechBuilder-360/business-directory-backend/internal/common/types"
	"github.com/TechBuilder-360/business-directory-backend/internal/database"
	"github.com/TechBuilder-360/business-directory-backend/internal/database/redis"
	"github.com/TechBuilder-360/business-directory-backend/internal/model"
	"gorm.io/gorm"
	"strings"
)

//go:generate mockgen -destination=../mocks/repository/auth.go -package=repository github.com/TechBuilder-360/business-directory-backend/repository AuthRepository
type AuthRepository interface {
	IsTokenValid(redis *redis.Client, body *types.AuthRequest) (bool, error)
	CreateToken(token *model.Token) error
	DoesUserEmailExist(string) (bool, error)
	GetByEmail(email string) (*model.User, error)
	WithTx(tx *gorm.DB) AuthRepository
	Create(user *model.User) error
	Activate(email string) error
	Get(email string) (*model.User, error)
}

type DefaultAuthRepo struct {
	db *gorm.DB
}

// Get returns user profile
func (r *DefaultAuthRepo) Get(email string) (*model.User, error) {
	user := &model.User{}
	result := r.db.Where("email_address = ?", email).First(&user)
	if result.Error != nil {
		return user, result.Error
	}
	return user, nil
}

func (r *DefaultAuthRepo) Activate(email string) error {
	return r.db.Model(&model.User{}).Where(" email= ?", email).Update("email_verified", true).Error
}
func (r *DefaultAuthRepo) GetByEmail(email string) (*model.User, error) {
	user := &model.User{}
	result := r.db.Where("email_address = ?", email).First(&user)
	if result.Error != nil {
		return user, result.Error
	}
	return user, nil
}
func (r *DefaultAuthRepo) DoesUserEmailExist(email string) (bool, error) {
	user := &model.User{}
	err := r.db.Where("email_address = ?", email).First(&user).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return false, errors.New("request failed")
	}
	return true, nil
}

func (r *DefaultAuthRepo) CreateToken(token *model.Token) error {
	panic("implement me")
}

func (r *DefaultAuthRepo) IsTokenValid(redis *redis.Client, body *types.AuthRequest) (bool, error) {
	token, err := redis.Get(strings.ToLower(body.EmailAddress))
	if err != nil {
		return false, err
	}
	if token == body.Token {
		return true, nil
	} else {
		return false, nil
	}

}
func (r *DefaultAuthRepo) Create(user *model.User) error {
	return r.db.WithContext(context.Background()).Create(user).Error
}

func NewAuthRepository() AuthRepository {
	return &DefaultAuthRepo{
		db: database.ConnectDB(),
	}
}

func (r *DefaultAuthRepo) WithTx(tx *gorm.DB) AuthRepository {
	return &DefaultAuthRepo{db: tx}
}
