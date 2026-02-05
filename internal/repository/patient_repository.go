package repository

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/naveeharn/hospital-information-service-backend/internal/dto"
	"github.com/naveeharn/hospital-information-service-backend/internal/entity"
)

type PatientRepository interface {
	FindPatientByNationalIdOrPassportId(hospital, id string) (entity.Patient, error)
	SearchPatient(p dto.PatientSerach) ([]entity.Patient, error)
}

type patientConnection struct {
	connection *sql.DB
}

// SearchPatient implements [PatientRepository].
func (db *patientConnection) SearchPatient(p dto.PatientSerach) ([]entity.Patient, error) {
	query := `
		SELECT
			id,
			first_name_th,
			middle_name_th,
			last_name_th,
			first_name_en,
			middle_name_en,
			last_name_en,
			TO_CHAR(date_of_birth, 'YYYY-MM-DD') AS date_of_birth,
			patient_hn,
			national_id,
			passport_id,
			phone_number,
			email,
			gender,
			hospital
		FROM patient
		WHERE
			1=1
			AND hospital = $1

	`
	params := []any{*p.Hospital}
	paramPos := len(params) + 1
	if p.FirstName != nil {
		query += fmt.Sprintf(" AND (first_name_th ILIKE $%d OR first_name_en ILIKE $%d)", paramPos, paramPos)
		
		params = append(params, "%" + *p.FirstName + "%")
		paramPos++
	}
	if p.MiddleName != nil {
		query += fmt.Sprintf(" AND (middle_name_th ILIKE $%d OR middle_name_en ILIKE $%d)", paramPos, paramPos)
		params = append(params, "%" + *p.MiddleName + "%")
		paramPos++
	}
	if p.LastName != nil {
		query += fmt.Sprintf(" AND (last_name_th ILIKE $%d OR last_name_en ILIKE $%d)", paramPos, paramPos)
		params = append(params, "%" + *p.LastName + "%")
		paramPos++
	}
	if p.DateOfBirth != nil {
		query += fmt.Sprintf(" AND (date_of_birth = ($%d)::TIMESTAMP)", paramPos)
		params = append(params, *p.DateOfBirth)
		paramPos++
	}
	if p.NationalId != nil {
		query += fmt.Sprintf(" AND (national_id LIKE $%d)", paramPos)
		params = append(params, "%" + *p.NationalId + "%")
		paramPos++
	}
	if p.PassportId != nil {
		query += fmt.Sprintf(" AND (passport_id LIKE $%d)", paramPos)
		params = append(params, "%" + *p.PassportId + "%")
		paramPos++
	}
	if p.PhoneNumber != nil {
		query += fmt.Sprintf(" AND (phone_number LIKE $%d)", paramPos)
		params = append(params, "%" + *p.PassportId + "%")
		paramPos++
	}
	if p.Email != nil {
		query += fmt.Sprintf(" AND (hospital LIKE $%d)", paramPos)
		params = append(params, "%" + *p.PassportId + "%")
		paramPos++
	}
	log.Printf("query: %v", query)
	log.Printf("params: %v", params)
	rows, err := db.connection.Query(query, params...)
	if err != nil {
		return []entity.Patient{}, err
	}
	defer rows.Close()
	var patients []entity.Patient
	for rows.Next() {
		var patient entity.Patient
		if err := rows.Scan(
			&patient.Id,
			&patient.FirstNameTH,
			&patient.MiddleNameTH,
			&patient.LastNameTH,
			&patient.FirstNameEN,
			&patient.MiddleNameEN,
			&patient.LastNameEN,
			&patient.DateOfBirth,
			&patient.PatientHn,
			&patient.NationalId,
			&patient.PassportId,
			&patient.PhoneNumber,
			&patient.Email,
			&patient.Gender,
			&patient.Hospital,
		); err != nil {
			return nil, err
		}
		patients = append(patients, patient)
	}
	return patients, err
}

// FindPatientByNationalIdOrPassportId implements [PatientRepository].
func (db *patientConnection) FindPatientByNationalIdOrPassportId(hospital, id string) (entity.Patient, error) {
	patient := entity.Patient{}
	query := `
		SELECT
			id,
			first_name_th,
			middle_name_th,
			last_name_th,
			first_name_en,
			middle_name_en,
			last_name_en,
			TO_CHAR(date_of_birth, 'YYYY-MM-DD') AS date_of_birth,
			patient_hn,
			national_id,
			passport_id,
			phone_number,
			email,
			gender,
			hospital
		FROM patient
		WHERE
			hospital = $1
			AND (national_id = $2 OR passport_id = $2)
	`
	params := []any{hospital, id}
	scanner := []any{
		&patient.Id,
		&patient.FirstNameTH,
		&patient.MiddleNameTH,
		&patient.LastNameTH,
		&patient.FirstNameEN,
		&patient.MiddleNameEN,
		&patient.LastNameEN,
		&patient.DateOfBirth,
		&patient.PatientHn,
		&patient.NationalId,
		&patient.PassportId,
		&patient.PhoneNumber,
		&patient.Email,
		&patient.Gender,
		&patient.Hospital,
	}
	err := db.connection.QueryRow(query, params...).Scan(scanner...)
	if err != nil {
		log.Printf("FindPatientByNationalIdOrPassportId Error: %s", err.Error())
		return entity.Patient{}, err
	}
	return patient, nil
}

func NewPatientRepository(db *sql.DB) PatientRepository {
	return &patientConnection{
		connection: db,
	}
}
