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
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// APIKeyPair Business Sec/Pub key pairs
type APIKeyPair struct {
	SecretKey string
	PublicKey string
}

type IBusinessService interface {
	CreateBusiness(body *types.BusinessReq, user *model.User, logger *log.Entry) (*types.BusinessResponse, error)
	//GetBusinessByPublicKey(publicKey string) (*model.Business, error)
	GenerateKeyPairs() *APIKeyPair
	//ChangeBusinessStatus(business *model.Business, user *model.User, body *types.Activate, logger *log.Entry) error
	//GetSingleBusiness(id string) (*types.Business, error)
	//GetAllBusiness(query types.Query, logger *log.Entry) (*types.PaginatedResponse, error)
}

type DefaultBusinessService struct {
	businessRepo repository.BusinessRepository
	branchRepo   repository.BranchRepository
	activityRepo repository.ActivityRepository
	userRepo     repository.UserRepository
	roleRepo     repository.RoleRepository
	countryRepo  repository.CountryRepository
	db           *gorm.DB
}

func NewBusinessService() IBusinessService {
	return &DefaultBusinessService{
		businessRepo: repository.NewBusinessRepository(),
		activityRepo: repository.NewActivityRepository(),
		userRepo:     repository.NewUserRepository(),
		branchRepo:   repository.NewBranchRepository(),
		roleRepo:     repository.NewRoleRepository(),
		countryRepo:  repository.NewCountryRepository(),
		db:           database.ConnectDB(),
	}
}

//func (b *DefaultBusinessService) ChangeBusinessStatus(business *model.Business, user *model.User, body *types.Activate, logger *log.Entry) error {
//	// todo: before activating an business there are some validations that needs be done
//	if business.Active != body.Status {
//		business.Active = body.Status
//		err := o.businessRepo.Update(business)
//		if err != nil {
//			logger.Error(err)
//			return err
//		}
//
//		status := "offline"
//		if body.Status {
//			status = "online"
//		}
//
//		if configs.Instance.GetEnv() != configs.SANDBOX {
//			message := fmt.Sprintf("Your business %s is %s.", business.Name, status)
//			// Send Activate email
//			mailTemplate := &sendgrid.GeneralMailRequest{
//				ToMail:  business.EmailAddress,
//				ToName:  business.Name,
//				Subject: "VerificationType Update",
//				Message: message,
//			}
//			err = sendgrid.GeneralMail(mailTemplate)
//			if err != nil {
//				log.Error("Error occurred when sending activation email. %s", err.Error())
//			}
//		}
//
//		// Activity log
//		go func() {
//			activity := &model.Activity{For: business.ID, By: user.ID, Message: fmt.Sprintf("Change business status to '%s'", status)}
//			err = o.activityRepo.Create(activity)
//			if err != nil {
//				logger.Error(err.Error())
//			}
//		}()
//	}
//
//	return nil
//}

//func (o *DefaultbusinessService) GetSinglebusiness(id string) (*types.business, error) {
//	business, err := o.businessRepo.Get(id)
//	if err != nil {
//		return nil, err
//	}
//
//	branches := make([]types.Branch, 0)
//
//	for _, b := range business.Branch {
//		country, _ := o.countryRepo.GetCountryByID(b.CountryID)
//		branches = append(branches, types.Branch{
//			Name:        b.Name,
//			IsHQ:        b.IsHQ,
//			PhoneNumber: b.PhoneNumber,
//			Country:     country.Code,
//			ZipCode:     b.ZipCode,
//			Street:      b.Street,
//			City:        b.City,
//			State:       b.State,
//			Longitude:   b.Longitude,
//			Latitude:    b.Latitude,
//		})
//	}
//
//	response := types.business{
//		ID:                 business.ID,
//		Name:               business.Name,
//		LogoURL:            business.LogoURL,
//		Website:            business.Website,
//		businessSize:       business.businessSize,
//		Description:        business.Description,
//		RegistrationNumber: business.RegistrationNumber,
//		Rating:             business.Rating,
//		FoundingDate:       business.FoundingDate,
//		Verified:           business.Verified,
//		Branch:             branches,
//	}
//
//	return &response, nil
//}
//
//func (o *DefaultbusinessService) GetAllbusiness(query types.Query, logger *log.Entry) (*types.PaginatedResponse, error) {
//	total, err := o.businessRepo.Total(query)
//	if err != nil {
//		logger.Error(err)
//		return &types.PaginatedResponse{Data: []interface{}{}}, nil
//	}
//
//	data, err := o.businessRepo.GetAll(query)
//	if err != nil {
//		logger.Error(err)
//		return &types.PaginatedResponse{Data: []interface{}{}}, nil
//	}
//
//	businesss := make([]types.businesss, 0)
//
//	for _, business := range data {
//		// todo: set location based on nearest location to user
//		businesss = append(businesss, types.businesss{
//			ID:          business.ID,
//			Name:        business.Name,
//			LogoURL:     business.LogoURL,
//			Description: business.Description,
//			Rating:      business.Rating,
//			Verified:    business.Verified,
//		})
//	}
//
//	return &types.PaginatedResponse{
//		Page:    query.Page,
//		PerPage: query.PageSize,
//		Total:   total,
//		Data:    businesss,
//	}, nil
//
//}

func (b *DefaultBusinessService) CreateBusiness(body *types.BusinessReq, user *model.User, logger *log.Entry) (*types.BusinessResponse, error) {
	uw := repository.NewGormUnitOfWork(b.db)
	tx, err := uw.Begin()
	if err != nil {
		return nil, err
	}

	defer tx.Rollback()

	if !user.Verified {
		logger.Error("Verify your account to create business")
		return nil, errors.New("verify your account to create business")
	}

	country, err := b.countryRepo.GetCountryByCode(body.Country)
	if err != nil {
		logger.Error(err.Error())
		return nil, errors.New("country approved for business location")
	}

	founding, err := dateparse.ParseLocal(body.FoundingDate)
	if err != nil {
		logger.Error(fmt.Sprintf("Founding date is not a valid date. %s %s", body.FoundingDate, err.Error()))
		return nil, errors.New("invalid founding date")
	}

	keys := b.GenerateKeyPairs()

	businessName := utils.CapitalizeFirstCharacter(body.Name)

	org, err := b.businessRepo.GetBusinessByName(businessName)
	if err != nil {
		logger.Error(err.Error())
		return nil, errors.New("request failed, please try again")
	}

	if org != nil {
		logger.Error("business name already exist")
		return nil, errors.New("business name already registered")
	}

	business := &model.Business{
		Name:         businessName,
		Description:  body.Description,
		FoundingDate: utils.FormatDate(founding),
		BusinessSize: body.BusinessSize,
		Category:     body.Category,
		CountryID:    country.ID,
		PublicKey:    keys.PublicKey,
		SecretKey:    keys.SecretKey,
	}
	business.ID = utils.GenerateUUID()

	err = b.businessRepo.WithTx(tx).Create(business)
	if err != nil {
		logger.Error(err.Error())
		return nil, errors.New("business creation failed")
	}

	branch := &model.Branch{
		BusinessID:  business.ID,
		Name:        business.Name,
		IsHQ:        true,
		Active:      true,
		PhoneNumber: business.PhoneNumber,
		CountryID:   country.ID,
	}

	role, err := b.roleRepo.GetByName(model.OWNER)
	if err != nil {
		logger.Error(err.Error())
		return nil, errors.New("business creation failed")
	}

	member := &model.Member{
		UserId:     user.ID,
		BusinessID: business.ID,
		RoleID:     role.ID,
		Status:     constant.Active,
	}

	err = b.businessRepo.WithTx(tx).AddBusinessMember(member)
	if err != nil {
		return nil, errors.New("business creation failed")
	}

	er := b.branchRepo.WithTx(tx).Create(branch)
	if er != nil {
		logger.Error(err.Error())
		return nil, errors.New("business creation failed")
	}

	if err = uw.Commit(tx); err != nil {
		logger.Error(err.Error())
		return nil, errors.New("business could not be created")
	}
	response := &types.BusinessResponse{
		ID:          business.ID,
		Name:        business.Name,
		Description: business.Description,
		IsHQ:        branch.IsHQ,
		Branch: []types.Branch{{
			Name:        branch.Name,
			IsHQ:        branch.IsHQ,
			PhoneNumber: branch.PhoneNumber,
			Country:     country.Code,
		}},
	}
	return response, nil
}

//func (o *DefaultbusinessService) GetbusinessByPublicKey(publicKey string) (*model.business, error) {
//	business, err := o.businessRepo.GetByPublicKey(publicKey)
//	if err != nil {
//		return nil, err
//	}
//
//	return business, nil
//}

func (b *DefaultBusinessService) GenerateKeyPairs() *APIKeyPair {
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
