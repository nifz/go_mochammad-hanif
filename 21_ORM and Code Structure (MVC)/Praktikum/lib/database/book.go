package database

import (
	"praktikum/configs"
	"praktikum/models"

	"github.com/labstack/echo"
)

func GetBooks() (interface{}, error) {
	var books []models.Book

	if err := configs.DB.Find(&books).Error; err != nil {
		return nil, err
	}
	return books, nil
}

func CreateBook(c echo.Context) (interface{}, error) {
	var book models.Book
	c.Bind(&book)

	if err := configs.DB.Save(&book).Error; err != nil {
		return nil, err
	}
	return book, nil
}

func GetBook(bookID int) (interface{}, error) {
	// query the database for the book with the given ID
	var book models.Book

	if err := configs.DB.Where("id = ?", bookID).First(&book).Error; err != nil {
		return nil, err
	}

	return book, nil
}

func DeleteBook(bookID int) (interface{}, error) {
	var book models.Book
	// delete the book with the given ID from the database

	if err := configs.DB.Where("id = ?", bookID).Delete(&book).Error; err != nil {
		return nil, err
	}

	return "Successfully deleted", nil
}

func UpdateBook(c echo.Context) (interface{}, error) {
	var book models.Book

	if err := configs.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		return nil, err
	}
	if err := c.Bind(&book); err != nil {
		return nil, err
	}
	configs.DB.Save(&book)

	return book, nil
}
