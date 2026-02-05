package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/naveeharn/hospital-information-service-backend/internal/service"
)

type HospitalController interface {
	GetHospital(ctx *gin.Context)
	CreateStaff(ctx *gin.Context)
}

type hospitalController struct {
	hospitalService service.HospitalService
}

func NewHospitalController(hospitalService service.HospitalService) HospitalController {
	return &hospitalController{
		hospitalService: hospitalService,
	}
}

// GetHospital implements [HospitalController].
func (controller *hospitalController) GetHospital(ctx *gin.Context) {
	hospitals, err := controller.hospitalService.GetHospital()
	if err != nil {
		return
	}
	ctx.JSON(http.StatusOK, hospitals)
}

// CreateStaff implements [HospitalController].
func (controller *hospitalController) CreateStaff(ctx *gin.Context) {
	panic("unimplemented")
}