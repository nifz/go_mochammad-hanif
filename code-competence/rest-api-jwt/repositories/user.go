package repositories

import (
	"code-competence/rest-api-jwt/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetUserByEmail(email string) (models.User, error)
	CreateUser(user models.User) (models.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) GetUserByEmail(email string) (models.User, error) {
	var item models.User
	err := r.db.Where("email = ?", email).First(&item).Error
	return item, err
}

func (r *userRepository) CreateUser(user models.User) (models.User, error) {
	err := r.db.Create(&user).Error
	return user, err
}
