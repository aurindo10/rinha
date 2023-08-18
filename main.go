package main

import (
	"rinha/db"
	handlers "rinha/handler"

	"github.com/gofiber/fiber/v2"
)



func main(){
	if err := db.ConnectToDb(); err != nil {
		panic(err)
	}
	app := fiber.New()
	handlers.PessoaHandler(app)

	if error := app.Listen(":4000"); error != nil {
		panic(error)
	}
	println("Server is running on port 3000")
}