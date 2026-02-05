package main

import (
	"database/sql"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/naveeharn/hospital-information-service-backend/config"
	"github.com/naveeharn/hospital-information-service-backend/internal/controller"
	"github.com/naveeharn/hospital-information-service-backend/internal/repository"
	"github.com/naveeharn/hospital-information-service-backend/internal/service"
	"github.com/naveeharn/hospital-information-service-backend/middleware"
	// "gorm.io/gorm"
)

var (
	db *sql.DB = config.SetupDatabaseConnection()

	// repository
	// hospitalRepository repository.HospitalRepository = repository.NewHospitalRepository(db)
	staffRepository repository.StaffRepository = repository.NewStaffRepository(db)
	patientRepository repository.PatientRepository = repository.NewPatientRepository(db)

	// service
	jwtService service.JwtService = service.NewJwtSwervice()
	// hospitalService service.HospitalService = service.NewHospitalService(hospitalRepository)
	staffService service.StaffService = service.NewStaffService(staffRepository)
	authService service.AuthService = service.NewAuthService(staffRepository)
	patientService service.PatientService = service.NewPatientService(patientRepository)

	// controller
	// hospitalController controller.HospitalController = controller.NewHospitalController(hospitalService)
	staffController controller.StaffController = controller.NewStaffController(staffService)
	authController controller.AuthController = controller.NewAuthController(authService, jwtService)
	patientController controller.PatientController = controller.NewPatientController(patientService, staffService)

)

func main() {

	// defer config.CloseDatabaseConnection(db)
	gin.SetMode((gin.ReleaseMode))
	routers := gin.Default()
	routers.Use(cors.Default())

	routers.GET("/healthcheck", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"status": "ok"})
	})

	// routers.GET("/hospital", hospitalController.GetHospital)

	routers.POST("/staff/create", staffController.CreateStaff)
	routers.POST("/staff/login", authController.Login)

	// authRoutes := routers.Group("staff")
	// {
	// 	authRoutes.POST("/create", staffController.CreateStaff)
	// 	authRoutes.POST("/login", authController.Login)
	// }

	patientRoutes := routers.Group("patient", middleware.AuthorizeJWT(jwtService))
	{
		patientRoutes.POST("/search", patientController.SearchPatient)
		patientRoutes.GET("/search/:id", patientController.FindPatientByNationalIdOrPassportId)
	}

	err := routers.Run(":8080")
	if err != nil {
		log.Fatal("Something went wrong")
	}
}
