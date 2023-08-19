package controllers

import (
	"fmt"
	"rinha/db"
	"rinha/db/schemas"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/lib/pq"
)

type PessoasRequest struct {
	Apelido    string       `gorm:"type:varchar(32)"`
	Nome       string       `gorm:"type:varchar(100)"`
	Nascimento string 		
	Stack      pq.StringArray  `gorm:"type:varchar(32)[]"`
}
func CreatePessoa (c *fiber.Ctx) error {
	pessoaRequet := &PessoasRequest{}
	if err := c.BodyParser(pessoaRequet); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"message": "Cannot parse JSON",
		})
	}
	if pessoaRequet.Nascimento == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"message": "Data não pode ser vazio",
		})
	}
	format := "2006-01-02"
	date, _ := time.Parse(format, pessoaRequet.Nascimento)
	pessoa := &schemas.Pessoas{Nascimento: date, Apelido: pessoaRequet.Apelido, Nome: pessoaRequet.Nome, Stack: pessoaRequet.Stack}
	if result := db.DB.Create(pessoa); result.Error != nil {
		if result.Error.Error() == "422" {
			return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
				"error": true,
				"message": result.Error.Error(),
			})
		} else {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": true,
				"message": result.Error.Error(),
			})
		}
	}
	location := fmt.Sprintf("/pessoas/%s", pessoa.ID)
	c.Set("Location", location)
	return c.Status(fiber.StatusCreated).JSON(pessoa)
}