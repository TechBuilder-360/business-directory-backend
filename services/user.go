package services

import (
	"github.com/TechBuilder-360/business-directory-backend/dto"
	"github.com/TechBuilder-360/business-directory-backend/errs"
	"github.com/TechBuilder-360/business-directory-backend/models"
	"github.com/TechBuilder-360/business-directory-backend/repository"
	"github.com/TechBuilder-360/business-directory-backend/utility"
	log "github.com/Toflex/oris_log"
	"net/http"
)

func RegisterUser(request *dto.Registration, repo repository.Repository, log log.Logger) *errs.AppError {
	// Check if email address exist
	ok,err:=repo.DoesUserEmailExist(request.EmailAddress)
	if err != nil {
		log.Error("An Error occurred while checking if user email exist. %s", err.Error())
		return errs.CustomError(http.StatusFailedDependency, utility.SMMERROR, nil)
	}
	if ok {
		log.Info("Email address already exist. '%s'", request.EmailAddress)
		return errs.CustomError(http.StatusBadRequest, utility.EMAILALREADYEXIST, nil)
	}

	// Save user details
	userId, err := repo.RegisterUser(request)
	if err != nil {
		log.Error("Error occurred when saving new user. %s", err.Error())
		return errs.CustomError(http.StatusFailedDependency, utility.SMMERROR, nil)
	}

	// Activity log
	activity := &models.Activity{By: userId, Message: "Registered"}
	go func() {
		if err = repo.AddActivity(activity); err!=nil {
			log.Error("User activity failed to log")
		}
	}()

	// TODO: Send Activate email

	return nil
}