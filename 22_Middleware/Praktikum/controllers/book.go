package controllers

import (
	"net/http"
	"praktikum/lib/database"
	"strconv"

	"github.com/labstack/echo"
)

// get all books
func GetBooksController(c echo.Context) error {
	books, err := database.GetBooks()

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "get all books successfully",
		"data":    books,
	})
}

// get book by id
func GetBookController(c echo.Context) error {
	// your solution here
	bookID, err := strconv.Atoi(c.Param("id"))

	book, err := database.GetBook(bookID)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "get book successfully",
		"data":    book,
	})
}

// create new book
func CreateBookController(c echo.Context) error {
	book, err := database.CreateBook(c)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "book created successfully",
		"data":    book,
	})
}

// delete book by id
func DeleteBookController(c echo.Context) error {
	// your solution here
	bookID, err := strconv.Atoi(c.Param("id"))

	book, err := database.DeleteBook(bookID)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "book deleted successfully",
		"data":    book,
	})
}

// update book by id
func UpdateBookController(c echo.Context) error {
	// your solution here
	book, err := database.UpdateBook(c)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "book updated successfully",
		"data":    book,
	})
}
