package services

import (
	"fmt"
	"github.com/TechBuilder-360/business-directory-backend/dto"
	"github.com/TechBuilder-360/business-directory-backend/errs"
	"github.com/TechBuilder-360/business-directory-backend/models"
	"github.com/TechBuilder-360/business-directory-backend/repository"
	"github.com/TechBuilder-360/business-directory-backend/utility"
	log "github.com/Toflex/oris_log"
	"github.com/araddon/dateparse"
	"net/http"
)

func CreateOrganisation(request *dto.CreateOrgReq, repo repository.Repository, log log.Logger) (*dto.CreateOrgResponse, *errs.AppError) {
	_, err := dateparse.ParseLocal(request.FoundingDate)
	if err != nil {
		errMsg := fmt.Sprintf("Founding date is not a valid date. %s %s", request.FoundingDate, err.Error())
		log.Error(errMsg)
		return nil, errs.CustomError(http.StatusBadRequest, utility.BAD_REQUEST, &errMsg)
	}
	val := repo.OrganisationExist(request)
	if val == true {
		log.Error("Organisation name Already Exist")
		return nil, errs.CustomError(http.StatusBadRequest, utility.ALREADY_EXIST, nil)

	}

	// Pass request to repo
	response, err := repo.CreateOrganisation(request)
	if err != nil {
		log.Error("Error occurred while creating organisation, %s", err.Error())
		return nil, errs.UnexpectedError(utility.SMMERROR)
	}

	// TODO: Add logged in user's ID to activity log

	// Activity log
	activity := &models.Activity{For: response.OrganisationID, Message: "Created an Organisation"}
	go func() {
		if err = repo.AddActivity(activity); err!=nil {
			log.Error("User activity failed to log")
		}
	}()

	return response, nil
}

func CreateBranch(request *dto.CreateBranch, repo repository.Repository, log log.Logger) *errs.AppError {
	val := repo.BranchExist(request)
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

	return nil
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