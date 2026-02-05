package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/naveeharn/hospital-information-service-backend/helper"
	"github.com/naveeharn/hospital-information-service-backend/internal/dto"
	"github.com/naveeharn/hospital-information-service-backend/internal/entity"
	"github.com/naveeharn/hospital-information-service-backend/internal/service"
)

type AuthController interface {
	Login(ctx *gin.Context)
}

type authController struct {
	authService service.AuthService
	jwtService  service.JwtService
}

// Login implements [AuthController].
func (controller *authController) Login(ctx *gin.Context) {
	loginDTO := dto.Login{}
	if err := ctx.ShouldBind(&loginDTO); err != nil {
		result := helper.Result{
			Code:    "P001999",
			Message: "Internal Service Error",
		}
		response := helper.CreateErrorResponse("bad request", err.Error(), result)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	verifiedStaff := controller.authService.VerifyCredential(loginDTO.Username, loginDTO.Password, loginDTO.Hospital)
	staff, ok := verifiedStaff.(entity.Staff)
	if !ok {
		result := helper.Result{
			Code:    "P001999",
			Message: "Internal Service Error",
		}
		response := helper.CreateErrorResponse("username or password invalid", nil, result)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	accessToken := controller.jwtService.GenerateToken(staff.Id)
	data := map[string]string{"accessToken": accessToken}
	result := helper.Result{
		Code:    "P001000",
		Message: "Success",
	}
	response := helper.CreateResponse(data, result)
	ctx.JSON(http.StatusOK, response)
}

func NewAuthController(authService service.AuthService, jwtService service.JwtService) AuthController {
	return &authController{
		authService: authService,
		jwtService:  jwtService,
	}
}
