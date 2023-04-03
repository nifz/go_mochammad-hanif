package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"praktikum/configs"
	"praktikum/models"

	"github.com/stretchr/testify/assert"
)

type BlogResponse struct {
	Message string        `json:"message"`
	Data    []models.User `json:"data"`
}

type DeleteBlogResponse struct {
	Message string `json:"message"`
}

func InsertDataBlogForGetBlogs() error {
	InsertDataUserForGetUsers()
	blog := models.Blog{
		UserID:      1,
		Title:       "ipsum dolor",
		Description: "lorem",
	}

	var err error
	if err = configs.DB.Save(&blog).Error; err != nil {
		return err
	}
	return nil
}

func TestGetBlogsController(t *testing.T) {
	var testCases = []struct {
		name       string
		path       string
		expectCode int
		sizeData   int
	}{
		{
			name:       "get blog normal",
			path:       "/blogs",
			expectCode: http.StatusOK,
			sizeData:   1,
		},
	}

	e := InitEchoTestAPI()
	InsertDataBlogForGetBlogs()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	for _, testCase := range testCases {
		c.SetPath(testCase.path)

		if assert.NoError(t, GetBlogsController(c)) {
			assert.Equal(t, testCase.expectCode, rec.Code)
			body := rec.Body.String()

			var blog BlogResponse
			err := json.Unmarshal([]byte(body), &blog)

			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, testCase.sizeData, len(blog.Data))
		}

	}

}

func TestGetBlogController(t *testing.T) {

	blog := models.Blog{
		UserID:      1,
		Title:       "ipsum dolor",
		Description: "lorem",
	}
	if err := configs.DB.Save(&blog).Error; err != nil {
		t.Errorf("Error setting up test blog: %v", err)
	}

	e := InitEchoTestAPI()
	InsertDataBlogForGetBlogs()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/blogs/:id")
	c.SetParamNames("id")
	c.SetParamValues(strconv.Itoa(1))

	if assert.NoError(t, GetBlogController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		var respBody struct {
			Message string      `json:"message"`
			Data    models.Blog `json:"data"`
		}
		err := json.Unmarshal(rec.Body.Bytes(), &respBody)
		if err != nil {
			t.Errorf("Error unmarshalling response body: %v", err)
		}
		assert.Equal(t, "get blog successfully", respBody.Message)
		assert.Equal(t, blog.Title, respBody.Data.Title)
		assert.Equal(t, blog.UserID, respBody.Data.UserID)
		assert.Equal(t, blog.Description, respBody.Data.Description)
	}

	if err := configs.DB.Delete(&blog).Error; err != nil {
		t.Errorf("Error cleaning up test blog: %v", err)
	}
}

func TestCreateBlogController(t *testing.T) {

	e := InitEchoTestAPI()

	blog := models.Blog{
		UserID: 1,
		User: models.User{
			Email:    "test@example.com",
			Password: "qwerty",
			Name:     "test",
		},
		Title:       "ipsum dolor",
		Description: "lorem",
	}
	if err := configs.DB.Save(&blog).Error; err != nil {
		t.Errorf("failed to create test blog: %v", err)
	}
	requestBodyBytes, err := json.Marshal(blog)
	if err != nil {
		t.Errorf("failed to marshal request body: %v", err)
	}

	req := httptest.NewRequest(http.MethodPost, "/blogs/", bytes.NewBuffer(requestBodyBytes))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, CreateBlogController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		var blogResponse models.Blog
		err = json.Unmarshal(requestBodyBytes, &blogResponse)
		if err != nil {
			t.Errorf("failed to unmarshal response body: %v", err)
		}
		assert.Equal(t, blog.Title, blogResponse.Title)
	}
}
func TestDeleteBlogController(t *testing.T) {
	e := InitEchoTestAPI()
	InsertDataBlogForGetBlogs()

	req := httptest.NewRequest(http.MethodDelete, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	if assert.NoError(t, DeleteBlogController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		body := rec.Body.String()

		var result DeleteUserResponse
		err := json.Unmarshal([]byte(body), &result)

		if err != nil {
			assert.Error(t, err, "error")
		}

		assert.Equal(t, "blog deleted successfully", result.Message)
	}
}

func TestUpdateBlogController(t *testing.T) {

	e := InitEchoTestAPI()

	blog := models.Blog{
		UserID: 1,
		User: models.User{
			Email:    "test@example.com",
			Password: "qwerty",
			Name:     "test",
		},
		Title:       "ipsum dolor",
		Description: "lorem",
	}
	if err := configs.DB.Save(&blog).Error; err != nil {
		t.Errorf("failed to create test blog: %v", err)
	}

	requestBody := map[string]interface{}{
		"title": "Jane Doe",
	}
	requestBodyBytes, err := json.Marshal(requestBody)
	if err != nil {
		t.Errorf("failed to marshal request body: %v", err)
	}

	req := httptest.NewRequest(http.MethodPut, "/blogs/"+strconv.Itoa(int(blog.ID)), bytes.NewBuffer(requestBodyBytes))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(strconv.Itoa(int(blog.ID)))

	if assert.NoError(t, UpdateBlogController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		var blogResponse models.Blog
		err = json.Unmarshal(requestBodyBytes, &blogResponse)
		if err != nil {
			t.Errorf("failed to unmarshal response body: %v", err)
		}
		assert.Equal(t, requestBody["title"], blogResponse.Title)
	}
}
