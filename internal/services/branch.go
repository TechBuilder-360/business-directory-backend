package services

import (
	"github.com/TechBuilder-360/business-directory-backend/internal/repository"
	"gorm.io/gorm"
)

//go:generate mockgen -destination=../mocks/services/branch.go -package=services github.com/TechBuilder-360/business-directory-backend/services BranchService
type BranchService interface {
}

type DefaultBranchService struct {
	repo repository.BranchRepository
	db   *gorm.DB
}

func NewBranchService() BranchService {
	return &DefaultBranchService{repo: repository.NewBranchRepository()}
}
