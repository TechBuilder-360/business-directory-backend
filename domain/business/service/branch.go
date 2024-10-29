package service

import (
	"github.com/TechBuilder-360/business-directory-backend/domain/business/model"
	"github.com/TechBuilder-360/business-directory-backend/domain/business/repository"
	cr "github.com/TechBuilder-360/business-directory-backend/domain/country/repository"
	"github.com/TechBuilder-360/business-directory-backend/internal/common/types"
	"github.com/TechBuilder-360/business-directory-backend/internal/database"
	"github.com/TechBuilder-360/business-directory-backend/pkg/log"
	"gorm.io/gorm"
)

type BranchService interface {
	Create(branch *model.Branch) error
	GetBusinessBranches(Business *model.Business, logger log.Entry) ([]types.Branch, error)
	Update(branch *model.Branch) error
	UpdateBranch(Business *model.Business, body types.CreateBranchRequest)
	CreateBranch(Business *model.Business, body types.CreateBranchRequest)
	Activate(Business *model.Business, id string, body types.Activate)
}

type DefaultBranchService struct {
	branchRepo  repository.BranchRepository
	countryRepo cr.CountryRepository
	db          *gorm.DB
}

func NewBranchService() BranchService {
	return &DefaultBranchService{
		branchRepo:  repository.NewBranchRepository(),
		countryRepo: cr.NewCountryRepository(),
		db:          database.ConnectDB(),
	}
}

func (b *DefaultBranchService) Update(branch *model.Branch) error {
	return b.branchRepo.Update(branch)
}

func (b *DefaultBranchService) UpdateBranch(Business *model.Business, body types.CreateBranchRequest) {
	//TODO implement me
	panic("implement me")
}

func (b *DefaultBranchService) CreateBranch(Business *model.Business, body types.CreateBranchRequest) {
	//TODO implement me
	panic("implement me")
}

func (b *DefaultBranchService) Activate(Business *model.Business, id string, body types.Activate) {
	//TODO implement me
	panic("implement me")
}

func (b *DefaultBranchService) Create(branch *model.Branch) error {
	return b.branchRepo.Create(branch)
}

func (b *DefaultBranchService) GetBusinessBranches(Business *model.Business, logger log.Entry) ([]types.Branch, error) {
	response := make([]types.Branch, 0)
	branches, err := b.branchRepo.GetByBusiness(Business.Name)
	if err != nil {
		logger.Error(err.Error())
	}

	for _, branch := range branches {
		country, _ := b.countryRepo.GetCountryByID(branch.CountryID)
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
