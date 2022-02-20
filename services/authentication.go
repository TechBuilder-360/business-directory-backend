package services

import "github.com/TechBuilder-360/business-directory-backend.git/repository"


//go:generate mockgen -destination=../mocks/service/mockAuthService.go -package=service github.com/TechBuilder-360/business-directory-backend.git/services AuthService
type AuthService interface {
	LoginUser(string, string) bool
}

type loginInformation struct {
	repo repository.Repository
}

func DefaultAuth(repo repository.Repository) AuthService {
	return &loginInformation{
		repo: repo,
	}
}

func (info *loginInformation) LoginUser(email string, token string) bool {
	return true
}


