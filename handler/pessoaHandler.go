package handlers

import (
	"rinha/controllers"

	"github.com/gofiber/fiber/v2"
)

func PessoaHandler(app *fiber.App)  {
	app.Post("/pessoas", controllers.CreatePessoa)
} 	