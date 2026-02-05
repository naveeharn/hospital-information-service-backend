package repository

import (
	"database/sql"
	"log"
	"runtime"
	"strings"

	// "github.com/gofrs/uuid"
	"github.com/naveeharn/hospital-information-service-backend/helper"
	"github.com/naveeharn/hospital-information-service-backend/internal/entity"
	"golang.org/x/crypto/bcrypt"
)

type StaffRepository interface {
	CreateStaff(staff entity.Staff) (entity.Staff, error)
	IsDuplicatedUsername(username string) (entity.Staff, error)
	IsDuplicatedUsernameAndHospital(username, hospital string) (entity.Staff, error)
	GetStaffByUsername(username string) (entity.Staff, error)
	GetStaffByUsernameAndHospital(username, hospital string) (entity.Staff, error)
	GetStaffById(id string) (entity.Staff, error)
}

type staffConnection struct {
	connection *sql.DB
}

// CreateStaff implements [StaffRepository].
func (db *staffConnection) CreateStaff(staff entity.Staff) (entity.Staff, error) {
	// id, err := uuid.NewV4()
	// if err != nil {
	// 	return entity.Staff{}, err
	// }
	// staff.Id = id.String()
	staff.Password = hashAndSalt(staff.Password)
	err := db.connection.QueryRow(`
		INSERT INTO staff (id, username, password, hospital)
		VALUES 
			(gen_random_uuid(), $1, $2, $3)
		RETURNING id
		;
	`, staff.Username, staff.Password, staff.Hospital).Scan(&staff.Id)
	if err != nil {
		return entity.Staff{}, err
	}
	return staff, nil
}

// IsDuplicatedUsername implements [StaffRepository].
func (db *staffConnection) IsDuplicatedUsername(username string) (entity.Staff, error) {
	staff := entity.Staff{}
	err := db.connection.QueryRow(`
		SELECT id
		FROM staff
		WHERE username = $1
		;
	`, username).Scan(&staff.Id)
	if err != nil {
		return entity.Staff{}, err
	}
	return staff, nil
}

// IsDuplicatedUsernameAndHospital implements [StaffRepository].
func (db *staffConnection) IsDuplicatedUsernameAndHospital(username string, hospital string) (entity.Staff, error) {
	staff := entity.Staff{}
	query := `
		SELECT id
		FROM staff
		WHERE 
			username = $1
			AND hospital = $2
		;
	`
	params := []any{username, hospital}
	scanner := []any{&staff.Id}
	err := db.connection.
		QueryRow(query, params...).
		Scan(scanner...)
	if err != nil && !strings.Contains(err.Error(), "sql: no rows in result set") {
		log.Panicf("IsDuplicatedUsernameAndHospital Error: %s", err.Error())
		return entity.Staff{}, err
	}
	return staff, nil
}

// GetStaffByUsername implements [StaffRepository].
func (db *staffConnection) GetStaffByUsername(username string) (entity.Staff, error) {
	staff := entity.Staff{}
	err := db.connection.QueryRow(`
		SELECT id, username, password, hospital
		FROM staff
		WHERE username = $1
		;
	`, username).Scan(&staff.Id, &staff.Username, &staff.Password, &staff.Hospital)
	if err != nil {
		return entity.Staff{}, err
	}
	return staff, nil
}

// GetStaffByUsernameAndHospital implements [StaffRepository].
func (db *staffConnection) GetStaffByUsernameAndHospital(username string, hospital string) (entity.Staff, error) {
	staff := entity.Staff{}
	// err := db.connection.QueryRow(`
	// 	SELECT id, username, password, hospital
	// 	FROM staff
	// 	WHERE
	// 		username = $1
	// 		AND hospital = $2
	// 	;
	// `, username, hospital).Scan(&staff.Id, &staff.Username, &staff.Password, &staff.Hospital)
	query := `
		SELECT id, username, password, hospital
		FROM staff
		WHERE 
			username = $1
			AND hospital = $2
		;
	`
	params := []any{username, hospital}
	scanner := []any{&staff.Id, &staff.Username, &staff.Password, &staff.Hospital}
	err := db.connection.
		QueryRow(query, params...).
		Scan(scanner...)
	if err != nil {
		return entity.Staff{}, err
	}
	return staff, nil
}

// GetStaffById implements [StaffRepository].
func (db *staffConnection) GetStaffById(id string) (entity.Staff, error) {
	staff := entity.Staff{}
	query := `
		SELECT id, username, password, hospital
		FROM staff
		WHERE 
			id = $1
		;
	`
	params := []any{id}
	scanner := []any{&staff.Id, &staff.Username, &staff.Password, &staff.Hospital}
	err := db.connection.
		QueryRow(query, params...).
		Scan(scanner...)
	if err != nil {
		return entity.Staff{}, err
	}
	return staff, nil
}

func NewStaffRepository(db *sql.DB) StaffRepository {
	return &staffConnection{connection: db}
}

func hashAndSalt(password string) string {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		helper.LoggerErrorPath(runtime.Caller(0))
		log.Fatalf("Failed to hash a password %s\n error:%s", password, err.Error())
	}
	return string(hashed)
}
