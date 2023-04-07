package usecases

import (
	"praktikum/dto"
	"praktikum/middlewares"
	"praktikum/models"
	"praktikum/repositories"
)

type UserUsecase interface {
	Create(input dto.CreateUserRequest) error
	GetAll() ([]models.User, error)
	LoginUser(input dto.CreateUserRequest) (interface{}, error)
}

type userUsecase struct {
	userRepository repositories.UserRepository
}

func NewUserUsecase(userRepo repositories.UserRepository) *userUsecase {
	return &userUsecase{userRepository: userRepo}
}

func (s *userUsecase) GetAll() ([]models.User, error) {

	users, err := s.userRepository.GetAll()

	if err != nil {
		return nil, err
	}
	return users, err
}

func (s *userUsecase) Create(input dto.CreateUserRequest) error {
	var userData models.User

	userData.Email = input.Email
	userData.Password = input.Password

	if err := s.userRepository.Create(userData); err != nil {
		return err
	}
	return nil
}

func (s *userUsecase) LoginUser(input dto.CreateUserRequest) (interface{}, error) {
	user, err := s.userRepository.LoginUser(input.Email, input.Password)

	if err != nil {
		return "nil", err
	}

	var userToken dto.UserToken
	userToken.Email = input.Email
	userToken.Password = input.Password
	userToken.Token, err = middlewares.CreateToken(int(user.ID))

	return userToken, nil
}
