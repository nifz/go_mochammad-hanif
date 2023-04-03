package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"praktikum/configs"
	"praktikum/models"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

type UserResponse struct {
	Message string        `json:"message"`
	Data    []models.User `json:"data"`
}

type DeleteUserResponse struct {
	Message string `json:"message"`
}

func InitEchoTestAPI() *echo.Echo {
	configs.InitDBTest()
	e := echo.New()
	return e
}

func InsertDataUserForGetUsers() error {
	user := models.User{
		Name:     "Alta",
		Password: "123",
		Email:    "alta@gmail.com",
	}

	var err error
	if err = configs.DB.Save(&user).Error; err != nil {
		return err
	}
	return nil
}

func TestGetUsersController(t *testing.T) {
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

	for _, testCase := range testCases {
		c.SetPath(testCase.path)

		if assert.NoError(t, GetUsersController(c)) {
			assert.Equal(t, testCase.expectCode, rec.Code)
			body := rec.Body.String()

			var user UserResponse
			err := json.Unmarshal([]byte(body), &user)

			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, testCase.sizeData, len(user.Data))
		}

	}

}

func TestGetUserController(t *testing.T) {

	user := models.User{
		Name:     "Alta",
		Password: "123",
		Email:    "alta@gmail.com",
	}
	if err := configs.DB.Save(&user).Error; err != nil {
		t.Errorf("Error setting up test user: %v", err)
	}

	e := InitEchoTestAPI()
	InsertDataUserForGetUsers()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/users/:id")
	c.SetParamNames("id")
	c.SetParamValues(strconv.Itoa(1))

	if assert.NoError(t, GetUserController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		var respBody struct {
			Message string      `json:"message"`
			Data    models.User `json:"data"`
		}
		err := json.Unmarshal(rec.Body.Bytes(), &respBody)
		if err != nil {
			t.Errorf("Error unmarshalling response body: %v", err)
		}
		assert.Equal(t, "get user successfully", respBody.Message)
		assert.Equal(t, user.Name, respBody.Data.Name)
		assert.Equal(t, user.Email, respBody.Data.Email)
		assert.Equal(t, user.Password, respBody.Data.Password)
	}

	if err := configs.DB.Delete(&user).Error; err != nil {
		t.Errorf("Error cleaning up test user: %v", err)
	}
}
func TestCreateUserController(t *testing.T) {
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

	e := InitEchoTestAPI()
	for _, testCase := range testCases {
		req := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(testCase.request))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		if assert.NoError(t, CreateUserController(c)) {
			assert.Equal(t, testCase.expectCode, rec.Code)

			if testCase.expectCode == http.StatusCreated {
				body := rec.Body.String()
				var user models.User
				err := json.Unmarshal([]byte(body), &user)

				if err != nil {
					assert.Error(t, err, "error")
				}
				assert.Equal(t, testCase.request[0], user.Name)
				assert.Equal(t, testCase.request[1], user.Email)
			}
		}
	}
}
func TestDeleteUserController(t *testing.T) {
	e := InitEchoTestAPI()
	InsertDataUserForGetUsers()

	req := httptest.NewRequest(http.MethodDelete, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	if assert.NoError(t, DeleteUserController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		body := rec.Body.String()

		var result DeleteUserResponse
		err := json.Unmarshal([]byte(body), &result)

		if err != nil {
			assert.Error(t, err, "error")
		}

		assert.Equal(t, "user deleted successfully", result.Message)
	}
}
func TestUpdateUserController(t *testing.T) {
	e := InitEchoTestAPI()

	user := models.User{
		Name:     "Alta",
		Password: "123",
		Email:    "alta@gmail.com",
	}
	if err := configs.DB.Save(&user).Error; err != nil {
		t.Errorf("failed to create test user: %v", err)
	}

	requestBody := map[string]interface{}{
		"name": "Jane Doe",
	}
	requestBodyBytes, err := json.Marshal(requestBody)
	if err != nil {
		t.Errorf("failed to marshal request body: %v", err)
	}

	req := httptest.NewRequest(http.MethodPut, "/users/"+strconv.Itoa(int(user.ID)), bytes.NewBuffer(requestBodyBytes))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(strconv.Itoa(int(user.ID)))

	if assert.NoError(t, UpdateUserController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		var userResponse models.User
		err = json.Unmarshal(requestBodyBytes, &userResponse)
		if err != nil {
			t.Errorf("failed to unmarshal response body: %v", err)
		}
		assert.Equal(t, requestBody["name"], userResponse.Name)
	}
}
