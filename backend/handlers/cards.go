package handlers

import (
	"fmt"
	"strconv"

	"git.bhasher.com/bhasher/focus/backend/db"
	"git.bhasher.com/bhasher/focus/backend/types"
	"github.com/gofiber/fiber/v2"
)

func cardsRouter(router fiber.Router) error {
	router.Post("/", CreateCard)
	router.Get("/:id", GetCard)
	router.Put("/:id", UpdateCard)
	router.Delete("/:id", DeleteCard)

	cardsTagsRouter(router.Group("/:card_id/tags"))
	return nil

}

func CreateCard(c *fiber.Ctx) error {
	card := types.Card{}
	if err := c.BodyParser(&card); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse request"})
	}

	id, err := db.CreateCard(card)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Cannot create card",
			"trace": fmt.Sprint(err),
		})
	}

	if err := c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"id": id,
	}); err != nil {
		return err
	}

	card.ID = id;

	source := c.Get("X-Request-Source");
	if source == "" {
		return nil;
	}

	publish(fiber.Map{
		"object": "card",
		"action": "create",
		"data": card,
		"X-Request-Source": source,
	})

	return nil;
}

func GetCard(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid card ID"})
	}

	card, err := db.GetCard(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Cannot retrieve card",
			"trace": fmt.Sprint(err),
		})
	}
	if card == nil {
		return c.SendStatus(fiber.StatusNotFound)
	}

	return c.Status(fiber.StatusOK).JSON(card)
}

func DeleteCard(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid card ID"})
	}

	count, err := db.DeleteCard(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Cannot delete card",
			"trace": fmt.Sprint(err),
		})
	}

	if count == 0 {
		return c.SendStatus(fiber.StatusNotFound)
	}

	if err := c.SendStatus(fiber.StatusNoContent); err != nil {
		return err
	}

	source := c.Get("X-Request-Source");
	if source == "" {
		return nil;
	}

	publish(fiber.Map{
		"object": "card",
		"action": "delete",
		"id": id,
		"X-Request-Source": source,
	});

	return nil;
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

	count, err := db.UpdateCard(card)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Cannot update card",
			"trace": fmt.Sprint(err),
		})
	}

	if count == 0 {
		return c.SendStatus(fiber.StatusNotFound)
	}

	if err := c.SendStatus(fiber.StatusNoContent); err != nil {
		return err
	}

	source := c.Get("X-Request-Source");
	if source == "" {
		return nil;
	}

	publish(fiber.Map{
		"object": "card",
		"action": "update",
		"id": id,
		"changes": card,
		"X-Request-Source": source,
	})

	return nil;
}
