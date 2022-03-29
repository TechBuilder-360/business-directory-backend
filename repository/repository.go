package repository

import (
	"github.com/TechBuilder-360/business-directory-backend/configs"
	"github.com/TechBuilder-360/business-directory-backend/dto"
	"github.com/TechBuilder-360/business-directory-backend/errs"
	"github.com/TechBuilder-360/business-directory-backend/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

//go:generate mockgen -destination=../mocks/repository/mockRepo.go -package=repository github.com/TechBuilder-360/business-directory-backend/repository Repository
type Repository interface {
	GetClientByID(string) (*models.Client, error)
	CreateOrganisation(*dto.CreateOrgReq) (*dto.CreateOrgResponse, error)
	CreateBranch(*dto.CreateBranch) (string, error)
	GetOrganisations(page string) (*dto.DataView, error)
	GetBranches(organisation string, page string) (*dto.DataView, error)
	OrganisationExist(*dto.CreateOrgReq) bool
	BranchExist(*dto.CreateBranch) bool
	OrganisationStatus(*dto.OrganStatus) (*mongo.UpdateResult, error)
	GetSingleOrganisation(id string) (primitive.M, error)
	GetSingleBranch(id string) (primitive.M, error)

	DoesUserEmailExist(string) (bool, error)
	RegisterUser(*dto.Registration) (string, error)
	AddActivity(activity *models.Activity) error
	CreateUserToken(string) (string, error)
	IsTokenValid(request *dto.AuthRequest) (bool, *errs.AppError)
	GetUserInformation(string) (dto.UserProfile, error)
}

type DefaultRepo struct {
	Config *configs.Config
	Client,
	Organisation,
	Branch,
	User,
	Activities,
	Token *mongo.Collection
	Cli *mongo.Client
}

func NewRepository(mdb *mongo.Client, config *configs.Config) Repository {
	database := mdb.Database(config.MongoDBName)
	client := database.Collection(config.ClientCol)
	organisation := database.Collection(config.OrganCol)
	branch := database.Collection(config.BranchCol)
	user := database.Collection(config.BranchCol)
	activity := database.Collection(config.ActivityCol)
	token := database.Collection(config.TokenCol)
	return &DefaultRepo{
		Client:       client,
		Config:       config,
		Organisation: organisation,
		Branch:       branch,
		User:         user,
		Activities:   activity,
		Token: 		  token,
		Cli:          mdb,
	}
}
