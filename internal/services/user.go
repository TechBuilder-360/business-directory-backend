package services

import (
	"errors"
	"github.com/TechBuilder-360/business-directory-backend/internal/common/constant"
	"github.com/TechBuilder-360/business-directory-backend/internal/common/types"
	"github.com/TechBuilder-360/business-directory-backend/internal/infrastructure/cloudinary"
	"github.com/TechBuilder-360/business-directory-backend/internal/model"
	"github.com/TechBuilder-360/business-directory-backend/internal/repository"
	log "github.com/sirupsen/logrus"
)

//go:generate mockgen -destination=../mocks/services/mockService.go -package=services github.com/TechBuilder-360/business-directory-backend/services UserService
type UserService interface {
	UpgradeTierOne(body *types.UpgradeUserTierRequest, user *model.User, logger *log.Entry) error
	Update(user *model.User) error
}

type DefaultUserService struct {
	userRepo repository.UserRepository
	activity repository.ActivityRepository
}

func (r *DefaultUserService) Update(user *model.User) error {
	return r.userRepo.Update(user)
}

func NewUserService() UserService {
	return &DefaultUserService{userRepo: repository.NewUserRepository()}
}

func (r *DefaultUserService) UpgradeTierOne(body *types.UpgradeUserTierRequest, user *model.User, logger *log.Entry) error {
	url, err := cloudinary.ImageUpload(body.IdentityImage)
	if err != nil {
		logger.Error(err)
		return errors.New(constant.InternalServerError)
	}

	// Todo: Identity needs to be verified
	user.IdentityName = body.IdentityName
	user.IdentityNumber = body.IdentityNumber
	user.IdentityImage = url
	user.Tier = uint8(1)

	err = r.Update(user)
	if err != nil {
		logger.Error(err)
		return errors.New(constant.InternalServerError)
	}
	return nil
}
