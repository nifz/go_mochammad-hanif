package models

import (
	"time"

	"gorm.io/gorm"
)

type Item struct {
	ID          string         `gorm:"primaryKey" json:"id"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Stock       uint           `json:"stock"`
	Price       uint           `json:"price"`
	CategoryID  uint           `json:"category_id"`
	Category    Category       `gorm:"foreignKey:CategoryID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"category"`
}
