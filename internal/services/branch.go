package services

import (
	"errors"
	"github.com/TechBuilder-360/business-directory-backend/internal/common/constant"
	"github.com/TechBuilder-360/business-directory-backend/internal/common/types"
	"github.com/TechBuilder-360/business-directory-backend/internal/model"
	"github.com/TechBuilder-360/business-directory-backend/internal/repository"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

//go:generate mockgen -destination=../mocks/services/branch.go -package=services github.com/TechBuilder-360/business-directory-backend/services BranchService
type BranchService interface {
	Create(branch *model.Branch) error
	GetSingleBranch(id string, logger *log.Entry) (*model.Branch, error)
	GetAllBranch(page int, logger *log.Entry) (*types.DataView, error)
}

type DefaultBranchService struct {
	branchRepo repository.BranchRepository
	db         *gorm.DB
}

func NewBranchService() BranchService {
	return &DefaultBranchService{
		branchRepo: repository.NewBranchRepository(),
	}
}

func (b *DefaultBranchService) Create(branch *model.Branch) error {
	err := b.branchRepo.Create(branch)
	if err != nil {
		return err
	}

	return nil
}

func (o *DefaultBranchService) GetSingleBranch(id string, logger *log.Entry) (*model.Branch, error) {
	branch := &model.Branch{}
	branch.Base.ID = id
	err := o.branchRepo.Get(branch)
	if err != nil {
		logger.Error(err)
		return nil, errors.New(constant.InternalServerError)
	}
	return branch, nil
}

func (o *DefaultBranchService) GetAllBranch(page int, logger *log.Entry) (*types.DataView, error) {
	data, err := o.branchRepo.GetAll(page)
	if err != nil {
		logger.Error(err)
		return nil, errors.New(constant.InternalServerError)
	}
	return data, nil
}
