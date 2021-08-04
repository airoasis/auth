package model

type UserAuthRequestDTO struct {
	Username    string	`json:"username" binding:"required,min=4"`
	Password	string	`json:"password" binding:"required,min=4"`
}

type UserRequestDTO struct {
	Username    string	`json:"username" binding:"required,min=4"`
	Password	string	`json:"password" binding:"required,min=4"`
	Email   	string	`json:"email" binding:"required,email"`
}

type UserResponseDTO struct {
	ID			uint	`json:"id"`
	Username    string	`json:"username"`
	Email   	string	`json:"email"`
	AcapyToken	string	`json:"acapyToken"`
}