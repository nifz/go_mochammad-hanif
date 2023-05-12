package usecases

import (
	"code-competence/rest-api-jwt/models"
	"code-competence/rest-api-jwt/repositories"
)

type CategoryUsecase interface {
	GetAllCategories() ([]models.Category, error)
	GetCategoryByID(id uint) (models.Category, error)
	CreateCategory(category *models.Category) error
	UpdateCategory(category *models.Category) error
	DeleteCategory(id uint) error
}

type categoryUsecase struct {
	categoryRepo repositories.CategoryRepository
}

func NewCategoryUsecase(categoryRepo repositories.CategoryRepository) CategoryUsecase {
	return &categoryUsecase{categoryRepo}
}

func (u *categoryUsecase) GetAllCategories() ([]models.Category, error) {
	return u.categoryRepo.GetAllCategories()
}

func (u *categoryUsecase) GetCategoryByID(id uint) (models.Category, error) {
	return u.categoryRepo.GetCategoryByID(id)
}

func (u *categoryUsecase) CreateCategory(category *models.Category) error {
	return u.categoryRepo.CreateCategory(category)
}

func (u *categoryUsecase) UpdateCategory(category *models.Category) error {
	return u.categoryRepo.UpdateCategory(category)
}

func (u *categoryUsecase) DeleteCategory(id uint) error {
	var category models.Category
	category.ID = id
	return u.categoryRepo.DeleteCategory(category)
}
