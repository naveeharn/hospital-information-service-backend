package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/naveeharn/hospital-information-service-backend/helper"
	"github.com/naveeharn/hospital-information-service-backend/internal/dto"
	"github.com/naveeharn/hospital-information-service-backend/internal/service"
)

type PatientController interface {
	FindPatientByNationalIdOrPassportId(ctx *gin.Context)
	SearchPatient(ctx *gin.Context)
}

type patientController struct {
	patientService service.PatientService
	staffService service.StaffService
}

// FindPatientByNationalIdOrPassportId implements [PatientController].
func (controller *patientController) FindPatientByNationalIdOrPassportId(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		result := helper.Result{
			Code:    "P004999",
			Message: "Internal Service Error",
		}
		response := helper.CreateErrorResponse("bad request param id", nil, result)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	staffId, ok := ctx.Get("id")
	if !ok {
		result := helper.Result{
			Code:    "P004999",
			Message: "Internal Service Error",
		}
		response := helper.CreateErrorResponse("Staff id was not found in ctx", nil, result)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	staff, err := controller.staffService.GetStaffById(staffId.(string))
	if err != nil {
		result := helper.Result{
			Code:    "P004999",
			Message: "Internal Service Error",
		}
		response := helper.CreateErrorResponse("Staff id was not found in staff table", err.Error(), result)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	patient, err := controller.patientService.FindPatientByNationalIdOrPassportId(staff.Hospital, id)
	if err != nil {
		result := helper.Result{
			Code:    "P004001",
			Message: "id was not existed",
		}
		response := helper.CreateErrorResponse("id was not existed", err.Error(), result)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	result := helper.Result{
		Code:    "P004000",
		Message: "Success",
	}
	response := helper.CreateResponse(patient, result)
	ctx.JSON(http.StatusOK, response)
}

// SearchPatient implements [PatientController].
func (controller *patientController) SearchPatient(ctx *gin.Context) {
	patientSearchDTO := dto.PatientSerach{}
	err := ctx.ShouldBind(&patientSearchDTO)
	if err != nil {
		result := helper.Result{
			Code:    "P003999",
			Message: "Internal Service Error",
		}
		response := helper.CreateErrorResponse("bad request", err.Error(), result)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	staffId, ok := ctx.Get("id")
	if !ok {
		result := helper.Result{
			Code:    "P003999",
			Message: "Internal Service Error",
		}
		response := helper.CreateErrorResponse("Staff id was not found in ctx", nil, result)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	staff, err := controller.staffService.GetStaffById(staffId.(string))
	if err != nil {
		result := helper.Result{
			Code:    "P003999",
			Message: "Internal Service Error",
		}
		response := helper.CreateErrorResponse("Staff id was not found in staff table", err.Error(), result)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	log.Printf("staff: %v", staff)
	patientSearchDTO.Hospital = &staff.Hospital
	patients, err := controller.patientService.SearchPatient(patientSearchDTO)
	if err != nil {
		result := helper.Result{
			Code:    "P003999",
			Message: "Internal Service Error",
		}
		response := helper.CreateErrorResponse(" ", err.Error(), result)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	result := helper.Result{
		Code:    "P003000",
		Message: "Success",
	}
	response := helper.CreateResponse(patients, result)
	ctx.JSON(http.StatusOK, response)
}

func NewPatientController(patientService service.PatientService, staffService service.StaffService) PatientController {
	return &patientController{
		patientService: patientService,
		staffService: staffService,
	}
}
