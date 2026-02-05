package repository

import (
	"database/sql"

	"github.com/naveeharn/hospital-information-service-backend/internal/entity"
)

type HospitalRepository interface {
	GetHospital() ([]entity.Hospital, error)
}

type hospitalConnection struct {
	connection *sql.DB
}

// GetHospital implements [HospitalRepository].
func (h *hospitalConnection) GetHospital() ([]entity.Hospital, error) {
	// panic("unimplemented")
	rows, err := h.connection.Query(`
		SELECT 
			id,
			hospital_name 
		FROM 
			hospital
	`)
	if err != nil {
		return []entity.Hospital{}, err
	}
	defer rows.Close()
	var hospitals []entity.Hospital
	for rows.Next() {
		var hospital entity.Hospital
		err := rows.Scan(&hospital.Id, &hospital.HospitalName)
		if err != nil {
			return []entity.Hospital{}, err
		}
		hospitals = append(hospitals, hospital)
	}
	if err := rows.Err(); err != nil {
		return []entity.Hospital{}, err
	}
	return hospitals, nil

}

func NewHospitalRepository(db *sql.DB) HospitalRepository {
	return &hospitalConnection{connection: db}
}
