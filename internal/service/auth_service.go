package service

import (
	// "log"
	// "runtime"

	// "github.com/mashingan/smapping"
	// "github.com/naveeharn/hospital-information-service-backend/helper"
	// "github.com/naveeharn/hospital-information-service-backend/internal/dto"
	// "github.com/naveeharn/hospital-information-service-backend/internal/entity"
	"github.com/naveeharn/hospital-information-service-backend/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	// VerifyCredential(username, password string) any
	VerifyCredential(username, password, hospital string) any
	// CreateStaff(staff dto.CreateStaff) (entity.Staff, error)
}

type authService struct {
	staffRepository repository.StaffRepository
}

// CreateStaff implements [AuthService].
// func (service *authService) CreateStaff(staff dto.CreateStaff) (entity.Staff, error) {
// 	staffBeforeCreate := entity.Staff{}
// 	if err := smapping.FillStruct(&staffBeforeCreate, smapping.MapFields(&staff)); err != nil {
// 		helper.LoggerErrorPath(runtime.Caller(0))
// 		log.Fatalf("Failed to map struct: %v", err)
// 	}
// 	createdStaff, err := service.staffRepository.CreateStaff(staffBeforeCreate)
// 	return createdStaff, err
// }

// VerifyCredential implements [AuthService].
func (service *authService) VerifyCredential(username string, password string, hospital string) any {
	staff, err := service.staffRepository.GetStaffByUsernameAndHospital(username, hospital)
	if err != nil {
		return nil
	}
	if !comparePassword(staff.Password, password) {
		return nil
	}
	return staff
}

func NewAuthService(staffRepository repository.StaffRepository) AuthService {
	return &authService{
		staffRepository: staffRepository,
	}
}

func comparePassword(hashedPassword, plainPassword string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword)); err != nil {
		return false
	}
	return true
}