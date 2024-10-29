package service

import (
	"errors"
	"github.com/TechBuilder-360/business-directory-backend/domain/user/model"
	ur "github.com/TechBuilder-360/business-directory-backend/domain/user/repository"
	"github.com/TechBuilder-360/business-directory-backend/internal/common/constant"
	"github.com/TechBuilder-360/business-directory-backend/internal/common/types"
	"github.com/TechBuilder-360/business-directory-backend/internal/configs"
	"github.com/TechBuilder-360/business-directory-backend/internal/repository"
	"github.com/TechBuilder-360/business-directory-backend/pkg/log"
)

//go:generate mockgen -destination=../mocks/services/mockService.go -package=services github.com/TechBuilder-360/business-directory-backend/services UserService
type UserService interface {
	UpgradeStatus(body *types.UpgradeUserTierRequest, user *model.User, logger log.Entry) error
	Update(user *model.User) error
}

type DefaultUserService struct {
	userRepo ur.UserRepository
	activity repository.ActivityRepository
}

func (r *DefaultUserService) Update(user *model.User) error {
	return r.userRepo.Update(user)
}

func NewUserService() UserService {
	return &DefaultUserService{userRepo: ur.NewUserRepository()}
}

func (r *DefaultUserService) UpgradeStatus(body *types.UpgradeUserTierRequest, user *model.User, logger log.Entry) error {
	if user.Disabled {
		return nil
	}

	//todo: Request body is yet to be known
	if configs.Instance.IsProduction() {
		// Todo: Identity needs to be verified
		user.Disabled = true

		err := r.Update(user)
		if err != nil {
			logger.Error(err.Error())
			return errors.New(constant.InternalServerError)
		}
	}

	return nil
}
