package dtos

import "time"

type CategoriesDTO struct {
	Name string `json:"name"`
}

type CategoryDTO struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"deleted_at"`
}

type CategoryDTOResponse struct {
	Message string      `json:"message"`
	Data    CategoryDTO `json:"data"`
}

type CategoryDTOsResponse struct {
	Message string        `json:"message"`
	Data    []CategoryDTO `json:"data"`
}
