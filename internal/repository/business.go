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

//go:generate mockgen -destination=../mocks/repository/Business.go -package=repository github.com/TechBuilder-360/business-directory-backend/repository BusinessRepository
type BusinessRepository interface {
	Create(Business *model.Business) error
	Get(id string) (*model.Business, error)
	GetByPublicKey(publicKey string) (*model.Business, error)
	GetAll(query types.Query) ([]model.Business, error)
	Total(query types.Query) (int64, error)
	Find(filter map[string]interface{}) ([]model.Business, error)
	Update(Business *model.Business) error
	WithTx(tx *gorm.DB) BusinessRepository
	GetBusinessByName(name string) (*model.Business, error)
	AddBusinessMember(member *model.Member) error
}

type DefaultBusinessRepo struct {
	db *gorm.DB
}

func (d *DefaultBusinessRepo) GetByPublicKey(publicKey string) (*model.Business, error) {
	Business := &model.Business{}
	err := d.db.WithContext(context.Background()).Where(&model.Business{PublicKey: publicKey}).First(Business).Error
	if errors.Is(err, gorm.ErrRecordNotFound) == false {
		return nil, errors.New("could not fetch Business")
	}

	return Business, nil
}

func (d *DefaultBusinessRepo) GetBusinessByName(name string) (*model.Business, error) {
	var Business *model.Business
	err := d.db.Where("lower(name) = lower(?)", name).First(&Business).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return Business, nil
}
func (d *DefaultBusinessRepo) Find(filter map[string]interface{}) ([]model.Business, error) {
	var Business []model.Business
	err := d.db.Where(filter).Find(&Business)
	if errors.Is(err.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	} else if err.Error != nil {
		return nil, err.Error
	}
	return Business, nil
}

func (d *DefaultBusinessRepo) Create(Business *model.Business) error {
	return d.db.WithContext(context.Background()).Create(Business).Error
}

func (d *DefaultBusinessRepo) Get(id string) (*model.Business, error) {
	var Business *model.Business
	err := d.db.Preload(clause.Associations).Where("id = ?", id).First(&Business).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("not found")
	}
	if err != nil {
		return nil, errors.New("an apiError occurred")
	}

	return Business, nil
}

func (d *DefaultBusinessRepo) GetAll(query types.Query) ([]model.Business, error) {
	var Businesss []model.Business
	stmt := d.db.Model(&model.Business{}).Limit(query.PageSize).Offset(query.PaginationOffset())
	if query.Search != "" {
		stmt = stmt.Where("name ilike '%?%'", query.Search)
	}
	if err := stmt.Find(&Businesss).Error; err != nil {
		return nil, err
	}

	return Businesss, nil
}

func (d *DefaultBusinessRepo) Total(query types.Query) (int64, error) {
	total := int64(0)
	stmt := d.db.Model(&model.Business{})
	if query.Search != "" {
		stmt = stmt.Where("name ilike '%?%'", query.Search)
	}
	if err := stmt.Count(&total).Error; err != nil {
		return total, err
	}

	return total, nil
}

func (d *DefaultBusinessRepo) Update(Business *model.Business) error {
	ctx := context.Background()
	return d.db.WithContext(ctx).Save(Business).Error
}

func (d *DefaultBusinessRepo) AddBusinessMember(member *model.Member) error {
	return d.db.WithContext(context.Background()).Create(member).Error
}

func (d *DefaultBusinessRepo) WithTx(tx *gorm.DB) BusinessRepository {
	return &DefaultBusinessRepo{
		db: tx,
	}
}

func NewBusinessRepository() BusinessRepository {
	return &DefaultBusinessRepo{
		db: database.ConnectDB(),
	}
}
