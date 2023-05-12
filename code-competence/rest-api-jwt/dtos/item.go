package dtos

import "time"

type ItemDTO struct {
	ID          string        `json:"id"`
	Name        string        `json:"name"`
	Description string        `json:"description"`
	Stock       uint          `json:"stock"`
	Price       uint          `json:"price"`
	CategoryID  uint          `json:"category_id"`
	Category    CategoriesDTO `json:"category"`
	CreatedAt   time.Time     `json:"created_at"`
	UpdatedAt   time.Time     `json:"updated_at"`
}

type ItemDTOResponse struct {
	Message string  `json:"message"`
	Data    ItemDTO `json:"data"`
}

type ItemDTOsResponse struct {
	Message string    `json:"message"`
	Data    []ItemDTO `json:"data"`
}
