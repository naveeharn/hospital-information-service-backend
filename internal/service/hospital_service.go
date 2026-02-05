package service

import (
	"github.com/naveeharn/hospital-information-service-backend/internal/entity"
	"github.com/naveeharn/hospital-information-service-backend/internal/repository"
)

type HospitalService interface {
	GetHospital() ([]entity.Hospital, error)
}

type hospitalService struct {
	hospitalRepository repository.HospitalRepository
}

func NewHospitalService(hospitalRepository repository.HospitalRepository) HospitalService  {
	return &hospitalService{
		hospitalRepository: hospitalRepository ,
	}
}

func (service *hospitalService) GetHospital() ([]entity.Hospital, error) {
	hospitals, err := service.hospitalRepository.GetHospital()
	return hospitals, err
}