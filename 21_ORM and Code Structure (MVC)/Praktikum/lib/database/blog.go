package database

import (
	"praktikum/configs"
	"praktikum/models"

	"github.com/labstack/echo"
)

func GetBlogs() (interface{}, error) {
	var blogs []models.Blog

	if err := configs.DB.Find(&blogs).Error; err != nil {
		return nil, err
	}
	return blogs, nil
}

func CreateBlog(c echo.Context) (interface{}, error) {
	var blog models.Blog
	c.Bind(&blog)

	if err := configs.DB.Save(&blog).Error; err != nil {
		return nil, err
	}
	return blog, nil
}

func GetBlog(blogID int) (interface{}, error) {
	// query the database for the book with the given ID
	var blog models.Blog

	if err := configs.DB.Where("id = ?", blogID).First(&blog).Error; err != nil {
		return nil, err
	}

	return blog, nil
}

func DeleteBlog(blogID int) (interface{}, error) {
	var blog models.Blog
	// delete the blog with the given ID from the database

	if err := configs.DB.Where("id = ?", blogID).Delete(&blog).Error; err != nil {
		return nil, err
	}

	return "Successfully deleted", nil
}

func UpdateBlog(c echo.Context) (interface{}, error) {
	var blog models.Blog

	if err := configs.DB.Where("id = ?", c.Param("id")).First(&blog).Error; err != nil {
		return nil, err
	}
	if err := c.Bind(&blog); err != nil {
		return nil, err
	}
	configs.DB.Save(&blog)

	return blog, nil
}
