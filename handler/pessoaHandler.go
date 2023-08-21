package handlers

import (
	"rinha/controllers"

	"github.com/gofiber/fiber/v2"
)

func PessoaHandler(app *fiber.App)  {
	app.Post("/pessoas", controllers.CreatePessoa)
	app.Get("/pessoas/:id", controllers.GetPessoa)
	app.Get("pessoas", controllers.GetPessoaByTerm)
	app.Get("/contagem-pessoas", controllers.GetNumberOfPessoas)
} 		