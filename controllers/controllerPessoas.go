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
type PessoaResponse struct {
	ID      string `json:"id"`
	Apelido string `json:"apelido"`
	Nome    string `json:"nome"`
	Nascimento string `json:"nascimento"`
	Stack   []string `json:"stack"`
}
func CreatePessoa(c *fiber.Ctx) error {
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

func GetPessoa(c *fiber.Ctx)error {
	pessoa := &schemas.Pessoas{}
	id := c.Params("id")
	if len(id) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"message": "Id não pode ser vazio",
		})
	}
	if result := db.DB.First(pessoa, "id = ?", id); result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": true,
			"message": result.Error.Error(),
		})
	}
	pessoaResp := PessoaResponse{
		ID:        pessoa.ID,
		Apelido:   pessoa.Apelido,
		Nome:      pessoa.Nome,
		Nascimento: pessoa.Nascimento.Format("2006-01-02"),
		Stack:     pessoa.Stack,
	}
	return c.Status(fiber.StatusOK).JSON(pessoaResp)

}