package services

import (
	"github.com/TechBuilder-360/business-directory-backend/internal/common/types"
	"github.com/TechBuilder-360/business-directory-backend/internal/database"
	"github.com/TechBuilder-360/business-directory-backend/internal/model"
	"github.com/TechBuilder-360/business-directory-backend/internal/repository"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

//go:generate mockgen -destination=../mocks/services/branch.go -package=services github.com/TechBuilder-360/business-directory-backend/services BranchService
type BranchService interface {
	Create(branch *model.Branch) error
	GetOrganisationBranches(organisation *model.Organisation, logger *log.Entry) ([]types.Branch, error)
	Update(branch *model.Branch) error
	UpdateBranch(organisation *model.Organisation, body types.CreateBranchRequest)
	CreateBranch(organisation *model.Organisation, body types.CreateBranchRequest)
	Activate(organisation *model.Organisation, id string, body types.Activate)
}

type DefaultBranchService struct {
	branchRepo  repository.BranchRepository
	countryRepo repository.CountryRepository
	db          *gorm.DB
}

func NewBranchService() BranchService {
	return &DefaultBranchService{
		branchRepo:  repository.NewBranchRepository(),
		countryRepo: repository.NewCountryRepository(),
		db:          database.ConnectDB(),
	}
}

func (o *DefaultBranchService) Update(branch *model.Branch) error {
	return o.branchRepo.Update(branch)
}

func (o *DefaultBranchService) UpdateBranch(organisation *model.Organisation, body types.CreateBranchRequest) {
	//TODO implement me
	panic("implement me")
}

func (o *DefaultBranchService) CreateBranch(organisation *model.Organisation, body types.CreateBranchRequest) {
	//TODO implement me
	panic("implement me")
}

func (o *DefaultBranchService) Activate(organisation *model.Organisation, id string, body types.Activate) {
	//TODO implement me
	panic("implement me")
}

func (o *DefaultBranchService) Create(branch *model.Branch) error {
	return o.branchRepo.Create(branch)
}

func (o *DefaultBranchService) GetOrganisationBranches(organisation *model.Organisation, logger *log.Entry) ([]types.Branch, error) {
	response := make([]types.Branch, 0)
	branches, err := o.branchRepo.GetByOrganisation(organisation.Name)
	if err != nil {
		logger.Error(err.Error())
	}

	for _, branch := range branches {
		country, _ := o.countryRepo.GetCountryByID(branch.CountryID)
		response = append(response, types.Branch{
			Name:        branch.Name,
			IsHQ:        branch.IsHQ,
			PhoneNumber: branch.PhoneNumber,
			Country:     country.Code,
			ZipCode:     branch.ZipCode,
			Street:      branch.Street,
			City:        branch.City,
			State:       branch.State,
			Longitude:   branch.Longitude,
			Latitude:    branch.Latitude,
		})
	}

	return response, nil
}
