package services

import (
	"errors"
	"fmt"
	"github.com/TechBuilder-360/business-directory-backend/internal/common/constant"
	"github.com/TechBuilder-360/business-directory-backend/internal/common/types"
	"github.com/TechBuilder-360/business-directory-backend/internal/common/utils"
	"github.com/TechBuilder-360/business-directory-backend/internal/database"
	"github.com/TechBuilder-360/business-directory-backend/internal/model"
	"github.com/TechBuilder-360/business-directory-backend/internal/repository"
	"github.com/araddon/dateparse"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

//go:generate mockgen -destination=../mocks/services/organisation.go -package=services github.com/TechBuilder-360/business-directory-backend/services OrganisationService
type OrganisationService interface {
	CreateOrganisation(body *types.CreateOrgReq, logger *log.Entry) (*types.CreateOrgResponse, error)
}

type DefaultOrganisationService struct {
	repo       repository.OrganisationRepository
	branchRepo repository.BranchRepository
	activity   repository.ActivityRepository
	db         *gorm.DB
	userRepo   repository.UserRepository
}

func NewOrganisationService() OrganisationService {
	return &DefaultOrganisationService{repo: repository.NewOrganisationRepository(),
		activity:   repository.NewActivityRepository(),
		db:         database.ConnectDB(),
		branchRepo: repository.NewBranchRepository(),
		userRepo:   repository.NewUserRepository(),
	}
}

func (d *DefaultOrganisationService) CreateOrganisation(body *types.CreateOrgReq, logger *log.Entry) (*types.CreateOrgResponse, error) {
	uw := repository.NewGormUnitOfWork(d.db)
	tx, err := uw.Begin()

	defer func() {
		if err != nil {
			logger.Error(err.Error())
			tx.Rollback()
		}
	}()

	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	user, _ := d.userRepo.GetUserByID(body.UserID)
	if user.Tier != 1 {
		logger.Error("Upgrade your acount to create Organisation")
		return nil, errors.New("Upgrade your acount to create Organisation")
	}
	founding, err := dateparse.ParseLocal(body.FoundingDate)

	if err != nil {
		logger.Error(fmt.Sprintf("Founding date is not a valid date. %s %s", body.FoundingDate, err.Error()))
		return nil, errors.New(constant.InternalServerError)
	}

	organisation := &model.Organisation{}
	body.OrganisationName = utils.CapitalizeFirstCharacter(body.OrganisationName)

	ok, _ := d.repo.GetOrganisationByName(body.OrganisationName)
	//if err != nil {
	//	logger.Error(err.Error())
	//	return errors.New(constant.InternalServerError)
	//}
	if ok {
		logger.Error(err.Error())
		return nil, errors.New("organisation name already exist")
	}

	organisation.UserID = user.ID
	organisation.OrganisationName = body.OrganisationName
	organisation.Description = utils.FormatDate(founding)
	organisation.OrganisationSize = body.OrganisationSize

	err = d.repo.WithTx(tx).Create(organisation)
	if err != nil {
		logger.Error(err.Error())
		return nil, errors.New(constant.InternalServerError)
	}
	branch := &model.Branch{}

	branch.OrganisationID = organisation.Base.ID
	branch.BranchName = organisation.OrganisationName
	branch.IsHQ = true
	er := d.branchRepo.WithTx(tx).Create(branch)

	if er != nil {
		logger.Error(err.Error())
		return nil, errors.New(constant.InternalServerError)
	}
	// Activity log
	activity := &model.Activity{For: user.ID, Message: fmt.Sprintf("Created an organisation %s", organisation.OrganisationName)}
	go func() {
		err := d.activity.Create(activity)
		if err != nil {
			logger.Error(err.Error())

		}
	}()
	if err = uw.Commit(tx); err != nil {
		logger.Error(err.Error())
		return nil, errors.New("organisation could not be created")
	}
	response := &types.CreateOrgResponse{
		ID:               organisation.Base.ID,
		OrganisationName: organisation.OrganisationName,
		IsHQ:             branch.IsHQ,
	}
	return response, nil
}

//func (d *DefaultOrganisationService) CreateBranch(body *types.CreateBranch, organisation *model.Organisation, log log.Logger) (*types.CreateBranch, error) {
//	uw := repository.NewGormUnitOfWork(d.db)
//	tx, err := uw.Begin()
//
//	defer func() {
//		if err != nil {
//			tx.Rollback()
//		}
//	}()
//
//	if err != nil {
//		return nil, err
//	}
//
//	branch := &model.Branch{}
//	branch.OrganisationID = organisation.ID
//	//filter := map[string]interface{}{"branch_name": body.BranchName, "organisation_id": body.OrganisationID}
//	//err := d.repo.WithTx(tx).FindBranch(filter)
//	//if val == true {
//	//	log.Error("Branch name Already Exist")
//	//	return errs.CustomError(http.StatusNotAcceptable, utils.ALREADY_EXIST, nil)
//	//}
//	//
//	//_, err := repo.CreateBranch(request)
//	//if err != nil {
//	//	log.Error("Error occurred while creating branch, %s", err.Error())
//	//	return errs.UnexpectedError(utils.SMMERROR)
//	//}
//	//
//	//// TODO: Add logged in user's ID to activity log
//	//// Activity log
//	//activity := &model.Activity{By: "", For: request.OrganisationID, Message: fmt.Sprintf("Added a branch '%s'", request.BranchName)}
//	//go func() {
//	//	if err = repo.AddActivity(activity); err != nil {
//	//		log.Error("User activity failed to log")
//	//	}
//	//}()
//
//	if err = uw.Commit(tx); err != nil {
//		return nil, err
//	}
//
//	return nil, nil
//}
//
//func GetOrganisations(page int, repo repository.Repository, log log.Logger) (*types.DataView, error) {
//	organisations, err := repo.GetOrganisations(page)
//	if err != nil {
//		log.Error("Error occured while getting list of organisations, %s", err.Error())
//		return nil, err
//	}
//
//	return organisations, nil
//}
//
//func GetSingleOrganisation(orgId string, repo repository.Repository, log log.Logger) (interface{}, error) {
//	data, err := repo.GetSingleOrganisation(orgId)
//	if err != nil {
//		log.Error("Error occurred while getting organisation, %s", err.Error())
//		return nil, err
//	}
//
//	return data, nil
//}
//
//func GetBranches(orgId, page string, repo repository.Repository, log log.Logger) (interface{}, error) {
//	data, err := repo.GetBranches(orgId, page)
//	if err != nil {
//		log.Error("Error occurred while getting organisation branches, %s", err.Error())
//		return nil, err
//	}
//
//	return data, nil
//}
//
//func GetSingleBranch(branchID string, repo repository.Repository, log log.Logger) (interface{}, error) {
//	branch, err := repo.GetSingleBranch(branchID)
//	if err != nil {
//		log.Error("Error occured while getting branch, %s", err.Error())
//		return nil, err
//	}
//
//	return branch, nil
//}
//
//func OrganisationStatus(Organs *types.OrganStatus, repo repository.Repository, log log.Logger) (interface{}, error) {
//	_, err := repo.OrganisationStatus(Organs)
//
//	if err != nil {
//		log.Error("Error occurred while deactivating or activating organisation, %s", err.Error())
//		return nil, err
//	}
//
//	// TODO: Add logged in user's ID to activity log
//	// Activity log
//	activity := &model.Activity{By: "", For: Organs.OrganisationID, Message: fmt.Sprintf("Changed organisation active status to %t", Organs.Active)}
//	go func() {
//		if err = repo.AddActivity(activity); err != nil {
//			log.Error("User activity failed to log")
//		}
//	}()
//
//	return nil, nil
//}
