package services

//go:generate mockgen -destination=../mocks/services/mockService.go -package=services github.com/TechBuilder-360/business-directory-backend/services Service
//type Service interface {
//}
//
//
//type DefaultService struct {
//	repo repository.
//	config *configs.Config
//}
//
//func NewService(repo repository.Repository) Service {
//	return DefaultService{repo: repo}
//}
