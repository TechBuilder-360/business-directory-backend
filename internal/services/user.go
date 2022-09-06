package services

import (
	"github.com/TechBuilder-360/business-directory-backend/internal/repository"
)

//go:generate mockgen -destination=../mocks/services/mockService.go -package=services github.com/TechBuilder-360/business-directory-backend/services UserService
type UserService interface {
}

type DefaultUserService struct {
	repo     repository.UserRepository
	activity repository.ActivityRepository
}

func NewUserService() UserService {
	return &DefaultUserService{repo: repository.NewUserRepository()}
}
