package service

import (
	"github.com/naveeharn/hospital-information-service-backend/internal/dto"
	"github.com/naveeharn/hospital-information-service-backend/internal/entity"
	"github.com/naveeharn/hospital-information-service-backend/internal/repository"
)

type PatientService interface {
	FindPatientByNationalIdOrPassportId(hospital, id string) (entity.Patient, error)
	SearchPatient(p dto.PatientSerach) ([]entity.Patient, error)
}

type patientService struct {
	patientRepository repository.PatientRepository
}

// FindPatientByNationalIdOrPassportId implements [PatientService].
func (service *patientService) FindPatientByNationalIdOrPassportId(hospital, id string) (entity.Patient, error) {
	patient, err := service.patientRepository.FindPatientByNationalIdOrPassportId(hospital, id)
	if err != nil {
		return entity.Patient{}, err
	}
	return patient, nil
}

// SearchPatient implements [PatientService].
func (service *patientService) SearchPatient(p dto.PatientSerach) ([]entity.Patient, error) {
	patients, err := service.patientRepository.SearchPatient(p)
	if err != nil {
		return []entity.Patient{}, err
	}
	return patients, nil
}

func NewPatientService(patientRepository repository.PatientRepository) PatientService {
	return &patientService{
		patientRepository: patientRepository,
	}
}
