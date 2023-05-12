package usecases

import (
	"code-competence/rest-api-jwt/helpers"
	"code-competence/rest-api-jwt/middlewares"
	"code-competence/rest-api-jwt/models"
	"code-competence/rest-api-jwt/repositories"
	"errors"
)

type UserUsecase interface {
	Login(input models.User) (string, error)
	Register(input models.User) (models.User, error)
}

type userUsecase struct {
	userRepo repositories.UserRepository
}

func NewUserUsecase(userRepo repositories.UserRepository) UserUsecase {
	return &userUsecase{userRepo}
}

func (u *userUsecase) Login(input models.User) (string, error) {
	var accessToken string

	user, _ := u.userRepo.GetUserByEmail(input.Email)
	if user.ID == 0 {
		return accessToken, errors.New("email/password is wrong")
	}

	valid := helpers.ComparePassword(input.Password, user.Password)
	if !valid {
		return accessToken, errors.New("email/password is wrong")
	}

	accessToken, err := middlewares.CreateToken(user.ID)
	if err != nil {
		return accessToken, err
	}

	return accessToken, nil
}

func (u *userUsecase) Register(users models.User) (models.User, error) {
	var user models.User

	user, _ = u.userRepo.GetUserByEmail(users.Email)
	if user.ID > 0 {
		return user, errors.New("Email already used")
	}

	password, err := helpers.HashPassword(users.Password)
	if err != nil {
		return user, err
	}

	if users.Email == "" || users.Name == "" || users.Password == "" {
		return user, errors.New("Failed to create user")
	}

	user.Name = users.Name
	user.Email = users.Email
	user.Password = password

	user, err = u.userRepo.CreateUser(user)
	if err != nil {
		return user, err
	}

	return user, nil
}
