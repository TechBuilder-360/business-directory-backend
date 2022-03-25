package services

import (
	"github.com/TechBuilder-360/business-directory-backend/configs"
	"github.com/TechBuilder-360/business-directory-backend/repository"
)


//go:generate mockgen -destination=../mocks/service/mockService.go -package=service github.com/TechBuilder-360/business-directory-backend/services Service
type Service interface {
}


type DefaultService struct {
	repo repository.Repository
	config *configs.Config
}

func NewService(repo repository.Repository) Service {
	return DefaultService{repo: repo}
}
