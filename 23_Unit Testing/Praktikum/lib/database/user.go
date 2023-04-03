package database

import (
	"praktikum/configs"
	"praktikum/models"

	"github.com/labstack/echo"
)

func GetUsers() (interface{}, error) {
	var users []models.User

	if err := configs.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func CreateUser(c echo.Context) (interface{}, error) {
	var user models.User
	c.Bind(&user)
	if err := configs.DB.Save(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func GetUser(userID int) (interface{}, error) {
	// query the database for the user with the given ID
	var user models.User

	if err := configs.DB.Where("id = ?", userID).First(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func DeleteUser(userID int) (interface{}, error) {
	var user models.User
	// delete the user with the given ID from the database

	if err := configs.DB.Where("id = ?", userID).Delete(&user).Error; err != nil {
		return nil, err
	}

	return nil, nil
}

func UpdateUser(c echo.Context) (interface{}, error) {
	var user models.User

	if err := configs.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		return nil, err
	}
	if err := c.Bind(&user); err != nil {
		return nil, err
	}
	configs.DB.Save(&user)

	return user, nil
}
