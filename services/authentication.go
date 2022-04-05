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

// Login
// Handles authentication logic
func Login(requestData *dto.AuthRequest, repo repository.Repository, service JWTService, log log.Logger) (*dto.JWTResponse, *errs.AppError) {
	response := &dto.JWTResponse{}
	// Validate user token
	isTokenValid, err := repo.IsTokenValid(requestData)
	if err != nil {
		log.Error("An Error occurred when validating login token. %s", err.AsMessage())
		return nil, err
	}
	if isTokenValid == false {
		log.Error("Invalid login token. %s", err.AsMessage())
		return nil, errs.CustomError(http.StatusUnauthorized, utility.AUTHERROR004, nil)
	}

	profile, profileErr := repo.GetUserInformation(requestData.Email)
	if profileErr != nil {
		log.Error("An error occurred when fetching user profile. %s", profileErr.Error())
		return nil, errs.NotFoundError(utility.USERNOTFOUND)
	}

	// Generate JWT for user
	token, Tokenerr := service.GenerateToken(profile.ID)
	if Tokenerr != nil {
		log.Error("An error occurred when generating jwt token. %s", Tokenerr.Error())
		return nil, errs.NotFoundError(utility.SMMERROR)
	}

	response.Profile = profile
	response.AccessToken = token

	// Activity log
	activity := &models.Activity{By: response.Profile.ID, Message: "Successful login"}
	go func() {
		if err := repo.AddActivity(activity); err!=nil {
			log.Error("User activity failed to log")
		}
	}()

	return response, nil
}

func AuthEmail(requestData *dto.EmailRequest, repo repository.Repository, log log.Logger) *errs.AppError {
	// Check if email address exist
	ok,err:= repo.DoesUserEmailExist(requestData.EmailAddress)
	if err != nil {
		log.Error("An Error occurred while checking if user email exist. %s", err.Error())
		return errs.UnexpectedError(utility.SMMERROR)
	}
	if !ok {
		log.Info("Email address does not exist. '%s'", requestData.EmailAddress)
		return errs.NotFoundError(utility.EMAILDOESNOTEXIST)
	}

	token, err := repo.CreateUserToken(requestData.EmailAddress)
	if err != nil {
		log.Error("Error occurred when saving sign-in token. %s", err.Error())
		return errs.UnexpectedError(utility.SMMERROR)
	}

	// TODO: Send Token to user email
	token = token

	// Activity log
	activity := &models.Activity{Message: "Requested for sign in token"}
	go func() {
		if err := repo.AddActivity(activity); err!=nil {
			log.Error("User activity failed to log")
		}
	}()

	return nil
}
