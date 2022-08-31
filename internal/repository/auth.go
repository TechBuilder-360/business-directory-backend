package repository

import (
	"github.com/TechBuilder-360/business-directory-backend/internal/common/utils"
	"github.com/TechBuilder-360/business-directory-backend/internal/database"
	"github.com/TechBuilder-360/business-directory-backend/internal/database/redis"
	"github.com/TechBuilder-360/business-directory-backend/internal/model"
	"gorm.io/gorm"
)

//go:generate mockgen -destination=../mocks/repository/auth.go -package=repository github.com/TechBuilder-360/business-directory-backend/repository AuthRepository
type AuthRepository interface {
	IsTokenValid(email, token string) (bool, error)
	CreateToken(token *model.Token) error
	WithTx(tx *gorm.DB) AuthRepository
}

type DefaultAuthRepo struct {
	db    *gorm.DB
	redis *redis.Client
}

func (r *DefaultAuthRepo) CreateToken(token *model.Token) error {
	panic("implement me")
}

func (r *DefaultAuthRepo) IsTokenValid(key, token string) (bool, error) {
	rToken, err := redis.RedisClient().Get(key)
	if err != nil {
		return false, err
	}

	isValid := utils.AddToStr(rToken) == token

	return isValid, nil
}

func NewAuthRepository() AuthRepository {
	return &DefaultAuthRepo{
		db: database.ConnectDB(),
	}
}

func (r *DefaultAuthRepo) WithTx(tx *gorm.DB) AuthRepository {
	return &DefaultAuthRepo{
		db:    tx,
		redis: redis.RedisClient(),
	}
}
