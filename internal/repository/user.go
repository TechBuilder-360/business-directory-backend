package repository

import (
	"context"
	"errors"
	"github.com/TechBuilder-360/business-directory-backend/internal/database"
	"github.com/TechBuilder-360/business-directory-backend/internal/model"
	"gorm.io/gorm"
)

//go:generate mockgen -destination=../mocks/repository/user.go -package=repository github.com/TechBuilder-360/business-directory-backend/repository UserRepository
type UserRepository interface {
	GetByID(id string) (*model.User, error)
	GetByEmail(email string) (*model.User, error)
	Update(user *model.User) error
	Create(user *model.User) error
	WithTx(tx *gorm.DB) UserRepository
}

type DefaultUserRepo struct {
	db *gorm.DB
}

func NewUserRepository() UserRepository {
	return &DefaultUserRepo{
		db: database.ConnectDB(),
	}
}

func (r *DefaultUserRepo) WithTx(tx *gorm.DB) UserRepository {
	return &DefaultUserRepo{db: tx}
}

func (r *DefaultUserRepo) Update(user *model.User) error {
	panic("implement me")
}

func (r *DefaultUserRepo) GetByID(id string) (*model.User, error) {
	user := &model.User{}
	if err := r.db.Where("id = ?", id).First(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (r *DefaultUserRepo) Create(user *model.User) error {
	return r.db.WithContext(context.Background()).Create(user).Error
}

func (r *DefaultUserRepo) GetByEmail(email string) (*model.User, error) {
	user := &model.User{}
	err := r.db.Where("email_address = ?", email).First(user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return user, nil
}
