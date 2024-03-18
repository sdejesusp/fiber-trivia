package handlers

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/sdejesusp/fiber-trivia/database"
	"github.com/sdejesusp/fiber-trivia/models"
)

func Welcome(c *fiber.Ctx) error {
	return c.SendString("Hello, from a Docker Fiber")
}

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

func UpdateFact(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	var fact models.Fact

	if err := FindFact(id, &fact); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	type UpdateFact struct {
		Question string `json:"question"`
		Answer   string `json:"answer"`
	}

	var updateData UpdateFact

	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	fact.Question = updateData.Question
	fact.Answer = updateData.Answer

	database.DB.Db.Save(&fact)
	return c.Status(fiber.StatusOK).JSON(fact)
}

func DeleteFact(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	var fact models.Fact
	if err := FindFact(id, &fact); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	if err := database.DB.Db.Delete(&fact).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(err.Error())
	}

	return c.Status(fiber.StatusOK).SendString("Record successfully deleted")

}
