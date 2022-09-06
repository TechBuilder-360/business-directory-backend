package services

import (
	"github.com/TechBuilder-360/business-directory-backend/internal/model"
	"github.com/TechBuilder-360/business-directory-backend/internal/repository"
	"gorm.io/gorm"
)

//go:generate mockgen -destination=../mocks/services/branch.go -package=services github.com/TechBuilder-360/business-directory-backend/services BranchService
type BranchService interface {
	Create(branch *model.Branch) error
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
