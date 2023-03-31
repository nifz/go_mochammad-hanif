package models

import "gorm.io/gorm"

type Blog struct {
	gorm.Model
	UserID      int    `json:"user_id" form:"user_id"`
	User        User   `gorm:"foreignKey:user_id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user"`
	Title       string `json:"title" form:"title"`
	Description string `json:"description" form:"description"`
}
