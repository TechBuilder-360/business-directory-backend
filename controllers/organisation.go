package controllers

import (
	"context"
	"encoding/json"
	"github.com/TechBuilder-360/business-directory-backend/dto"
	"github.com/TechBuilder-360/business-directory-backend/utility"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
)

func (c *NewController) CreateOrganisation(w http.ResponseWriter, r *http.Request) {
	log := c.Logger.NewContext()
	log.SetLogID(r.Header.Get("LogID"))
	log.Info("creating organisation")
	response := utility.NewResponse()
	var Organs *dto.CreateOrganisation
	err := json.NewDecoder(r.Body).Decode(&Organs)
	val := c.Repo.AlreadyOrganisation(Organs)
	if val == true {
		log.Error("Organisation name Already Exist")
		json.NewEncoder(w).Encode(response.Error(utility.ALREADY_EXIST, utility.GetCodeMsg(utility.ALREADY_EXIST)))

		return
	}
	if err != nil {
		log.Error("Error occured while parsing the request body, %s", err.Error())
		json.NewEncoder(w).Encode(response.Error(utility.BAD_REQUEST, utility.GetCodeMsg(utility.BAD_REQUEST)))
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}
	OrganisationId, err := c.Repo.CreateOrganisation(Organs)
	if err != nil {
		log.Error("Error occured while creating organisation, %s", err.Error())
		json.NewEncoder(w).Encode(response.Error(utility.SERVER_ERROR, utility.GetCodeMsg(utility.SERVER_ERROR)))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(response.Success(utility.SYSTEM001, utility.GetCodeMsg(utility.SYSTEM001), OrganisationId))

}

func (c *NewController) CreateBranch(w http.ResponseWriter, r *http.Request) {
	log := c.Logger.NewContext()
	log.SetLogID(r.Header.Get("LogID"))
	log.Info("creating branch")
	response := utility.NewResponse()
	var br *dto.CreateBranch
	err := json.NewDecoder(r.Body).Decode(&br)
	val := c.Repo.AlreadyBranch(br)
	if val == true {
		log.Error("Organisation name Already Exist")
		json.NewEncoder(w).Encode(response.Error(utility.ALREADY_EXIST, utility.GetCodeMsg(utility.ALREADY_EXIST)))

		return
	}
	if err != nil {
		log.Error("Error occured while parsing the request body, %s", err.Error())
		json.NewEncoder(w).Encode(response.Error(utility.BAD_REQUEST, utility.GetCodeMsg(utility.BAD_REQUEST)))
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	branchId, err := c.Repo.CreateBranch(br)
	if err != nil {
		log.Error("Error occured while creating branch, %s", err.Error())
		json.NewEncoder(w).Encode(response.Error(utility.SERVER_ERROR, utility.GetCodeMsg(utility.SERVER_ERROR)))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(response.Success(utility.SYSTEM001, utility.GetCodeMsg(utility.SYSTEM001), branchId))

}

func (c *NewController) GetOrganisation(w http.ResponseWriter, r *http.Request) {
	log := c.Logger.NewContext()
	log.SetLogID(r.Header.Get("LogID"))
	log.Info("getting organisation")
	response := utility.NewResponse()
	var br []bson.M
	cursor, err := c.Repo.GetOrganisation()
	if err != nil {
		log.Error("Error occured while getting organisation, %s", err.Error())
		json.NewEncoder(w).Encode(response.Error(utility.SERVER_ERROR, utility.GetCodeMsg(utility.SERVER_ERROR)))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = cursor.All(context.Background(), &br)
	json.NewEncoder(w).Encode(response.Success(utility.SYSTEM001, utility.GetCodeMsg(utility.SYSTEM001), br))
}

func (c *NewController) GetBranch(w http.ResponseWriter, r *http.Request) {
	log := c.Logger.NewContext()
	log.SetLogID(r.Header.Get("LogID"))
	log.Info("getting organisation")
	response := utility.NewResponse()
	params := mux.Vars(r)
	organisationID := params["organisationId"]
	var br []bson.M
	cursor, err := c.Repo.GetBranch(organisationID)
	if err != nil {
		log.Error("Error occured while getting organisation branches, %s", err.Error())
		json.NewEncoder(w).Encode(response.Error(utility.SERVER_ERROR, utility.GetCodeMsg(utility.SERVER_ERROR)))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = cursor.All(context.Background(), &br)
	json.NewEncoder(w).Encode(response.Success(utility.SYSTEM001, utility.GetCodeMsg(utility.SYSTEM001), br))
}

func (c *NewController) DeactivateOrganisation(w http.ResponseWriter, r *http.Request) {
	log := c.Logger.NewContext()
	log.Info("deactivating and activating organisation")
	response := utility.NewResponse()
	var Organs *dto.CreateOrganisation
	err := json.NewDecoder(r.Body).Decode(&Organs)
	if err != nil {
		log.Error("Error occured while parsing the request body, %s", err.Error())
		json.NewEncoder(w).Encode(response.Error(utility.BAD_REQUEST, utility.GetCodeMsg(utility.BAD_REQUEST)))
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}
	_, err = c.Repo.DeactivateOrganisation(Organs.OrganisationID, Organs.Active)

	if err != nil {
		log.Error("Error occured while deactivating or activating organisation, %s", err.Error())
		json.NewEncoder(w).Encode(response.Error(utility.SERVER_ERROR, utility.GetCodeMsg(utility.SERVER_ERROR)))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(response.PlainSuccess(utility.SYSTEM001, utility.GetCodeMsg(utility.SYSTEM001)))

}
