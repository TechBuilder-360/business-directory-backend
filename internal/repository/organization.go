package repository

import (
	"context"
	"errors"
	"github.com/TechBuilder-360/business-directory-backend/internal/database"
	"github.com/TechBuilder-360/business-directory-backend/internal/model"
	"gorm.io/gorm"
)

//go:generate mockgen -destination=../mocks/repository/organisation.go -package=repository github.com/TechBuilder-360/business-directory-backend/repository OrganisationRepository
type OrganisationRepository interface {
	Create(organisation *model.Organisation) error
	Get(organisation *model.Organisation) error
	GetAll(page, limit uint) (*[]model.Organisation, error)
	Find(filter map[string]interface{}, organisation *model.Organisation) error
	Update(organisation *model.Organisation) error
	WithTx(tx *gorm.DB) OrganisationRepository
}

type DefaultOrganisationRepo struct {
	db *gorm.DB
}

func (d DefaultOrganisationRepo) Find(filter map[string]interface{}, organisation *model.Organisation) error {
	err := d.db.Where(filter).Find(&organisation).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("record not found")
	} else if err != nil {
		return errors.New("not found")
	}
	return nil
}

func (d *DefaultOrganisationRepo) Create(organisation *model.Organisation) error {
	return d.db.WithContext(context.Background()).Create(organisation).Error
}

func (d *DefaultOrganisationRepo) Get(organisation *model.Organisation) error {
	panic("implement me")
}

func (d *DefaultOrganisationRepo) GetAll(page, limit uint) (*[]model.Organisation, error) {
	panic("implement me")
}

func (d *DefaultOrganisationRepo) Update(organisation *model.Organisation) error {
	panic("implement me")
}

func (d *DefaultOrganisationRepo) UpdateStatus(organisation *model.Organisation) error {
	panic("implement me")
}

func (d *DefaultOrganisationRepo) WithTx(tx *gorm.DB) OrganisationRepository {
	return &DefaultOrganisationRepo{
		db: tx,
	}
}

func NewOrganisationRepository() OrganisationRepository {
	return &DefaultOrganisationRepo{
		db: database.ConnectDB(),
	}
}
