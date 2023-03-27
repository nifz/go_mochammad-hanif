package main

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
)

// -------------------- controller --------------------

// get movie by id
func GetMovieByID(c echo.Context) error {
	imdbID := c.Param("imdbID")
	resp, _ := http.Get("https://www.omdbapi.com/?apikey=a818034f&i=" + imdbID)

	defer resp.Body.Close()

	var movie map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&movie); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Failed to decode movie data",
		})
	}

	return c.JSON(http.StatusOK, movie)

}

// get movie with query parameters
func GetMovieWithQueryParam(c echo.Context) error {

	var query string

	if len(c.QueryParam("page")) > 0 {
		query += "&page=" + c.QueryParam("page")
	}

	if len(c.QueryParam("search")) > 0 {
		query += "&s=" + c.QueryParam("search")
	}

	if len(c.QueryParam("type")) > 0 {
		query += "&type=" + c.QueryParam("type")
	}

	resp, _ := http.Get("https://www.omdbapi.com/?apikey=a818034f" + query)

	defer resp.Body.Close()

	var movie map[string]interface{}

	if err := json.NewDecoder(resp.Body).Decode(&movie); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Failed to decode movie data",
		})
	}

	return c.JSON(http.StatusOK, movie)
}

// ---------------------------------------------------
func main() {
	e := echo.New()
	// routing with query parameter
	e.GET("/:imdbID", GetMovieByID)
	e.GET("/", GetMovieWithQueryParam)

	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":8000"))
}
