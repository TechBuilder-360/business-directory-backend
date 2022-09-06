package repository

import (
	"github.com/TechBuilder-360/business-directory-backend/internal/common/types"
	"github.com/TechBuilder-360/business-directory-backend/internal/database"
	"github.com/TechBuilder-360/business-directory-backend/internal/model"
	"gorm.io/gorm"
)

//go:generate mockgen -destination=../mocks/repository/role.go -package=repository github.com/TechBuilder-360/business-directory-backend/repository RoleRepository
type RoleRepository interface {
	GetByName(roleName types.RoleType) (*model.Role, error)
}

type DefaultRoleRepo struct {
	db *gorm.DB
}

func (r *DefaultRoleRepo) GetByName(roleName types.RoleType) (*model.Role, error) {
	role := &model.Role{}
	err := r.db.Where(&model.Role{Name: roleName}).First(role).Error
	if err != nil {
		return nil, err
	}

	return role, nil
}

func NewRoleRepository() RoleRepository {
	return &DefaultRoleRepo{
		db: database.ConnectDB(),
	}
}
