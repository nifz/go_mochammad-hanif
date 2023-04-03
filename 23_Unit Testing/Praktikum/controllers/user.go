package controllers

import (
	"net/http"
	"praktikum/lib/database"
	"strconv"

	"github.com/labstack/echo"
)

// get all users
func GetUsersController(c echo.Context) error {
	users, _ := database.GetUsers()

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "get all users successfully",
		"data":    users,
	})
}

// get user by id
func GetUserController(c echo.Context) error {
	// your solution here
	userID, _ := strconv.Atoi(c.Param("id"))

	user, _ := database.GetUser(userID)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "get user successfully",
		"data":    user,
	})
}

// create new user
func CreateUserController(c echo.Context) error {
	user, _ := database.CreateUser(c)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "user created successfully",
		"data":    user,
	})
}

// delete user by id
func DeleteUserController(c echo.Context) error {
	// your solution here
	userID, _ := strconv.Atoi(c.Param("id"))

	user, _ := database.DeleteUser(userID)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "user deleted successfully",
		"data":    user,
	})
}

// update user by id
func UpdateUserController(c echo.Context) error {
	// your solution here
	user, _ := database.UpdateUser(c)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "user updated successfully",
		"data":    user,
	})
}
