package main

import (
	"rinha/db"

	"github.com/gofiber/fiber/v2"
)



func main(){
	if err := db.ConnectToDb(); err != nil {
		panic(err)
	}
	app := fiber.New()

	if error := app.Listen(":3000"); error != nil {
		panic(error)
	}
	println("Server is running on port 3000")
}