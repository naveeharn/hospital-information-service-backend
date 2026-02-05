package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/naveeharn/hospital-information-service-backend/helper"
	"github.com/naveeharn/hospital-information-service-backend/internal/dto"
	"github.com/naveeharn/hospital-information-service-backend/internal/service"
)

type StaffController interface {
	CreateStaff(ctx *gin.Context)
}

type staffController struct {
	staffService service.StaffService
}

// CreateStaff implements [StaffController].
func (controller *staffController) CreateStaff(ctx *gin.Context) {
	staffDTO := dto.CreateStaff{}
	if err := ctx.ShouldBind(&staffDTO); err != nil {
		result := helper.Result{
			Code: "P002999",
			Message: "Internal Service Error",
		}
		response := helper.CreateErrorResponse("bad request", err.Error(), result)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	if controller.staffService.IsDuplicatedUsernameAndHospital(staffDTO.Username, staffDTO.Hospital) {
		result := helper.Result{
			Code: "P002998",
			Message: "username and hospital must not be duplicated",
		}
		response := helper.CreateErrorResponse("username and hospital were duplicated", nil, result)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	staff, err := controller.staffService.CreateStaff(staffDTO)
	if err != nil || staff.Id == "" {
		result := helper.Result{
			Code: "P002999",
			Message: "Internal Service Error",
		}
		response := helper.CreateErrorResponse("internal server error", nil, result)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	result := helper.Result{
		Code: "P002000",
		Message: "Success",
	}
	response := helper.CreateResponse(nil, result)
	ctx.JSON(http.StatusOK, response)
}

func NewStaffController(staffService service.StaffService) StaffController {
	return &staffController{
		staffService: staffService,
	}
}
