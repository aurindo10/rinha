package controllers

import (
	"fmt"
	"rinha/db"
	"rinha/db/schemas"

	"github.com/gofiber/fiber/v2"
)


func CreatePessoa (c *fiber.Ctx) error {
	pessoa := &schemas.Pessoas{}
	if err := c.BodyParser(pessoa); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"message": "Cannot parse JSON",
		})
	}
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