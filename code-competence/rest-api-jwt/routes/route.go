package routes

import (
	"code-competence/rest-api-jwt/controllers"
	"code-competence/rest-api-jwt/repositories"
	"code-competence/rest-api-jwt/usecases"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

func init() {
	middleware.ErrJWTMissing.Code = 401
	middleware.ErrJWTMissing.Message = "Unauthorized"
}

func Init(e *echo.Echo, db *gorm.DB) {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	jwtMiddleware := middleware.JWT([]byte(os.Getenv("SECRET_JWT")))

	userRepository := repositories.NewUserRepository(db)
	userUsecase := usecases.NewUserUsecase(userRepository)
	userController := controllers.NewUserController(userUsecase)

	e.POST("/users/login", userController.SignIn)
	e.POST("/users/register", userController.SignUp)

	itemRepo := repositories.NewItemRepository(db)
	itemUsecase := usecases.NewItemUsecase(itemRepo)
	itemController := controllers.NewItemController(itemUsecase)

	e.GET("/items", itemController.GetAllItems, jwtMiddleware)
	e.GET("/items/:id", itemController.GetItemByID, jwtMiddleware)
	e.POST("/items", itemController.CreateItem, jwtMiddleware)
	e.PUT("/items/:id", itemController.UpdateItem, jwtMiddleware)
	e.DELETE("/items/:id", itemController.DeleteItem, jwtMiddleware)
	e.GET("/items/category/:category_id", itemController.GetItemsByCategoryID, jwtMiddleware)
	e.GET("/items", itemController.GetItemsByName, jwtMiddleware)

	categoryRepo := repositories.NewCategoryRepository(db)
	categoryUsecase := usecases.NewCategoryUsecase(categoryRepo)
	categoryController := controllers.NewCategoryController(categoryUsecase)

	e.GET("/categories", categoryController.GetAllCategories, jwtMiddleware)
	e.POST("/categories", categoryController.CreateCategory, jwtMiddleware)
	e.PUT("/categories/:id", categoryController.UpdateCategory, jwtMiddleware)
	e.DELETE("/categories/:id", categoryController.DeleteCategory, jwtMiddleware)
}
