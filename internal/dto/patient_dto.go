package dto

type PatientSerach struct {
	FirstName  *string `json:"first_name,omitempty"`
	MiddleName *string `json:"middle_name,omitempty"`
	LastName   *string `json:"last_name,omitempty"`

	DateOfBirth *string `json:"date_of_birth"`

	NationalId *string `json:"national_id,omitempty"`
	PassportId *string `json:"passport_id,omitempty"`

	PhoneNumber *string `json:"phone_number,omitempty"`
	Email       *string `json:"email,omitempty"`

	Hospital *string `json:"hospital,omitempty"`
}
