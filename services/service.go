package services

import "github.com/TechBuilder-360/business-directory-backend.git/repository"


//go:generate mockgen -destination=../mocks/service/mockService.go -package=service github.com/TechBuilder-360/business-directory-backend.git/services Service
type Service interface {
	GetAuthor() (Person)
}


type DefaultService struct {
	repo *repository.Repository
}

func NewService(repo *repository.Repository) DefaultService {
	return DefaultService{repo: repo}
}

type Person struct {
	Name string
	ID   int
}

func (s DefaultService) GetAuthor() (Person) {

	person:=Person{
		Name:"Adegunwa Toluwalope",
		ID: 120,
	}
	return person
}