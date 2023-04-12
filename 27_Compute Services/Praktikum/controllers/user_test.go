package controllers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"praktikum/configs"
	"praktikum/dto"
	"praktikum/models"
	"praktikum/repositories"
	"praktikum/usecases"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func InitEchoTestAPI() *echo.Echo {
	configs.InitDBTest()
	e := echo.New()
	return e
}

func InsertDataUserForGetUsers() error {
	user := models.User{
		Email:    "alta@gmail.com",
		Password: "123",
	}

	var err error
	if err = configs.DB.Save(&user).Error; err != nil {
		return err
	}
	return nil
}

func TestGetAllUsers(t *testing.T) {
	var testCases = []struct {
		name       string
		path       string
		expectCode int
		sizeData   int
	}{
		{
			name:       "get user normal",
			path:       "/users",
			expectCode: http.StatusOK,
			sizeData:   1,
		},
	}

	e := InitEchoTestAPI()
	InsertDataUserForGetUsers()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	userRepository := repositories.NewUserRepository(configs.DB)

	userService := usecases.NewUserUsecase(userRepository)

	userController := NewUserController(userService)

	for _, testCase := range testCases {
		c.SetPath(testCase.path)

		if assert.NoError(t, userController.GetAllUsers(c)) {
			assert.Equal(t, testCase.expectCode, rec.Code)
			body := rec.Body.String()

			var user dto.UserResponse
			err := json.Unmarshal([]byte(body), &user)

			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, testCase.sizeData, len(user.Data))
		}
	}

}
func TestCreateUser(t *testing.T) {
	var testCases = []struct {
		name       string
		request    string
		expectCode int
	}{
		{
			name:       "create user success",
			request:    `{"name":"Test","email":"test@test.com","password":"123"}`,
			expectCode: http.StatusOK,
		},
	}

	userRepository := repositories.NewUserRepository(configs.DB)

	userService := usecases.NewUserUsecase(userRepository)

	userController := NewUserController(userService)

	e := InitEchoTestAPI()
	for _, testCase := range testCases {
		req := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(testCase.request))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		if assert.NoError(t, userController.CreateUser(c)) {
			assert.Equal(t, testCase.expectCode, rec.Code)

			if testCase.expectCode == http.StatusCreated {
				body := rec.Body.String()
				var user models.User
				err := json.Unmarshal([]byte(body), &user)

				if err != nil {
					assert.Error(t, err, "error")
				}
				assert.Equal(t, testCase.request[0], user.Email)
				assert.Equal(t, testCase.request[1], user.Password)
			}
		}
	}
}
func TestLoginUser(t *testing.T) {
	var testCases = []struct {
		name       string
		request    string
		expectCode int
		sizeData   int
	}{
		{
			name:       "login user success",
			request:    `{"email":"test@test.com","password":"123"}`,
			expectCode: http.StatusOK,
			sizeData:   1,
		},
	}

	userRepository := repositories.NewUserRepository(configs.DB)

	userService := usecases.NewUserUsecase(userRepository)

	userController := NewUserController(userService)

	e := InitEchoTestAPI()
	for _, testCase := range testCases {
		req := httptest.NewRequest(http.MethodPost, "/users/login", strings.NewReader(testCase.request))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		if assert.NoError(t, userController.LoginUser(c)) {
			// assert.Equal(t, testCase.expectCode, rec.Code)

			if testCase.expectCode == http.StatusCreated {
				body := rec.Body.String()
				var user models.User
				err := json.Unmarshal([]byte(body), &user)

				if err != nil {
					assert.Error(t, err, "error")
				}
				assert.Equal(t, testCase.sizeData, len(user.Email))
			}
		}
	}
}
