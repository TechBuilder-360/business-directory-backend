package services

import (
	"github.com/TechBuilder-360/business-directory-backend/internal/repository"
	"gorm.io/gorm"
)

//go:generate mockgen -destination=../mocks/services/organisation.go -package=services github.com/TechBuilder-360/business-directory-backend/services OrganisationService
type OrganisationService interface {
	//CreateOrganisation(body *types.CreateOrgReq, user *model.User, logger *log.Entry) (*types.Organisation, error)
}

type DefaultOrganisationService struct {
	repo     repository.OrganisationRepository
	activity repository.ActivityRepository
	db       *gorm.DB
}

func NewOrganisationService() OrganisationService {
	return &DefaultOrganisationService{repo: repository.NewOrganisationRepository()}
}

//func (d *DefaultOrganisationService) CreateOrganisation(body *types.CreateOrgReq, user *model.User, log *log.Entry) (*types.Organisation, error) {
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
//	founding, err := dateparse.ParseLocal(body.FoundingDate)
//	if err != nil {
//		errMsg := fmt.Sprintf("Founding date is not a valid date. %s %s", body.FoundingDate, err.Error())
//		return nil, errors.New(errMsg)
//	}
//
//	body.OrganisationName = utils.CapitalizeFirstCharacter(body.OrganisationName)
//
//	organisation := &model.Organisation{}
//	filter := map[string]interface{}{"organisation_name": body.OrganisationName}
//	err = d.repo.WithTx(tx).Find(filter, organisation)
//	if err != nil {
//		return nil, err
//	}
//
//	if organisation.ID != "" {
//		return nil, errors.New("organisation name already exist")
//	}
//
//	organisation.UserID = user.ID
//	organisation.OrganisationName = body.OrganisationName
//	organisation.Description = utils.FormatDate(founding)
//	organisation.OrganisationSize = &body.OrganisationSize
//	organisation.Description = body.Description
//
//	err = d.repo.WithTx(tx).Create(organisation)
//	if err != nil {
//		return nil, errors.New("organisation could not be created")
//	}
//
//	// Activity log
//	activity := &model.Activity{For: user.ID, Message: fmt.Sprintf("Created an organisation %s", organisation.OrganisationName)}
//	go func() {
//		err := d.activity.Create(activity)
//		if err != nil {
//
//		}
//	}()
//
//	response := &types.Organisation{
//		OrganisationID:     organisation.ID,
//		OrganisationName:   organisation.OrganisationName,
//		LogoURL:            *organisation.LogoURL,
//		Website:            *organisation.Website,
//		OrganisationSize:   *organisation.OrganisationSize,
//		Description:        organisation.Description,
//		RegistrationNumber: *organisation.RegistrationNumber,
//		FoundingDate:       organisation.FoundingDate,
//	}
//
//	if err = uw.Commit(tx); err != nil {
//		return nil, errors.New("organisation could not be created")
//	}
//
//	return response, nil
//}

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
