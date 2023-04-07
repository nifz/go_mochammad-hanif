package main

import (
	"praktikum/configs"
	"praktikum/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	db, err := configs.ConnectDB()
	if err != nil {
		panic(err)
	}

	err = configs.MigrateDB(db)
	if err != nil {
		panic(err)
	}

	app := echo.New()
	routes.New(app, db)

	app.Start(":8000")
}
