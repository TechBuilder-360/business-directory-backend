package repository

import (
	"errors"
	"github.com/TechBuilder-360/business-directory-backend/models"
	"github.com/TechBuilder-360/business-directory-backend/utility"
	"gorm.io/gorm"
)

//go:generate mockgen -destination=../mocks/repository/branch.go -package=repository github.com/TechBuilder-360/business-directory-backend/repository BranchRepository
type BranchRepository interface {
	Create(branch *models.Branch) error
	Get(branch *models.Branch) error
	Find(filter map[string]interface{}, branch *models.Branch) error
	Update(branch *models.Branch) error
	GetAll(organisationId string, page, limit uint) (*[]models.Branch, error)
	WithTx(tx *gorm.DB) OrganisationRepository
}

type DefaultBranchRepo struct {
	db *gorm.DB
}

func (d DefaultBranchRepo) Create(branch *models.Branch) error {
	panic("implement me")
}

func (d DefaultBranchRepo) Get(branch *models.Branch) error {
	panic("implement me")
}

func (d DefaultBranchRepo) Find(filter map[string]interface{}, branch *models.Branch) error {
	err := d.db.Where(filter).Find(&branch).Error
	if errors.Is(err, gorm.ErrRecordNotFound){
		return errors.New("record not found")
	}else if err !=nil {
		return errors.New(utility.SMMERROR)
	}
	return nil
}

func (d DefaultBranchRepo) Update(branch *models.Branch) error {
	panic("implement me")
}

func (d DefaultBranchRepo) GetAll(organisationId string, page, limit uint) (*[]models.Branch, error) {
	panic("implement me")
}

func (d DefaultBranchRepo) WithTx(tx *gorm.DB) OrganisationRepository {
	panic("implement me")
}

func NewBranchRepository(db *gorm.DB) BranchRepository {
	return &DefaultBranchRepo{
		db: db,
	}
}
