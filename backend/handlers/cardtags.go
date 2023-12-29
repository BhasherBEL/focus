package handlers

import (
	"fmt"
	"strconv"

	"git.bhasher.com/bhasher/focus/backend/db"
	"git.bhasher.com/bhasher/focus/backend/types"
	"github.com/gofiber/fiber/v2"
)

func CreateTagOfCard(c *fiber.Ctx) error {
	cardtag := types.CardTag{}
	if err := c.BodyParser(&cardtag); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse request"})
	}

	err := db.AddTagToCard(cardtag)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Cannot add tag to card"})
	}

	return c.SendStatus(fiber.StatusOK)
}

func GetAllTagsOfCard(c *fiber.Ctx) error {
	cardID, err := strconv.Atoi(c.Params("card_id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "error": "Invalid card_id"})
	}

	cardtags, err := db.GetAllTagsOfCard(cardID, -1)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "error": "Cannot retrieve tags of card", "stack": fmt.Sprint(err)})
	}

	return c.JSON(fiber.Map{"status": "ok", "data": cardtags})
}

func DeleteTagOfCard(c *fiber.Ctx) error {
	card_id, err := strconv.Atoi(c.Params("card_id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid card ID"})
	}

	tag_id, err := strconv.Atoi(c.Params("tag_id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid tag ID"})
	}

	err = db.DeleteTagOfCard(card_id, tag_id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Cannot delete tag of card"})
	}

	return c.SendStatus(fiber.StatusNoContent)
}

func DeleteTagsOfCard(c *fiber.Ctx) error {
	card_id, err := strconv.Atoi(c.Params("card_id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid card ID"})
	}

	err = db.DeleteTagsOfCard(card_id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Cannot delete tags of card"})
	}

	return c.SendStatus(fiber.StatusNoContent)
}

func UpdateTagOfCard(c *fiber.Ctx) error {
	card_id, err := strconv.Atoi(c.Params("card_id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid card ID"})
	}

	tag_id, err := strconv.Atoi(c.Params("tag_id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "error": "Invalid tag ID", "stack": fmt.Sprint(err)})
	}

	cardtag := types.CardTag{CardID: card_id, TagID: tag_id}
	if err := c.BodyParser(&cardtag); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse request"})
	}

	err = db.UpdateTagOfCard(cardtag)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Cannot update tag of card"})
	}

	return c.SendStatus(fiber.StatusOK)
}
