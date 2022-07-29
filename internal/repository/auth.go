package repository

import (
	"github.com/TechBuilder-360/business-directory-backend/internal/common/types"
	"github.com/TechBuilder-360/business-directory-backend/internal/database"
	"github.com/TechBuilder-360/business-directory-backend/internal/model"
	"gorm.io/gorm"
)

//go:generate mockgen -destination=../mocks/repository/auth.go -package=repository github.com/TechBuilder-360/business-directory-backend/repository AuthRepository
type AuthRepository interface {
	IsTokenValid(body *types.AuthRequest) error
	CreateToken(token *model.Token) error
	WithTx(tx *gorm.DB) AuthRepository
}

type DefaultAuthRepo struct {
	db *gorm.DB
}

func (r *DefaultAuthRepo) CreateToken(token *model.Token) error {
	panic("implement me")
}

func (r *DefaultAuthRepo) IsTokenValid(body *types.AuthRequest) error {
	panic("implement me")
}

func NewAuthRepository() AuthRepository {
	return &DefaultAuthRepo{
		db: database.ConnectDB(),
	}
}

func (r *DefaultAuthRepo) WithTx(tx *gorm.DB) AuthRepository {
	return &DefaultAuthRepo{db: tx}
}
