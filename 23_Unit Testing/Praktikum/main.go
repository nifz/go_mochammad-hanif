package main

import (
	"praktikum/configs"
	"praktikum/routes"
)

func main() {
	configs.InitDB()
	e := routes.New()
	e.Logger.Fatal(e.Start(":8080"))
}
