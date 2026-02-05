package dto

type CreateStaff struct {
	Username string `json:"username" binding:"required,alpha,min=3,max=255"`
	Password string `json:"password" binding:"required,min=3,max=255"`
	Hospital string `json:"hospital" binding:"required,min=3,max=255"`
}