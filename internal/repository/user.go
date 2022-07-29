package repository

import (
	"context"
	"github.com/TechBuilder-360/business-directory-backend/internal/database"
	"github.com/TechBuilder-360/business-directory-backend/internal/model"
	"gorm.io/gorm"
)

//go:generate mockgen -destination=../mocks/repository/user.go -package=repository github.com/TechBuilder-360/business-directory-backend/repository UserRepository
type UserRepository interface {
	DoesUserEmailExist(string) (bool, error)
	Create(user *model.User) error
	Get(user *model.User) error
	GetByEmail(user *model.User) error
	Update() error
	Deactivate() error
	Activate() error
	WithTx(tx *gorm.DB) UserRepository
}

type DefaultUserRepo struct {
	db *gorm.DB
}

func (r *DefaultUserRepo) GetByEmail(user *model.User) error {
	panic("implement me")
}

func NewUserRepository() UserRepository {
	return &DefaultUserRepo{
		db: database.ConnectDB(),
	}
}

func (r *DefaultUserRepo) WithTx(tx *gorm.DB) UserRepository {
	return &DefaultUserRepo{db: tx}
}

func (r *DefaultUserRepo) Deactivate() error {
	panic("implement me")
}

func (r *DefaultUserRepo) Activate() error {
	panic("implement me")
}

func (r *DefaultUserRepo) Update() error {
	panic("implement me")
}

// DoesUserEmailExist ...
func (r *DefaultUserRepo) DoesUserEmailExist(email string) (bool, error) {
	panic("not implemented")
	return true, nil
}

// Create ...
func (r *DefaultUserRepo) Create(user *model.User) error {
	return r.db.WithContext(context.Background()).Create(user).Error
}

// Get returns user profile
func (r *DefaultUserRepo) Get(user *model.User) error {
	panic("not implemented")
	return nil
}
