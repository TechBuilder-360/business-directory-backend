package controller

import (
	"github.com/TechBuilder-360/business-directory-backend/domain/business/service"
	"github.com/gofiber/fiber/v2"
)

type BranchController interface {
	RegisterRoutes(router fiber.Router)
	//GetBranches(ctx *fiber.Ctx) apiError
}

type NewBranchController struct {
	Service service.BranchService
}

func (c *NewBranchController) RegisterRoutes(router fiber.Router) {
	_ = router.Group("/branches")

	//apis.Get("", middleware.Adapt(http.HandlerFunc(c.GetBranches), middleware.AuthorizeUserJWT()).ServeHTTP).Methods(http.MethodGet)
}

func DefaultBranchController() BranchController {
	return &NewBranchController{
		Service: service.NewBranchService(),
	}
}

// GetBranches godoc
// @Summary      get branches
// @Description  Get branches
// @Tags         Branch
// @Accept       json
// @Produce      json
// @Success      200      {object}  utils.SuccessResponse{Data=[]types.Branch}
// @Router       /branches [get]
//func (c *NewBranchController) GetBranches(ctx *fiber.Ctx) apiError {
//	logger := log.WithFields(log.Fields{constant.RequestIdentifier: utils.GenerateUUID()})
//	logger.Info("fetching single branch.")
//
//	// get organisation from context
//	organisation, err := middleware.OrganisationFromContext(r)
//	if err != nil {
//		logger.Error(err.Error())
//		w.WriteHeader(http.StatusBadRequest)
//		json.NewEncoder(w).Encode(utils.ErrorResponse{
//			Status:  false,
//			Message: "organisation not found",
//		})
//		return
//	}
//
//	data, err := c.Service.GetOrganisationBranches(organisation, logger)
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
