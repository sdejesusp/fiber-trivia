package handlers

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/sdejesusp/fiber-trivia/database"
	"github.com/sdejesusp/fiber-trivia/models"
)

func ListFacts(c *fiber.Ctx) error {
	facts := []models.Fact{}

	database.DB.Db.Find(&facts)

	return c.Status(200).JSON(facts)
}

func CreateFact(c *fiber.Ctx) error {
	fact := new(models.Fact)
	if err := c.BodyParser(fact); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	database.DB.Db.Create(&fact)

	return c.Status(fiber.StatusOK).JSON(fact)
}

func FindFact(id int, fact *models.Fact) error {
	database.DB.Db.Find(&fact, "id=?", id)
	if fact.ID == 0 {
		return errors.New("Record not found")
	}
	return nil
}

func GetSingleFact(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	var fact models.Fact
	if err := FindFact(id, &fact); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(fact)
}
