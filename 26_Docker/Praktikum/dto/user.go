package dto

import "praktikum/models"

type CreateUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserToken struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Token    string `json:"token"`
}

type UserResponse struct {
	Message string        `json:"message"`
	Data    []models.User `json:"data"`
}
