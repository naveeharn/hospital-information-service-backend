package service

import (
	"log"
	"runtime"

	"github.com/mashingan/smapping"
	"github.com/naveeharn/hospital-information-service-backend/helper"
	"github.com/naveeharn/hospital-information-service-backend/internal/dto"
	"github.com/naveeharn/hospital-information-service-backend/internal/entity"
	"github.com/naveeharn/hospital-information-service-backend/internal/repository"
)

type StaffService interface {
	CreateStaff(staff dto.CreateStaff) (entity.Staff, error)
	IsDuplicatedUsernameAndHospital(username, hospital string) bool
	GetStaffById(id string) (entity.Staff, error)
}

type staffService struct {
	staffRepository repository.StaffRepository
}

// CreateStaff implements [StaffService].
func (service *staffService) CreateStaff(staff dto.CreateStaff) (entity.Staff, error) {
	staffBeforeCreate := entity.Staff{}
	if err := smapping.FillStruct(&staffBeforeCreate, smapping.MapFields(&staff)); err != nil {
		helper.LoggerErrorPath(runtime.Caller(0))
		log.Fatalf("Failed to map struct: %v", err)
	}
	createdStaff, err := service.staffRepository.CreateStaff(staffBeforeCreate)
	return createdStaff, err
}

// IsDuplicatedUsernameAndHospital implements [StaffService].
func (service *staffService) IsDuplicatedUsernameAndHospital(username, hospital string) bool {
	staff, _ := service.staffRepository.IsDuplicatedUsernameAndHospital(username, hospital)
	// if err != nil {
	// 	log.Printf("err:%s", err.Error())
	// 	return true
	// }
	return staff.Id != ""
}

// GetStaffById implements [StaffService].
func (service *staffService) GetStaffById(id string) (entity.Staff, error) {
	staff, err := service.staffRepository.GetStaffById(id)
	if err != nil {
		return entity.Staff{}, err
	}
	return staff, nil
}

func NewStaffService(staffRepository repository.StaffRepository) StaffService {
	return &staffService{
		staffRepository: staffRepository,
	}
}
