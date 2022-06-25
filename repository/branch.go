package repository

import (
	"errors"
	"github.com/TechBuilder-360/business-directory-backend/database"
	"github.com/TechBuilder-360/business-directory-backend/model"
	"github.com/TechBuilder-360/business-directory-backend/utility"
	"gorm.io/gorm"
)

//go:generate mockgen -destination=../mocks/repository/branch.go -package=repository github.com/TechBuilder-360/business-directory-backend/repository BranchRepository
type BranchRepository interface {
	Create(branch *model.Branch) error
	Get(branch *model.Branch) error
	Find(filter map[string]interface{}, branch *model.Branch) error
	Update(branch *model.Branch) error
	GetAll(organisationId string, page, limit uint) (*[]model.Branch, error)
	WithTx(tx *gorm.DB) BranchRepository
}

type DefaultBranchRepo struct {
	db *gorm.DB
}

func (d DefaultBranchRepo) Create(branch *model.Branch) error {
	panic("implement me")
}

func (d DefaultBranchRepo) Get(branch *model.Branch) error {
	panic("implement me")
}

func (d DefaultBranchRepo) Find(filter map[string]interface{}, branch *model.Branch) error {
	err := d.db.Where(filter).Find(&branch).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("record not found")
	} else if err != nil {
		return errors.New(utility.SMMERROR)
	}
	return nil
}

func (d DefaultBranchRepo) Update(branch *model.Branch) error {
	panic("implement me")
}

func (d DefaultBranchRepo) GetAll(organisationId string, page, limit uint) (*[]model.Branch, error) {
	panic("implement me")
}

func (d DefaultBranchRepo) WithTx(tx *gorm.DB) BranchRepository {
	return &DefaultBranchRepo{db: tx}
}

func NewBranchRepository() BranchRepository {
	return &DefaultBranchRepo{
		db: database.ConnectDB(),
	}
}
