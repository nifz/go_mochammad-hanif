package routes

import (
	"log"
	"os"
	"praktikum/controllers"
	"praktikum/repositories"
	"praktikum/usecases"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

func init() {
	middleware.ErrJWTMissing.Code = 401
	middleware.ErrJWTMissing.Message = "Unauthorized"
}

func New(e *echo.Echo, db *gorm.DB) {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	jwtMiddleware := middleware.JWT([]byte(os.Getenv("SECRET_JWT")))

	userRepository := repositories.NewUserRepository(db)

	userService := usecases.NewUserUsecase(userRepository)

	userController := controllers.NewUserController(userService)

	e.GET("/users", userController.GetAllUsers, jwtMiddleware)
	e.POST("/users", userController.CreateUser)
	e.POST("/users/login", userController.LoginUser)

}
