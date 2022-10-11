package services

import (
	"errors"
	"fmt"
	"github.com/TechBuilder-360/business-directory-backend/internal/common/constant"
	"github.com/TechBuilder-360/business-directory-backend/internal/common/types"
	"github.com/TechBuilder-360/business-directory-backend/internal/common/utils"
	"github.com/TechBuilder-360/business-directory-backend/internal/configs"
	"github.com/TechBuilder-360/business-directory-backend/internal/database"
	"github.com/TechBuilder-360/business-directory-backend/internal/infrastructure/sendgrid"
	"github.com/TechBuilder-360/business-directory-backend/internal/model"
	"github.com/TechBuilder-360/business-directory-backend/internal/repository"
	"github.com/araddon/dateparse"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// APIKeyPair Business Sec/Pub key pairs
type APIKeyPair struct {
	SecretKey string
	PublicKey string
}

//go:generate mockgen -destination=../mocks/services/organisation.go -package=services github.com/TechBuilder-360/business-directory-backend/services OrganisationService
type IOrganisationService interface {
	CreateOrganisation(body *types.CreateOrganisationReq, user *model.User, logger *log.Entry) (*types.CreateOrganisationResponse, error)
	GetOrganisationByPublicKey(publicKey string) (*model.Organisation, error)
	GenerateKeyPairs() *APIKeyPair
	ChangeOrganisationStatus(organisation *model.Organisation, user *model.User, body *types.Activate, logger *log.Entry) error
	GetSingleOrganisation(id string) (*types.Organisation, error)
	GetAllOrganisation(query types.Query, logger *log.Entry) (*types.PaginatedResponse, error)
}

type DefaultOrganisationService struct {
	organisationRepo repository.OrganisationRepository
	branchRepo       repository.BranchRepository
	activityRepo     repository.ActivityRepository
	userRepo         repository.UserRepository
	roleRepo         repository.RoleRepository
	countryRepo      repository.CountryRepository
	db               *gorm.DB
}

func NewOrganisationService() IOrganisationService {
	return &DefaultOrganisationService{
		organisationRepo: repository.NewOrganisationRepository(),
		activityRepo:     repository.NewActivityRepository(),
		userRepo:         repository.NewUserRepository(),
		branchRepo:       repository.NewBranchRepository(),
		roleRepo:         repository.NewRoleRepository(),
		countryRepo:      repository.NewCountryRepository(),
		db:               database.ConnectDB(),
	}
}

func (o *DefaultOrganisationService) ChangeOrganisationStatus(organisation *model.Organisation, user *model.User, body *types.Activate, logger *log.Entry) error {
	// todo: before activating an organisation there are some validations that needs be done
	if organisation.Active != body.Status {
		organisation.Active = body.Status
		err := o.organisationRepo.Update(organisation)
		if err != nil {
			logger.Error(err)
			return err
		}

		status := "offline"
		if body.Status {
			status = "online"
		}

		if configs.Instance.GetEnv() != configs.SANDBOX {
			message := fmt.Sprintf("Your organisation %s is %s.", organisation.Name, status)
			// Send Activate email
			mailTemplate := &sendgrid.GeneralMailRequest{
				ToMail:  organisation.EmailAddress,
				ToName:  organisation.Name,
				Subject: "Status Update",
				Message: message,
			}
			err = sendgrid.GeneralMail(mailTemplate)
			if err != nil {
				log.Error("Error occurred when sending activation email. %s", err.Error())
			}
		}

		// Activity log
		go func() {
			activity := &model.Activity{For: organisation.ID, By: user.ID, Message: fmt.Sprintf("Change organisation status to '%s'", status)}
			err = o.activityRepo.Create(activity)
			if err != nil {
				logger.Error(err.Error())
			}
		}()
	}

	return nil
}

func (o *DefaultOrganisationService) GetSingleOrganisation(id string) (*types.Organisation, error) {
	organisation, err := o.organisationRepo.Get(id)
	if err != nil {
		return nil, err
	}

	branches := make([]types.Branch, 0)

	for _, b := range organisation.Branch {
		country, _ := o.countryRepo.GetCountryByID(b.CountryID)
		branches = append(branches, types.Branch{
			Name:        b.Name,
			IsHQ:        b.IsHQ,
			PhoneNumber: b.PhoneNumber,
			Country:     country.Code,
			ZipCode:     b.ZipCode,
			Street:      b.Street,
			City:        b.City,
			State:       b.State,
			Longitude:   b.Longitude,
			Latitude:    b.Latitude,
		})
	}

	response := types.Organisation{
		ID:                 organisation.ID,
		Name:               organisation.Name,
		LogoURL:            organisation.LogoURL,
		Website:            organisation.Website,
		OrganisationSize:   organisation.OrganisationSize,
		Description:        organisation.Description,
		RegistrationNumber: organisation.RegistrationNumber,
		Rating:             organisation.Rating,
		FoundingDate:       organisation.FoundingDate,
		Verified:           organisation.Verified,
		Branch:             branches,
	}

	return &response, nil
}

func (o *DefaultOrganisationService) GetAllOrganisation(query types.Query, logger *log.Entry) (*types.PaginatedResponse, error) {
	total, err := o.organisationRepo.Total(query)
	if err != nil {
		logger.Error(err)
		return &types.PaginatedResponse{Data: []interface{}{}}, nil
	}

	data, err := o.organisationRepo.GetAll(query)
	if err != nil {
		logger.Error(err)
		return &types.PaginatedResponse{Data: []interface{}{}}, nil
	}

	organisations := make([]types.Organisations, 0)

	for _, organisation := range data {
		// todo: set location based on nearest location to user
		organisations = append(organisations, types.Organisations{
			ID:          organisation.ID,
			Name:        organisation.Name,
			LogoURL:     organisation.LogoURL,
			Description: organisation.Description,
			Rating:      organisation.Rating,
			Verified:    organisation.Verified,
		})
	}

	return &types.PaginatedResponse{
		Page:    query.Page,
		PerPage: query.PageSize,
		Total:   total,
		Data:    organisations,
	}, nil

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

	country, err := o.countryRepo.GetCountryByCode(body.Country)
	if err != nil {
		logger.Error(err.Error())
		return nil, errors.New("country not found")
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
		Name:             organisationName,
		Description:      body.Description,
		FoundingDate:     utils.FormatDate(founding),
		OrganisationSize: body.OrganisationSize,
		Category:         body.Category,
		CountryID:        country.ID,
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
		Name:           organisation.Name,
		IsHQ:           true,
		Active:         true,
		PhoneNumber:    organisation.PhoneNumber,
		CountryID:      country.ID,
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
	activity := &model.Activity{For: user.ID, Message: fmt.Sprintf("Created an organisation %s", organisation.Name)}
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
		Name:        organisation.Name,
		Description: organisation.Description,
		IsHQ:        branch.IsHQ,
		Branch: []types.Branch{{
			Name:        branch.Name,
			IsHQ:        branch.IsHQ,
			PhoneNumber: branch.PhoneNumber,
			Country:     country.Code,
			ZipCode:     branch.ZipCode,
			Street:      branch.Street,
			City:        branch.City,
			State:       branch.State,
		}},
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
