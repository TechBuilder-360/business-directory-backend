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
	CreateOrganisation(*dto.CreateOrganisation)(string,error)
	CreateBranch(*dto.CreateBranch)(string,error)
	GetOrganisation() (*mongo.Cursor, error)
	GetBranch(organisation string) (*mongo.Cursor, error)
	AlreadyOrganisation(*dto.CreateOrganisation) (bool) 
	AlreadyBranch(*dto.CreateBranch) (bool) 
	DeactivateOrganisation(id string,value bool) (*mongo.UpdateResult, error)

}

type DefaultRepo struct {
	Config  *configs.Config
	Client	*mongo.Collection
	Organisation *mongo.Collection
	Branch *mongo.Collection
 
}

func NewRepository(mdb *mongo.Database, config *configs.Config) Repository {
	client:= mdb.Collection("Clients")
	organisation:= mdb.Collection("Organisations")
	branch:= mdb.Collection("Branch")
	return &DefaultRepo{
		Client: client,
		Config: config,
		Organisation: organisation,
		Branch:branch,
	}
}




