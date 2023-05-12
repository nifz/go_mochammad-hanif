package controllers

import (
	"code-competence/rest-api-jwt/dtos"
	"code-competence/rest-api-jwt/models"
	"code-competence/rest-api-jwt/usecases"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type CategoryController interface {
	GetAllCategories(c echo.Context) error
	CreateCategory(c echo.Context) error
	UpdateCategory(c echo.Context) error
	DeleteCategory(c echo.Context) error
}

type categoryController struct {
	categoryUsecase usecases.CategoryUsecase
}

func NewCategoryController(categoryUsecase usecases.CategoryUsecase) CategoryController {
	return &categoryController{categoryUsecase}
}

func (c *categoryController) GetAllCategories(ctx echo.Context) error {
	category, err := c.categoryUsecase.GetAllCategories()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, dtos.ErrorDTO{
			Message: err.Error(),
		})
	}

	var categoryDTOs []dtos.CategoryDTO

	for _, categories := range category {
		categoryDTO := dtos.CategoryDTO{
			ID:        categories.ID,
			Name:      categories.Name,
			CreatedAt: categories.CreatedAt,
			UpdatedAt: categories.UpdatedAt,
		}
		categoryDTOs = append(categoryDTOs, categoryDTO)
	}

	return ctx.JSON(http.StatusOK, dtos.CategoryDTOsResponse{
		Message: "Successfully get all categories.",
		Data:    categoryDTOs,
	})
}

func (c *categoryController) CreateCategory(ctx echo.Context) error {
	var categoryDTO dtos.CategoryDTO
	if err := ctx.Bind(&categoryDTO); err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	category := models.Category{
		Name: categoryDTO.Name,
	}

	err := c.categoryUsecase.CreateCategory(&category)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, dtos.ErrorDTO{
			Message: err.Error(),
		})
	}
	return ctx.JSON(http.StatusCreated, dtos.CategoryDTOResponse{
		Message: "Successfully created category.",
		Data: dtos.CategoryDTO{
			ID:        category.ID,
			Name:      category.Name,
			CreatedAt: category.CreatedAt,
			UpdatedAt: category.UpdatedAt,
		},
	})
}

func (c *categoryController) UpdateCategory(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.Param("id"))
	var categoryDTO dtos.CategoryDTO
	if err := ctx.Bind(&categoryDTO); err != nil {
		return ctx.JSON(http.StatusBadRequest, dtos.ErrorDTO{
			Message: err.Error(),
		})
	}

	category, err := c.categoryUsecase.GetCategoryByID(uint(id))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, dtos.ErrorDTO{
			Message: err.Error(),
		})
	}

	category.Name = categoryDTO.Name

	err = c.categoryUsecase.UpdateCategory(&category)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, dtos.ErrorDTO{
			Message: err.Error(),
		})
	}
	return ctx.JSON(http.StatusOK, dtos.CategoryDTOResponse{
		Message: "Successfully updated category.",
		Data: dtos.CategoryDTO{
			ID:        category.ID,
			Name:      category.Name,
			CreatedAt: category.CreatedAt,
			UpdatedAt: category.UpdatedAt,
		},
	})
}

func (c *categoryController) DeleteCategory(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.Param("id"))

	err := c.categoryUsecase.DeleteCategory(uint(id))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, dtos.ErrorDTO{
			Message: err.Error(),
		})
	}
	return ctx.JSON(http.StatusOK, dtos.ErrorDTO{
		Message: "Successfully deleted category.",
	})
}
