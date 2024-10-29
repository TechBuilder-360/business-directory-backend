package repository

import (
	"github.com/TechBuilder-360/business-directory-backend/domain/business/model"
	"github.com/TechBuilder-360/business-directory-backend/internal/common/types"
	"github.com/TechBuilder-360/business-directory-backend/internal/database"
	"gorm.io/gorm"
)

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
