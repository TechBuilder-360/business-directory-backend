package repository

import (
	"context"
	"errors"
	"github.com/TechBuilder-360/business-directory-backend/internal/common/types"
	"github.com/TechBuilder-360/business-directory-backend/internal/database"
	"github.com/TechBuilder-360/business-directory-backend/internal/model"
	"gorm.io/gorm"
	"math"
)

//go:generate mockgen -destination=../mocks/repository/branch.go -package=repository github.com/TechBuilder-360/business-directory-backend/repository BranchRepository
type BranchRepository interface {
	Create(branch *model.Branch) error
	Get(id string) (*model.Branch, error)
	GetByOrganisation(organisationId string) ([]model.Branch, error)
	Find(filter map[string]interface{}, branch *model.Branch) error
	Update(branch *model.Branch) error
	GetAll(page int) (*types.PaginatedResponse, error)
	WithTx(tx *gorm.DB) BranchRepository
}

type DefaultBranchRepo struct {
	db *gorm.DB
}

func (d *DefaultBranchRepo) Create(branch *model.Branch) error {
	return d.db.WithContext(context.Background()).Create(branch).Error
}

func (d *DefaultBranchRepo) Get(id string) (*model.Branch, error) {
	var branch *model.Branch
	err := d.db.Where("id = ?", id).First(branch).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, errors.New("an error occurred")
	}

	return branch, nil
}

func (d *DefaultBranchRepo) GetByOrganisation(organisationId string) ([]model.Branch, error) {
	var branch []model.Branch
	err := d.db.Where("organisation_id = ? and active = true", organisationId).Find(&branch).Error
	if err != nil {
		return nil, errors.New("an error occurred")
	}

	return branch, nil
}

func (d *DefaultBranchRepo) Find(filter map[string]interface{}, branch *model.Branch) error {
	err := d.db.Where(filter).Find(&branch).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("record not found")
	} else if err != nil {
		return errors.New("not found")
	}
	return nil
}

func (d *DefaultBranchRepo) Update(branch *model.Branch) error {
	panic("implement me")
}

func (d *DefaultBranchRepo) GetAll(page int) (*types.PaginatedResponse, error) {
	var branch []model.Branch
	limit := 10
	brans := &types.Branch{}
	offset := (page - 1) * limit
	query := d.db.Select(brans).Limit(limit).Offset(offset).Find(&branch)
	if query.Error != nil {
		return nil, query.Error
	}

	data := &types.PaginatedResponse{
		Page:    page,
		PerPage: limit,
		Total:   int64(math.Ceil(float64(query.RowsAffected) / float64(limit))),
		Data:    query,
	}

	return data, nil
}

func (d *DefaultBranchRepo) WithTx(tx *gorm.DB) BranchRepository {
	return &DefaultBranchRepo{db: tx}
}

func NewBranchRepository() BranchRepository {
	return &DefaultBranchRepo{
		db: database.ConnectDB(),
	}
}
