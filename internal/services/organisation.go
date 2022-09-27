package services

import (
	"errors"
	"fmt"
	"github.com/TechBuilder-360/business-directory-backend/internal/common/constant"
	"github.com/TechBuilder-360/business-directory-backend/internal/common/types"
	"github.com/TechBuilder-360/business-directory-backend/internal/common/utils"
	"github.com/TechBuilder-360/business-directory-backend/internal/configs"
	"github.com/TechBuilder-360/business-directory-backend/internal/database"
	"github.com/TechBuilder-360/business-directory-backend/internal/model"
	"github.com/TechBuilder-360/business-directory-backend/internal/repository"
	"github.com/araddon/dateparse"
	"github.com/google/uuid"
	"github.com/pariz/gountries"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// APIKeyPair Business Sec/Pub key pairs
type APIKeyPair struct {
	SecretKey string
	PublicKey string
}

//go:generate mockgen -destination=../mocks/services/organisation.go -package=services github.com/TechBuilder-360/business-directory-backend/services OrganisationService
type OrganisationService interface {
	CreateOrganisation(body *types.CreateOrganisationReq, user *model.User, logger *log.Entry) (*types.CreateOrganisationResponse, error)
	GetOrganisationByPublicKey(publicKey string) (*model.Organisation, error)
	GenerateKeyPairs() *APIKeyPair
	DeactivateOrganisation(id string, logger *log.Entry) error
	ActivateOrganisation(id string, logger *log.Entry) error
	GetSingleOrganisation(id string, logger *log.Entry) (*model.Organisation, error)
	GetAllOrganisation(page int, logger *log.Entry) (*types.DataView, error)
}

type DefaultOrganisationService struct {
	organisationRepo repository.OrganisationRepository
	branchRepo       repository.BranchRepository
	activityRepo     repository.ActivityRepository
	userRepo         repository.UserRepository
	roleRepo         repository.RoleRepository
	db               *gorm.DB
}

func NewOrganisationService() OrganisationService {
	return &DefaultOrganisationService{
		organisationRepo: repository.NewOrganisationRepository(),
		activityRepo:     repository.NewActivityRepository(),
		userRepo:         repository.NewUserRepository(),
		branchRepo:       repository.NewBranchRepository(),
		roleRepo:         repository.NewRoleRepository(),
		db:               database.ConnectDB(),
	}
}
func (o *DefaultOrganisationService) DeactivateOrganisation(id string, logger *log.Entry) error {
	organization := &model.Organisation{}
	organization.Base.ID = id
	err := o.organisationRepo.Get(organization)
	if err != nil {
		logger.Error(err)
		return errors.New(constant.InternalServerError)
	}

	organization.Active = false
	err = o.organisationRepo.Update(organization)
	if err != nil {
		logger.Error(err)
		return errors.New(constant.InternalServerError)
	}
	return nil
}
func (o *DefaultOrganisationService) ActivateOrganisation(id string, logger *log.Entry) error {
	organization := &model.Organisation{}
	organization.Base.ID = id
	err := o.organisationRepo.Get(organization)
	if err != nil {
		logger.Error(err)
		return errors.New(constant.InternalServerError)
	}

	organization.Active = true
	err = o.organisationRepo.Update(organization)
	if err != nil {
		logger.Error(err)
		return errors.New(constant.InternalServerError)
	}
	return nil
}

func (o *DefaultOrganisationService) GetSingleOrganisation(id string, logger *log.Entry) (*model.Organisation, error) {
	organization := &model.Organisation{}
	organization.Base.ID = id
	err := o.organisationRepo.Get(organization)
	if err != nil {
		logger.Error(err)
		return nil, errors.New(constant.InternalServerError)
	}
	return organization, nil
}

func (o *DefaultOrganisationService) GetAllOrganisation(page int, logger *log.Entry) (*types.DataView, error) {
	data, err := o.organisationRepo.GetAll(page)
	if err != nil {
		logger.Error(err)
		return nil, errors.New(constant.InternalServerError)
	}
	return data, nil
}
func (o *DefaultOrganisationService) CreateOrganisation(body *types.CreateOrganisationReq, user *model.User, logger *log.Entry) (*types.CreateOrganisationResponse, error) {
	uw := repository.NewGormUnitOfWork(o.db)
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

	if user.Tier < 1 {
		logger.Error("Upgrade your account to create organisation")
		return nil, errors.New("upgrade your account to create organisation")
	}

	q := gountries.New()
	country, err := q.FindCountryByAlpha(body.Country)
	if err != nil {
		return nil, errors.New("invalid country")
	}

	founding, err := dateparse.ParseLocal(body.FoundingDate)

	if err != nil {
		logger.Error(fmt.Sprintf("Founding date is not a valid date. %s %s", body.FoundingDate, err.Error()))
		return nil, errors.New(constant.InternalServerError)
	}

	keys := o.GenerateKeyPairs()

	organisationName := utils.CapitalizeFirstCharacter(body.Name)

	org, err := o.organisationRepo.GetOrganisationByName(organisationName)
	if err != nil {
		logger.Error(err.Error())
		return nil, errors.New(constant.InternalServerError)
	}
	if org != nil {
		logger.Error("organisation name already exist")
		return nil, errors.New("organisation name already exist")
	}

	organisation := &model.Organisation{
		UserID:           user.ID,
		OrganisationName: organisationName,
		Description:      body.Description,
		FoundingDate:     utils.FormatDate(founding),
		OrganisationSize: body.OrganisationSize,
		Category:         body.Category,
		Country:          country.Name.Official,
		PublicKey:        keys.PublicKey,
		SecretKey:        keys.SecretKey,
	}
	organisation.ID = utils.GenerateUUID()

	err = o.organisationRepo.WithTx(tx).Create(organisation)
	if err != nil {
		logger.Error(err.Error())
		return nil, errors.New("organisation creation failed")
	}

	branch := &model.Branch{
		OrganisationID: organisation.ID,
		BranchName:     organisation.OrganisationName,
		IsHQ:           true,
	}

	role, err := o.roleRepo.GetByName(model.OWNER)
	if err != nil {
		logger.Error(err.Error())
		return nil, errors.New("organisation creation failed")
	}

	member := &model.OrganisationMember{
		UserID:         user.ID,
		OrganizationID: organisation.ID,
		RoleID:         role.ID,
	}

	err = o.organisationRepo.WithTx(tx).AddOrganisationMember(member)
	if err != nil {
		return nil, errors.New("organisation creation failed")
	}

	er := o.branchRepo.WithTx(tx).Create(branch)
	if er != nil {
		logger.Error(err.Error())
		return nil, errors.New("organisation creation failed")
	}
	// Activity log
	activity := &model.Activity{For: user.ID, Message: fmt.Sprintf("Created an organisation %s", organisation.OrganisationName)}
	go func() {
		err = o.activityRepo.Create(activity)
		if err != nil {
			logger.Error(err.Error())
		}
	}()

	if err = uw.Commit(tx); err != nil {
		logger.Error(err.Error())
		return nil, errors.New("organisation could not be created")
	}
	response := &types.CreateOrganisationResponse{
		ID:          organisation.ID,
		Name:        organisation.OrganisationName,
		Description: organisation.Description,
		IsHQ:        branch.IsHQ,
	}
	return response, nil
}

func (o *DefaultOrganisationService) GetOrganisationByPublicKey(publicKey string) (*model.Organisation, error) {
	organisation, err := o.organisationRepo.GetByPublicKey(publicKey)
	if err != nil {
		return nil, err
	}

	return organisation, nil
}

func (o *DefaultOrganisationService) GenerateKeyPairs() *APIKeyPair {
	var (
		secretKey = uuid.NewString()
		publicKey = utils.ToMd5(uuid.NewString())
	)

	var (
		skHeader = "bd_sk_"
		pkHeader = "bd_pk_"
	)

	if configs.Instance.GetEnv() != configs.PRODUCTION {
		skHeader = "bd_sandbox_sk_"
		pkHeader = "bd_sandbox_pk_"
	}

	return &APIKeyPair{
		SecretKey: skHeader + secretKey,
		PublicKey: pkHeader + publicKey,
	}
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
