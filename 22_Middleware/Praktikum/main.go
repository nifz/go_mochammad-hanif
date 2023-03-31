package main

import (
	"praktikum/configs"
	"praktikum/middlewares"
	"praktikum/routes"
)

func main() {
	configs.InitDB()
	e := routes.New()
	middlewares.LogMiddleware(e)
	e.Logger.Fatal(e.Start(":8080"))
}
