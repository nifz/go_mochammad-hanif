package controllers

import (
	"code-competence/rest-api-jwt/dtos"
	"code-competence/rest-api-jwt/models"
	"code-competence/rest-api-jwt/usecases"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	userUsecase usecases.UserUsecase
}

func NewUserController(userUsecase usecases.UserUsecase) UserController {
	return UserController{userUsecase}
}

func (c *UserController) SignIn(ctx echo.Context) error {
	var user models.User

	err := ctx.Bind(&user)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, dtos.ErrorDTO{
			Message: err.Error(),
		})
	}

	tokenString, err := c.userUsecase.Login(user)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, dtos.ErrorDTO{
			Message: err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, dtos.LoginDTOResponse{
		Message: "Successfully logged in.",
		Data: dtos.UserLoginDTO{
			Email: user.Email,
			Token: tokenString,
		},
	})
}

func (c *UserController) SignUp(ctx echo.Context) error {
	var user models.User
	err := ctx.Bind(&user)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, dtos.ErrorDTO{
			Message: err.Error(),
		})
	}

	user, err = c.userUsecase.Register(user)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, dtos.ErrorDTO{
			Message: err.Error(),
		})
	}

	return ctx.JSON(http.StatusCreated, dtos.RegisterDTOResponse{
		Message: "Successfully registered",
		Data: dtos.UserDTO{
			Name:      user.Name,
			Email:     user.Email,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		},
	})
}
