package repository

import (
	"github.com/TechBuilder-360/business-directory-backend.git/configs"
	"github.com/TechBuilder-360/business-directory-backend.git/models"
	"go.mongodb.org/mongo-driver/mongo"
)

//go:generate mockgen -destination=../mocks/repository/mockRepo.go -package=repository github.com/TechBuilder-360/business-directory-backend.git/repository Repository
type Repository interface {
	GetClientByID(string) ( *models.Client , error)
}

type DefaultRepo struct {
	Config  *configs.Config
	Client	*mongo.Collection
 
}

func NewRepository(mdb *mongo.Database, config *configs.Config) Repository {
	client:= mdb.Collection(config.ClientCollection)
	return &DefaultRepo{
		Client: client,
		Config: config,
	}
}

