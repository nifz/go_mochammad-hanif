package repositories

import (
	"praktikum/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(data models.User) error
	GetAll() ([]models.User, error)
	LoginUser(email, password string) (models.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (r *userRepository) Create(data models.User) error {
	return r.db.Create(&data).Error
}

func (r *userRepository) GetAll() ([]models.User, error) {
	var users []models.User
	err := r.db.Find(&users).Error
	return users, err
}

func (r *userRepository) LoginUser(email, password string) (models.User, error) {
	var user models.User
	err := r.db.Where("email = ? AND password = ?", email, password).First(&user).Error

	return user, err

}
