package handlers

import (
	"fmt"
	"strconv"

	"git.bhasher.com/bhasher/focus/backend/db"
	"git.bhasher.com/bhasher/focus/backend/types"
	"git.bhasher.com/bhasher/focus/backend/utils"
	"github.com/gofiber/fiber/v2"
)

func CreateCard(c *fiber.Ctx) error {
	card := types.Card{}
	if err := c.BodyParser(&card); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse request",
			"trace": fmt.Sprint(err),
		})
	}

	id, err := db.CreateCard(card)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Cannot create card",
			"trace": fmt.Sprint(err),
		})
	}

	c.Status(fiber.StatusCreated)
	c.Location(fmt.Sprintf("/api/cards/%v", id))
	return c.JSON(fiber.Map{
		"id":     id,
		"_links": utils.HALProjectLinks(id),
	})
}

func GetAllCardsOf(c *fiber.Ctx) error {
	projectID, err := strconv.Atoi(c.Params("project_id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "error": "Invalid project_id", "trace": fmt.Sprint(err)})
	}

	projects, err := db.GetAllCardsOf(projectID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "error": "Cannot retrieve cards", "trace": fmt.Sprint(err)})
	}

	return c.JSON(fiber.Map{"status": "ok", "data": projects})
}

func GetCard(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid card ID"})
	}

	card, err := db.GetCard(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Cannot retrieve card"})
	}
	if card == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Card not found"})
	}

	return c.JSON(card)
}

func DeleteCard(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "error": "Invalid card ID", "trace": fmt.Sprint(err)})
	}

	err = db.DeleteCard(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "error": "Cannot delete card", "trace": fmt.Sprint(err)})
	}

	return c.SendStatus(fiber.StatusNoContent)
}

func UpdateCard(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid card ID"})
	}

	card := types.Card{ID: id}
	if err := c.BodyParser(&card); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse request"})
	}

	err = db.UpdateCard(card)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Cannot update card"})
	}

	return c.SendStatus(fiber.StatusOK)
}
