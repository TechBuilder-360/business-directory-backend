package repository

import (
	"github.com/TechBuilder-360/business-directory-backend/dto"
	"github.com/TechBuilder-360/business-directory-backend/models"
	"gorm.io/gorm"
)

//go:generate mockgen -destination=../mocks/repository/auth.go -package=repository github.com/TechBuilder-360/business-directory-backend/repository AuthRepository
type AuthRepository interface {
	IsTokenValid(body *dto.AuthRequest) error
	CreateToken(token *models.Token) error
	WithTx(tx *gorm.DB) AuthRepository
}

type DefaultAuthRepo struct {
	db *gorm.DB
}

func (r *DefaultAuthRepo) CreateToken(token *models.Token) error {
	panic("implement me")
}

func (r *DefaultAuthRepo) IsTokenValid(body *dto.AuthRequest) error {
	panic("implement me")
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &DefaultAuthRepo{
		db: db,
	}
}

func (r *DefaultAuthRepo) WithTx(tx *gorm.DB) AuthRepository {
	return NewAuthRepository(tx)
}