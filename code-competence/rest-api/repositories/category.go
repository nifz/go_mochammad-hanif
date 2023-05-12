package repositories

import (
	"code-competence/rest-api/models"

	"gorm.io/gorm"
)

type CategoryRepository interface {
	GetAllCategories() ([]models.Category, error)
	GetCategoryByID(id uint) (models.Category, error)
	CreateCategory(category *models.Category) error
	UpdateCategory(category *models.Category) error
	DeleteCategory(category models.Category) error
}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepository{db}
}

func (r *categoryRepository) GetAllCategories() ([]models.Category, error) {
	var categories []models.Category
	err := r.db.Find(&categories).Error
	return categories, err
}

func (r *categoryRepository) GetCategoryByID(id uint) (models.Category, error) {
	var category models.Category
	err := r.db.Where("id = ?", id).First(&category).Error
	return category, err
}

func (r *categoryRepository) CreateCategory(category *models.Category) error {
	var _ = r.db.Create(category).Error
	return r.db.Where("id = ?", category.ID).First(&category).Error
}

func (r *categoryRepository) UpdateCategory(category *models.Category) error {
	return r.db.Save(category).Error
}

func (r *categoryRepository) DeleteCategory(category models.Category) error {
	// return nil
	return r.db.Delete(&category).Error
}
