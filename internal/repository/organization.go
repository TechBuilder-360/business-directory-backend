package repository

import "C"
import (
	"context"
	"errors"
	"github.com/TechBuilder-360/business-directory-backend/internal/common/types"
	"github.com/TechBuilder-360/business-directory-backend/internal/database"
	"github.com/TechBuilder-360/business-directory-backend/internal/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

//go:generate mockgen -destination=../mocks/repository/organisation.go -package=repository github.com/TechBuilder-360/business-directory-backend/repository OrganisationRepository
type OrganisationRepository interface {
	Create(organisation *model.Organisation) error
	Get(id string) (*model.Organisation, error)
	GetByPublicKey(publicKey string) (*model.Organisation, error)
	GetAll(query types.Query) ([]model.Organisation, error)
	Total(query types.Query) (int64, error)
	Find(filter map[string]interface{}) ([]model.Organisation, error)
	Update(organisation *model.Organisation) error
	WithTx(tx *gorm.DB) OrganisationRepository
	GetOrganisationByName(name string) (*model.Organisation, error)
	AddOrganisationMember(member *model.Member) error
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
	err := d.db.Where("name=?", name).First(&organisation).Error
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

func (d *DefaultOrganisationRepo) Get(id string) (*model.Organisation, error) {
	var organisation *model.Organisation
	err := d.db.Preload(clause.Associations).Where("id = ?", id).First(&organisation).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("not found")
	}
	if err != nil {
		return nil, errors.New("an error occurred")
	}

	return organisation, nil
}

func (d *DefaultOrganisationRepo) GetAll(query types.Query) ([]model.Organisation, error) {
	var organisations []model.Organisation
	stmt := d.db.Model(&model.Organisation{}).Limit(query.PageSize).Offset(query.PaginationOffset())
	if query.Search != "" {
		stmt = stmt.Where("name ilike '%?%'", query.Search)
	}
	if err := stmt.Find(&organisations).Error; err != nil {
		return nil, err
	}

	return organisations, nil
}

func (d *DefaultOrganisationRepo) Total(query types.Query) (int64, error) {
	total := int64(0)
	stmt := d.db.Model(&model.Organisation{})
	if query.Search != "" {
		stmt = stmt.Where("name ilike '%?%'", query.Search)
	}
	if err := stmt.Count(&total).Error; err != nil {
		return total, err
	}

	return total, nil
}

func (d *DefaultOrganisationRepo) Update(organisation *model.Organisation) error {
	ctx := context.Background()
	return d.db.WithContext(ctx).Save(organisation).Error
}

func (d *DefaultOrganisationRepo) AddOrganisationMember(member *model.Member) error {
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
