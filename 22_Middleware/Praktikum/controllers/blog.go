package controllers

import (
	"net/http"
	"praktikum/lib/database"
	"strconv"

	"github.com/labstack/echo"
)

// get all blogs
func GetBlogsController(c echo.Context) error {
	blogs, err := database.GetBlogs()

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "get all blogs successfully",
		"data":    blogs,
	})
}

// get blog by id
func GetBlogController(c echo.Context) error {
	// your solution here
	blogID, err := strconv.Atoi(c.Param("id"))

	blog, err := database.GetBlog(blogID)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "get blog successfully",
		"data":    blog,
	})
}

// create new blog
func CreateBlogController(c echo.Context) error {
	blog, err := database.CreateBlog(c)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "blog created successfully",
		"data":    blog,
	})
}

// delete blog by id
func DeleteBlogController(c echo.Context) error {
	// your solution here
	blogID, err := strconv.Atoi(c.Param("id"))

	blog, err := database.DeleteBlog(blogID)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "blog deleted successfully",
		"data":    blog,
	})
}

// update blog by id
func UpdateBlogController(c echo.Context) error {
	// your solution here
	blog, err := database.UpdateBlog(c)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "blog updated successfully",
		"data":    blog,
	})
}
