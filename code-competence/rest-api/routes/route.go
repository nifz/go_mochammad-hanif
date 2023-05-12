package routes

import (
	"code-competence/rest-api/controllers"
	"code-competence/rest-api/repositories"
	"code-competence/rest-api/usecases"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Init(e *echo.Echo, db *gorm.DB) {
	itemRepo := repositories.NewItemRepository(db)
	itemUsecase := usecases.NewItemUsecase(itemRepo)
	itemController := controllers.NewItemController(itemUsecase)

	e.GET("/items", itemController.GetAllItems)
	e.GET("/items/:id", itemController.GetItemByID)
	e.POST("/items", itemController.CreateItem)
	e.PUT("/items/:id", itemController.UpdateItem)
	e.DELETE("/items/:id", itemController.DeleteItem)
	e.GET("/items/category/:category_id", itemController.GetItemsByCategoryID)
	e.GET("/items", itemController.GetItemsByName)

	categoryRepo := repositories.NewCategoryRepository(db)
	categoryUsecase := usecases.NewCategoryUsecase(categoryRepo)
	categoryController := controllers.NewCategoryController(categoryUsecase)

	e.GET("/categories", categoryController.GetAllCategories)
	e.POST("/categories", categoryController.CreateCategory)
	e.PUT("/categories/:id", categoryController.UpdateCategory)
	e.DELETE("/categories/:id", categoryController.DeleteCategory)
}
