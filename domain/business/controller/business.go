package controller

import (
	"github.com/TechBuilder-360/business-directory-backend/domain/business/service"
	"github.com/TechBuilder-360/business-directory-backend/domain/user"
	"github.com/TechBuilder-360/business-directory-backend/internal/common/apiError"
	"github.com/TechBuilder-360/business-directory-backend/internal/common/types"
	"github.com/TechBuilder-360/business-directory-backend/internal/common/utils"
	"github.com/TechBuilder-360/business-directory-backend/internal/validation"
	"github.com/TechBuilder-360/business-directory-backend/pkg/log"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type IBusinessController interface {
	CreateBusiness(ctx *fiber.Ctx) error
	//ChangeActiveStatus(ctx *fiber.Ctx) apiError
	//GetSingleBusiness(ctx *fiber.Ctx) apiError
	//GetAllBusiness(ctx *fiber.Ctx) apiError
	RegisterRoutes(router *fiber.App)
}

type BusinessController struct {
	Service service.IBusinessService
}

func (c *BusinessController) RegisterRoutes(router *fiber.App) {
	_ = router.Group("/businesss")

	//apis.HandleFunc("", middleware.CacheClient.Middleware(middleware.Adapt(http.HandlerFunc(c.GetAllBusiness), middleware.AuthorizeUserJWT())).ServeHTTP).Methods(http.MethodGet)
	//apis.HandleFunc("", middleware.Adapt(http.HandlerFunc(c.CreateBusiness), middleware.AuthorizeUserJWT()).ServeHTTP).Methods(http.MethodPost)
	//apis.HandleFunc("/{id}", middleware.CacheClient.Middleware(middleware.Adapt(http.HandlerFunc(c.GetSingleBusiness), middleware.AuthorizeUserJWT())).ServeHTTP).Methods(http.MethodGet)
	//apis.HandleFunc("/activate", middleware.Adapt(http.HandlerFunc(c.ChangeActiveStatus), middleware.AuthorizeUserJWT(), middleware.AuthorizeBusinessJWT).ServeHTTP).Methods(http.MethodPatch)
}

func DefaultBusinessController() IBusinessController {
	return &BusinessController{
		Service: service.NewBusinessService(),
	}
}

func (c *BusinessController) CreateBusiness(ctx *fiber.Ctx) error {
	logger := log.LoggerInContext(ctx.UserContext())
	logger.Info("Creating Business")

	body := new(types.BusinessReq)

	if err, ok := validation.ValidateStruct(body, logger); !ok {
		return ctx.Status(http.StatusBadRequest).JSON(utils.ValidationError(apiError.AppError{
			Error:   err,
			Message: err,
		}))
	}

	// get user from context
	user, err := user.UserFromContext(ctx)
	if err != nil {
		logger.Error(err.Error())
		ctx.Status(http.StatusBadRequest)
		return ctx.JSON(utils.Error(apiError.AppError{
			Error:   err.Error(),
			Message: "request failed",
		}))
	}

	data, err := c.Service.CreateBusiness(body, user, logger)
	if err != nil {
		logger.Error(err.Error())
		ctx.Status(http.StatusBadRequest)
		return ctx.JSON(utils.Error(apiError.AppError{
			Error:   err.Error(),
			Message: err.Error(),
		}))
	}

	ctx.Status(http.StatusCreated)
	return ctx.JSON(utils.Success("successful", data, nil))

}

// GetBusiness godoc
// @Summary      get Business
// @Description  get Business
// @Tags         Business
// @Accept       json
// @Produce      json
// @Param        default  path	string  true  "Business ID"
// @Success      200      {object}  utils.SuccessResponse{types.Business}
// @Router       /Business/{id} [get]
//func (c *BusinessController) GetBusiness(ctx *fiber.Ctx) apiError {
//	logger := log.WithFields(log.Fields{constant.RequestIdentifier: utils.GenerateUUID()})
//	logger.Info("GetBusiness")
//
//	vars := mux.Vars(r)
//	id := vars["id"]
//
//	data, err := c.Service.GetBusiness(id)
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
// @Summary      activate/deactivate an Business
// @Description  activate/deactivate an Business
// @Tags         Businesss
// @Accept       json
// @Produce      json
// @Param        default  body	types.Activate  true  "change Business status"
// @Success      200      {object}  utils.SuccessResponse
// @Router       /Businesss/status [patch]
//func (c *BusinessController) ChangeActiveStatus(ctx *fiber.Ctx) apiError {
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
//	// get Business from context
//	Business, err := middleware.BusinessFromContext(r)
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
//	user, err := middleware.UserFromContext(r)
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
//	err = c.Service.ChangeBusinessStatus(Business, user, body, logger)
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

// GetSingleBusiness godoc
// @Summary      fetch an Business
// @Description  fetch an Business
// @Tags        Businesss
// @Accept       json
// @Produce      json
// @Param        default  body	id  true  "fetch an Business"
// @Success      200      {object}  utils.SuccessResponse{Data=types.Business}
// @Router       /Businesss/{id} [get]
//func (c *BusinessController) GetSingleBusiness(ctx *fiber.Ctx) apiError {
//	logger := log.WithFields(log.Fields{constant.RequestIdentifier: utils.GenerateUUID()})
//	logger.Info("GetSingleBusiness")
//	params := mux.Vars(r)
//	id := params["id"]
//
//	data, err := c.Service.GetSingleBusiness(id)
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

// GetAllBusiness godoc
// @Summary      fetch all Business
// @Description  fetch all Business
// @Tags         Businesss
// @Accept       json
// @Produce      json
// @Param        token    query     string  false  "token"
// @Success      200      {object}  utils.SuccessResponse{Data=data}
// @Router       /Businesss [get]
//func (c *BusinessController) GetAllBusiness(ctx *fiber.Ctx) apiError {
//	logger := log.WithFields(log.Fields{constant.RequestIdentifier: utils.GenerateUUID()})
//	logger.Info("Get All Businesss")
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
//	data, err := c.Service.GetAllBusiness(*filter, logger)
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
