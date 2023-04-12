package controllers

import (
	"praktikum/dto"
	"praktikum/usecases"

	"github.com/labstack/echo/v4"
)

type userController struct {
	userUsecase usecases.UserUsecase
}

func NewUserController(userUsecase usecases.UserUsecase) *userController {
	return &userController{userUsecase}
}

func (u *userController) GetAllUsers(c echo.Context) error {
	users, _ := u.userUsecase.GetAll()

	return c.JSON(200, echo.Map{
		"data": users,
	})
}

func (u *userController) CreateUser(c echo.Context) error {
	user := dto.CreateUserRequest{}

	if err := c.Bind(&user); err != nil {
		return c.JSON(400, echo.Map{
			"error": err.Error(),
		})
	}

	u.userUsecase.Create(user)

	return c.JSON(200, echo.Map{
		"data": user,
	})
}

func (u *userController) LoginUser(c echo.Context) error {
	var user dto.CreateUserRequest
	if err := c.Bind(&user); err != nil {
		return c.JSON(400, echo.Map{
			"error": err.Error(),
		})
	}
	users, _ := u.userUsecase.LoginUser(user)

	return c.JSON(200, echo.Map{
		"data": users,
	})
}
