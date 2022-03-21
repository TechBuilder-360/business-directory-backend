package repository

import (
	"github.com/TechBuilder-360/business-directory-backend/configs"
	"github.com/TechBuilder-360/business-directory-backend/dto"
	"github.com/TechBuilder-360/business-directory-backend/models"
	"go.mongodb.org/mongo-driver/mongo"
)

//go:generate mockgen -destination=../mocks/repository/mockRepo.go -package=repository github.com/TechBuilder-360/business-directory-backend/repository Repository
type Repository interface {
	GetClientByID(string)( *models.Client , error)
	CreateOrganisation(req *dto.CreateOrgReq)(*dto.CreateOrgResponse,error)
}

type DefaultRepo struct {
	Config  *configs.Config
	Client	*mongo.Collection
	Organisation *mongo.Collection
 
}

func NewRepository(mdb *mongo.Database, config *configs.Config) Repository {
	client:= mdb.Collection(config.ClientCollection)
	Organisation:= mdb.Collection(config.OrganisationCollection)
	return &DefaultRepo{
		Client: client,
		Config: config,
		Organisation: Organisation,
	}
}




