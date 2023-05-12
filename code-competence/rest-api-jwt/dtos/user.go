package dtos

import "time"

type UserDTO struct {
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserLoginDTO struct {
	Email string `json:"email"`
	Token string `json:"token"`
}

type RegisterDTOResponse struct {
	Message string `json:"message"`
	Data    UserDTO
}

type LoginDTOResponse struct {
	Message string       `json:"message"`
	Data    UserLoginDTO `json:"data"`
}
