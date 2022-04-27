package repository

import (
	"context"
	"errors"
	"github.com/TechBuilder-360/business-directory-backend/models"
	"github.com/TechBuilder-360/business-directory-backend/utility"
	"gorm.io/gorm"
)

//go:generate mockgen -destination=../mocks/repository/organisation.go -package=repository github.com/TechBuilder-360/business-directory-backend/repository OrganisationRepository
type OrganisationRepository interface {
	Create(organisation *models.Organisation) error
	Get(organisation *models.Organisation) error
	GetAll(page, limit uint) (*[]models.Organisation, error)
	Find(filter map[string]interface{}, organisation *models.Organisation) error
	Update(organisation *models.Organisation) error
	WithTx(tx *gorm.DB) OrganisationRepository
}

type DefaultOrganisationRepo struct {
	db *gorm.DB
}

func (d DefaultOrganisationRepo) Find(filter map[string]interface{}, organisation *models.Organisation) error {
	err := d.db.Where(filter).Find(&organisation).Error
	if errors.Is(err, gorm.ErrRecordNotFound){
		return errors.New("record not found")
	}else if err !=nil {
		return errors.New(utility.SMMERROR)
	}
	return nil
}

func (d *DefaultOrganisationRepo) Create(organisation *models.Organisation) error {
	return d.db.WithContext(context.Background()).Create(organisation).Error
}

func (d DefaultOrganisationRepo) Get(organisation *models.Organisation) error {
	panic("implement me")
}

func (d DefaultOrganisationRepo) GetAll(page, limit uint) (*[]models.Organisation, error) {
	panic("implement me")
}

func (d DefaultOrganisationRepo) Update(organisation *models.Organisation) error {
	panic("implement me")
}

func (d DefaultOrganisationRepo) UpdateStatus(organisation *models.Organisation) error {
	panic("implement me")
}

func (d DefaultOrganisationRepo) WithTx(tx *gorm.DB) OrganisationRepository {
	panic("implement me")
}

func NewOrganisationRepository(db *gorm.DB) OrganisationRepository {
	return &DefaultOrganisationRepo{
		db: db,
	}
}
