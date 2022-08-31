package repository

import (
	"github.com/TechBuilder-360/business-directory-backend/internal/database"
	"github.com/TechBuilder-360/business-directory-backend/internal/model"
	"gorm.io/gorm"
)

//go:generate mockgen -destination=../mocks/repository/activity.go -package=repository github.com/TechBuilder-360/business-directory-backend/repository ActivityRepository
type ActivityRepository interface {
	Create(activity *model.Activity) error
	Get(activity *model.Activity) error
	WithTx(tx *gorm.DB) ActivityRepository
}

type DefaultActivityRepo struct {
	db *gorm.DB
}

func (r *DefaultActivityRepo) Get(activity *model.Activity) error {
	panic("implement me")
}

func (r *DefaultActivityRepo) WithTx(tx *gorm.DB) ActivityRepository {
	panic("implement me")
}

func (r *DefaultActivityRepo) Create(activity *model.Activity) error {
	return r.db.Create(activity).Error
}

func NewActivityRepository() ActivityRepository {
	return &DefaultActivityRepo{
		db: database.ConnectDB(),
	}
}
