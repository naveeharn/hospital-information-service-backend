package entity

import (
	"github.com/gofrs/uuid"
)

type Patient struct {
	Id            uuid.UUID `json:"-"`

	FirstNameTH   *string `json:"first_name_th,omitempty"`
	MiddleNameTH  *string `json:"middle_name_th,omitempty"`
	LastNameTH    *string `json:"last_name_th,omitempty"`

	FirstNameEN   *string `json:"first_name_en,omitempty"`
	MiddleNameEN  *string `json:"middle_name_en,omitempty"`
	LastNameEN    *string `json:"last_name_en,omitempty"`

	DateOfBirth   string `json:"date_of_birth"`

	PatientHn     string  `json:"patient_hn"`

	NationalId    *string `json:"national_id,omitempty"`
	PassportId    *string `json:"passport_id,omitempty"`

	PhoneNumber   *string `json:"phone_number,omitempty"`
	Email         *string `json:"email,omitempty"`

	Gender        *string `json:"gender,omitempty"`
	Hospital      *string `json:"hospital,omitempty"`
}