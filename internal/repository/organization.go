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
	GetByPublicKey(publicKey string) (*model.Organisation, error)
	GetAll(page, limit uint) (*[]model.Organisation, error)
	Find(filter map[string]interface{}) ([]model.Organisation, error)
	Update(organisation *model.Organisation) error
	WithTx(tx *gorm.DB) OrganisationRepository
	GetOrganisationByName(name string) (*model.Organisation, error)
	AddOrganisationMember(member *model.OrganisationMember) error
}

type DefaultOrganisationRepo struct {
	db *gorm.DB
}

func (d *DefaultOrganisationRepo) GetByPublicKey(publicKey string) (*model.Organisation, error) {
	organisation := &model.Organisation{}
	err := d.db.WithContext(context.Background()).Where(&model.Organisation{PublicKey: publicKey}).First(organisation).Error
	if errors.Is(err, gorm.ErrRecordNotFound) == false {
		return nil, errors.New("could not fetch organisation")
	}

	return organisation, nil
}

func (d *DefaultOrganisationRepo) GetOrganisationByName(name string) (*model.Organisation, error) {
	var organisation *model.Organisation
	err := d.db.Where("organisation_name=?", name).First(&organisation).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return organisation, nil
}
func (d *DefaultOrganisationRepo) Find(filter map[string]interface{}) ([]model.Organisation, error) {
	var organisation []model.Organisation
	err := d.db.Where(filter).Find(&organisation)
	if errors.Is(err.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	} else if err.Error != nil {
		return nil, err.Error
	}
	return organisation, nil
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

func (d *DefaultOrganisationRepo) AddOrganisationMember(member *model.OrganisationMember) error {
	return d.db.WithContext(context.Background()).Create(member).Error
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
