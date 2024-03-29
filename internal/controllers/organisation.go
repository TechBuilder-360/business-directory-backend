package controllers

import (
	"github.com/TechBuilder-360/business-directory-backend/internal/services"
	"github.com/gofiber/fiber/v2"
)

type IOrganisationController interface {
	//CreateOrganisation(ctx *fiber.Ctx) error
	//ChangeActiveStatus(ctx *fiber.Ctx) error
	//GetSingleOrganisation(ctx *fiber.Ctx) error
	//GetAllOrganisation(ctx *fiber.Ctx) error
	RegisterRoutes(router *fiber.App)
}

type organisationController struct {
	Service services.IOrganisationService
}

func (c *organisationController) RegisterRoutes(router *fiber.App) {
	_ = router.Group("/organisations")

	//apis.HandleFunc("", middlewares.CacheClient.Middleware(middlewares.Adapt(http.HandlerFunc(c.GetAllOrganisation), middlewares.AuthorizeUserJWT())).ServeHTTP).Methods(http.MethodGet)
	//apis.HandleFunc("", middlewares.Adapt(http.HandlerFunc(c.CreateOrganisation), middlewares.AuthorizeUserJWT()).ServeHTTP).Methods(http.MethodPost)
	//apis.HandleFunc("/{id}", middlewares.CacheClient.Middleware(middlewares.Adapt(http.HandlerFunc(c.GetSingleOrganisation), middlewares.AuthorizeUserJWT())).ServeHTTP).Methods(http.MethodGet)
	//apis.HandleFunc("/activate", middlewares.Adapt(http.HandlerFunc(c.ChangeActiveStatus), middlewares.AuthorizeUserJWT(), middlewares.AuthorizeOrganisationJWT).ServeHTTP).Methods(http.MethodPatch)
}

func DefaultOrganisationController() IOrganisationController {
	return &organisationController{
		Service: services.NewOrganisationService(),
	}
}

// CreateOrganisation godoc
// @Summary      create an organisation
// @Description  create an organisation
// @Tags         organisations
// @Accept       json
// @Produce      json
// @Param        default  body	types.CreateOrganisationReq  true  "create an organisation
// @Success      201      {object}  utils.SuccessResponse{Data=types.CreateOrganisationResponse
// @Router       /organisations [post]
//func (c *organisationController) CreateOrganisation(ctx *fiber.Ctx) error {
//	logger := log.WithFields(log.Fields{constant.RequestIdentifier: utils.GenerateUUID()})
//	logger.Info("Creating Organisation.")
//
//	body := &types.CreateOrganisationReq{}
//
//	err := json.NewDecoder(r.Body).Decode(body)
//	if err != nil {
//		logger.Error(err.Error())
//		w.WriteHeader(http.StatusBadRequest)
//		json.NewEncoder(w).Encode(utils.ErrorResponse{
//			Status:  false,
//			Message: "bad request",
//		})
//		return
//	}
//
//	if validation.ValidateStruct(w, body, logger) {
//		return
//	}
//
//	// get user from context
//	user, err := middlewares.UserFromContext(r)
//	if err != nil {
//		logger.Error(err.Error())
//		w.WriteHeader(http.StatusBadRequest)
//		json.NewEncoder(w).Encode(utils.ErrorResponse{
//			Status:  false,
//			Message: err.Error(),
//		})
//		return
//	}
//
//	data, err := c.Service.CreateOrganisation(body, user, logger)
//	if err != nil {
//		logger.Error(err.Error())
//		w.WriteHeader(http.StatusBadRequest)
//		json.NewEncoder(w).Encode(utils.ErrorResponse{
//			Status:  false,
//			Message: err.Error(),
//		})
//		return
//	}
//
//	w.WriteHeader(http.StatusCreated)
//	json.NewEncoder(w).Encode(utils.SuccessResponse{
//		Status:  true,
//		Message: "Successful",
//		Data:    data,
//	})
//
//}

// GetOrganisation godoc
// @Summary      get organisation
// @Description  get organisation
// @Tags         Organisation
// @Accept       json
// @Produce      json
// @Param        default  path	string  true  "organisation ID"
// @Success      200      {object}  utils.SuccessResponse{types.Organisation}
// @Router       /organisation/{id} [get]
//func (c *organisationController) GetOrganisation(ctx *fiber.Ctx) error {
//	logger := log.WithFields(log.Fields{constant.RequestIdentifier: utils.GenerateUUID()})
//	logger.Info("GetOrganisation")
//
//	vars := mux.Vars(r)
//	id := vars["id"]
//
//	data, err := c.Service.GetOrganisation(id)
//	if err != nil {
//		logger.Error(err.Error())
//		w.WriteHeader(http.StatusBadRequest)
//		json.NewEncoder(w).Encode(utils.ErrorResponse{
//			VerificationType:  false,
//			Message: err.Error(),
//		})
//		return
//	}
//
//	w.WriteHeader(http.StatusOK)
//	json.NewEncoder(w).Encode(utils.SuccessResponse{
//		VerificationType:  true,
//		Message: "Successful",
//		Data:    data,
//	})
//
//}

// ChangeActiveStatus godoc
// @Summary      activate/deactivate an organisation
// @Description  activate/deactivate an organisation
// @Tags         organisations
// @Accept       json
// @Produce      json
// @Param        default  body	types.Activate  true  "change organisation status"
// @Success      200      {object}  utils.SuccessResponse
// @Router       /organisations/status [patch]
//func (c *organisationController) ChangeActiveStatus(ctx *fiber.Ctx) error {
//	logger := log.WithFields(log.Fields{constant.RequestIdentifier: utils.GenerateUUID()})
//	logger.Info("ChangeStatus")
//
//	body := &types.Activate{}
//
//	err := json.NewDecoder(r.Body).Decode(body)
//	if err != nil {
//		logger.Error(err.Error())
//		w.WriteHeader(http.StatusBadRequest)
//		json.NewEncoder(w).Encode(utils.ErrorResponse{
//			Status:  false,
//			Message: "bad request",
//		})
//		return
//	}
//
//	if validation.ValidateStruct(w, body, logger) {
//		return
//	}
//
//	// get organisation from context
//	organisation, err := middlewares.OrganisationFromContext(r)
//	if err != nil {
//		logger.Error(err.Error())
//		w.WriteHeader(http.StatusBadRequest)
//		json.NewEncoder(w).Encode(utils.ErrorResponse{
//			Status:  false,
//			Message: err.Error(),
//		})
//		return
//	}
//
//	// get user from context
//	user, err := middlewares.UserFromContext(r)
//	if err != nil {
//		logger.Error(err.Error())
//		w.WriteHeader(http.StatusBadRequest)
//		json.NewEncoder(w).Encode(utils.ErrorResponse{
//			Status:  false,
//			Message: err.Error(),
//		})
//		return
//	}
//
//	err = c.Service.ChangeOrganisationStatus(organisation, user, body, logger)
//	if err != nil {
//		logger.Error(err.Error())
//		w.WriteHeader(http.StatusBadRequest)
//		json.NewEncoder(w).Encode(utils.ErrorResponse{
//			Status:  false,
//			Message: err.Error(),
//		})
//		return
//	}
//
//	w.WriteHeader(http.StatusOK)
//	json.NewEncoder(w).Encode(utils.SuccessResponse{
//		Status:  true,
//		Message: "Successful",
//		Data:    nil,
//	})
//
//}

// GetSingleOrganisation godoc
// @Summary      fetch an organisation
// @Description  fetch an organisation
// @Tags        organisations
// @Accept       json
// @Produce      json
// @Param        default  body	id  true  "fetch an organisation"
// @Success      200      {object}  utils.SuccessResponse{Data=types.Organisation}
// @Router       /organisations/{id} [get]
//func (c *organisationController) GetSingleOrganisation(ctx *fiber.Ctx) error {
//	logger := log.WithFields(log.Fields{constant.RequestIdentifier: utils.GenerateUUID()})
//	logger.Info("GetSingleOrganisation")
//	params := mux.Vars(r)
//	id := params["id"]
//
//	data, err := c.Service.GetSingleOrganisation(id)
//	if err != nil {
//		logger.Error(err.Error())
//		w.WriteHeader(http.StatusBadRequest)
//		json.NewEncoder(w).Encode(utils.ErrorResponse{
//			Status:  false,
//			Message: err.Error(),
//		})
//		return
//	}
//
//	w.WriteHeader(http.StatusOK)
//	json.NewEncoder(w).Encode(utils.SuccessResponse{
//		Status:  true,
//		Message: "Successful",
//		Data:    data,
//	})
//
//}

// GetAllOrganisation godoc
// @Summary      fetch all organisation
// @Description  fetch all organisation
// @Tags         organisations
// @Accept       json
// @Produce      json
// @Param        token    query     string  false  "token"
// @Success      200      {object}  utils.SuccessResponse{Data=data}
// @Router       /organisations [get]
//func (c *organisationController) GetAllOrganisation(ctx *fiber.Ctx) error {
//	logger := log.WithFields(log.Fields{constant.RequestIdentifier: utils.GenerateUUID()})
//	logger.Info("Get All Organisations")
//
//	filter := &types.Query{}
//	if err := schema.NewDecoder().Decode(filter, r.Form); err != nil {
//		logger.Error(err)
//		w.WriteHeader(http.StatusBadRequest)
//		json.NewEncoder(w).Encode(utils.ErrorResponse{
//			Status:  false,
//			Message: err.Error(),
//		})
//		return
//	}
//
//	filter.CleanUp()
//
//	data, err := c.Service.GetAllOrganisation(*filter, logger)
//	if err != nil {
//		logger.Error(err.Error())
//		w.WriteHeader(http.StatusBadRequest)
//		json.NewEncoder(w).Encode(utils.ErrorResponse{
//			Status:  false,
//			Message: err.Error(),
//		})
//		return
//	}
//
//	w.WriteHeader(http.StatusOK)
//	json.NewEncoder(w).Encode(utils.SuccessResponse{
//		Status:  true,
//		Message: "Successful",
//		Data:    data.Data,
//		Meta:    data,
//	})
//
//}
