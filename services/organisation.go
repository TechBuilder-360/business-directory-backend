package services

import (
	"errors"
	"fmt"
	"gorm.io/gorm"

	"github.com/TechBuilder-360/business-directory-backend/dto"
	"github.com/TechBuilder-360/business-directory-backend/models"
	"github.com/TechBuilder-360/business-directory-backend/repository"
	"github.com/TechBuilder-360/business-directory-backend/utility"

	log "github.com/Toflex/oris_log"
	"github.com/araddon/dateparse"
	"net/http"
)


//go:generate mockgen -destination=../mocks/services/organisation.go -package=services github.com/TechBuilder-360/business-directory-backend/services OrganisationService
type OrganisationService interface {
	CreateOrganisation(body *dto.CreateOrgReq, user *models.UserProfile, logger log.Logger) (*dto.Organisation, error)
	Create() error
}


type DefaultOrganisationService struct {
	repo repository.OrganisationRepository
	activity repository.ActivityRepository
	db *gorm.DB
}


func NewOrganisationService(repo repository.OrganisationRepository, activity repository.ActivityRepository) OrganisationService {
	return &DefaultOrganisationService{repo: repo, activity: activity}
}

func (d *DefaultOrganisationService) Create() error {




}

func (d *DefaultOrganisationService) CreateOrganisation(body *dto.CreateOrgReq, user *models.UserProfile, log log.Logger) (*dto.Organisation, error) {
	uw := repository.NewGormUnitOfWork(d.db)
	tx, err := uw.Begin()

	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	if err != nil {
		return nil, err
	}

	founding, err := dateparse.ParseLocal(body.FoundingDate)
	if err != nil {
		errMsg := fmt.Sprintf("Founding date is not a valid date. %s %s", body.FoundingDate, err.Error())
		return nil, errors.New(errMsg)
	}

	body.OrganisationName = utility.CapitalizeFirstCharacter(body.OrganisationName)

	organisation := &models.Organisation{}
	filter := map[string]interface{}{"organisation_name": body.OrganisationName}
	err = d.repo.WithTx(tx).FindOrganisation(filter, organisation)
	if err != nil {
		return nil, err
	}

	if organisation.ID == "" {
		return nil, errors.New("organisation name already exist")
	}

	organisation.CreatorID = user.ID
	organisation.OrganisationName = body.OrganisationName
	organisation.Description = utility.FormatDate(founding)
	organisation.OrganisationSize = body.OrganisationSize
	organisation.Description = body.Description

	err = d.repo.WithTx(tx).Create(organisation)
	if err != nil {
		return nil, err
	}

	// Activity log
	activity := &models.Activity{For: user.ID, Message: fmt.Sprintf("Created an organisation %s", organisation.OrganisationName)}
	go d.activity.Create(activity)

	response := &dto.Organisation{}
	response.ToDTO(organisation)

	if err = uw.Commit(tx); err!= nil{
		return nil, err
	}

	return response, nil
}

func (d *DefaultOrganisationService) CreateBranch(body *dto.CreateBranch, organisation *models.Organisation, log log.Logger) (*dto.CreateBranch, error) {
	uw := repository.NewGormUnitOfWork(d.db)
	tx, err := uw.Begin()

	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	if err != nil {
		return nil, err
	}

	branch := &models.Branch{}
	branch.OrganisationID = body.OrganisationID
	filter := map[string]interface{}{"branch_name": body.BranchName, "organisation_id": body.OrganisationID}
	err := d.repo.WithTx(tx).FindBranch(filter, )
	if val == true {
		log.Error("Branch name Already Exist")
		return errs.CustomError(http.StatusNotAcceptable, utility.ALREADY_EXIST, nil)
	}

	_, err := repo.CreateBranch(request)
	if err != nil {
		log.Error("Error occurred while creating branch, %s", err.Error())
		return errs.UnexpectedError(utility.SMMERROR)
	}

	// TODO: Add logged in user's ID to activity log
	// Activity log
	activity := &models.Activity{By: "", For: request.OrganisationID, Message: fmt.Sprintf("Added a branch '%s'", request.BranchName)}
	go func() {
		if err = repo.AddActivity(activity); err != nil {
			log.Error("User activity failed to log")
		}
	}()

	if err = uw.Commit(tx); err!= nil{
		return nil, err
	}

	return nil, nil
}

func GetOrganisations(page int, repo repository.Repository, log log.Logger) (*dto.DataView, *errs.AppError) {
	organisations, err := repo.GetOrganisations(page)
	if err != nil {
		log.Error("Error occured while getting list of organisations, %s", err.Error())
		return nil, errs.UnexpectedError(utility.SMMERROR)
	}

	return organisations, nil
}

func GetSingleOrganisation(orgId string, repo repository.Repository, log log.Logger) (interface{}, *errs.AppError) {
	data, err := repo.GetSingleOrganisation(orgId)
	if err != nil {
		log.Error("Error occurred while getting organisation, %s", err.Error())
		return nil, errs.UnexpectedError(utility.SMMERROR)
	}

	return data, nil
}

func GetBranches(orgId, page string, repo repository.Repository, log log.Logger) (interface{}, *errs.AppError) {
	data, err := repo.GetBranches(orgId, page)
	if err != nil {
		log.Error("Error occurred while getting organisation branches, %s", err.Error())
		return nil, errs.UnexpectedError(utility.SMMERROR)
	}

	return data, nil
}

func GetSingleBranch(branchID string, repo repository.Repository, log log.Logger) (interface{}, *errs.AppError) {
	branch, err := repo.GetSingleBranch(branchID)
	if err != nil {
		log.Error("Error occured while getting branch, %s", err.Error())
		return nil, errs.NewValidationError(utility.SMMERROR)
	}

	return branch, nil
}

func OrganisationStatus(Organs *dto.OrganStatus, repo repository.Repository, log log.Logger) (interface{}, *errs.AppError) {
	_, err := repo.OrganisationStatus(Organs)

	if err != nil {
		log.Error("Error occurred while deactivating or activating organisation, %s", err.Error())
		return nil, errs.UnexpectedError(utility.SMMERROR)
	}

	// TODO: Add logged in user's ID to activity log
	// Activity log
	activity := &models.Activity{By: "", For: Organs.OrganisationID, Message: fmt.Sprintf("Changed organisation active status to %t", Organs.Active)}
	go func() {
		if err = repo.AddActivity(activity); err!=nil {
			log.Error("User activity failed to log")
		}
	}()

	return nil, nil
}