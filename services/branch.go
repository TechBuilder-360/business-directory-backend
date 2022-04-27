package services

import (
	"gorm.io/gorm"

	"github.com/TechBuilder-360/business-directory-backend/repository"
)


//go:generate mockgen -destination=../mocks/services/branch.go -package=services github.com/TechBuilder-360/business-directory-backend/services BranchService
type BranchService interface {
}


type DefaultBranchService struct {
	repo repository.BranchRepository
	activity repository.ActivityRepository
	db *gorm.DB
}


func NewBranchService(repo repository.BranchRepository, activity repository.ActivityRepository) BranchService {
	return &DefaultBranchService{repo: repo, activity: activity}
}
