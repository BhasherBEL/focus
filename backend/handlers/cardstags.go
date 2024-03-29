package handlers

import (
	"fmt"
	"strconv"
	"strings"

	"git.bhasher.com/bhasher/focus/backend/db"
	"git.bhasher.com/bhasher/focus/backend/types"
	"github.com/gofiber/fiber/v2"
)

func cardsTagsRouter(router fiber.Router) error {
	router.Get("/", GetCardTags)
	router.Delete("/", DeleteCardTags)
	router.Post("/:tag_id", CreateCardTag)
	router.Put("/:tag_id", UpdateCardTag)
	router.Delete("/:tag_id", DeleteCardTag)
	return nil
}

func CreateCardTag(c *fiber.Ctx) error {
	cardID, err := strconv.Atoi(c.Params("card_id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid card_id"})
	}
	tagID, err := strconv.Atoi(c.Params("tag_id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid tag_id"})
	}

	cardtag := types.CardTag{CardID: cardID, TagID: tagID}
	if err := c.BodyParser(&cardtag); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse request"})
	}

	exist, err := db.ExistCard(cardID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error finding card",
			"trace": fmt.Sprint(err),
		})
	}
	if !exist {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Card not found"})
	}

	exist, err = db.ExistTag(tagID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error finding tag",
			"trace": fmt.Sprint(err),
		})
	}
	if !exist {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Tag not found"})
	}

	err = db.CreateCardTag(cardtag)
	if err != nil {
		if strings.HasPrefix(err.Error(), "UNIQUE constraint failed") {
			return c.SendStatus(fiber.StatusConflict)
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Cannot add tag to card",
			"trace": fmt.Sprint(err),
		})
	}

	err = c.SendStatus(fiber.StatusCreated)
	if err != nil {
		return err
	}

	source := c.Get("X-Request-Source");
	if source == "" {
		return nil;
	}

	publish(fiber.Map{
		"object": "cardTag",
		"action": "create",
		"card_id": cardID,
		"data": cardtag,
		"X-Request-Source": source,
	});

	return nil
}

func GetCardTags(c *fiber.Ctx) error {
	cardID, err := strconv.Atoi(c.Params("card_id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid card_id"})
	}

	exist, err := db.ExistCard(cardID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error finding card",
			"trace": fmt.Sprint(err),
		})
	}
	if !exist {
		return c.SendStatus(fiber.StatusNotFound)
	}

	cardtags, err := db.GetCardTags(cardID, -1)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Cannot retrieve tags of card",
			"stack": fmt.Sprint(err),
		})
	}

	return c.Status(fiber.StatusOK).JSON(cardtags)
}

func DeleteCardTag(c *fiber.Ctx) error {
	cardID, err := strconv.Atoi(c.Params("card_id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid card ID"})
	}

	tagID, err := strconv.Atoi(c.Params("tag_id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid tag ID"})
	}

	exist, err := db.ExistCard(cardID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error finding card",
			"trace": fmt.Sprint(err),
		})
	}
	if !exist {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Card not found"})
	}

	exist, err = db.ExistTag(tagID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error finding tag",
			"trace": fmt.Sprint(err),
		})
	}
	if !exist {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Tag not found"})
	}

	count, err := db.DeleteCardTag(cardID, tagID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Cannot delete tag of card",
			"stack": fmt.Sprint(err),
		})
	}

	if count == 0 {
		return c.SendStatus(fiber.StatusNotFound)
	}

	err = c.SendStatus(fiber.StatusNoContent)

	if err != nil {	
		return err
	}

	source := c.Get("X-Request-Source");
	if source == "" {
		return nil		
	}

	publish(fiber.Map{
		"object": "cardTag",
		"action": "delete",
		"card_id": cardID,
		"tag_id": tagID,
		"X-Request-Source": source,
	});

	return nil
}

func DeleteCardTags(c *fiber.Ctx) error {
	cardID, err := strconv.Atoi(c.Params("card_id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid card ID"})
	}

	exist, err := db.ExistCard(cardID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error finding card",
			"trace": fmt.Sprint(err),
		})
	}
	if !exist {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Card not found"})
	}

	_, err = db.DeleteCardTags(cardID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Cannot delete tags of card",
			"stack": fmt.Sprint(err),
		})
	}

	return c.SendStatus(fiber.StatusNoContent)

	// TODO publish event
}

func UpdateCardTag(c *fiber.Ctx) error {
	cardID, err := strconv.Atoi(c.Params("card_id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid card ID"})
	}

	tagID, err := strconv.Atoi(c.Params("tag_id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid tag ID"})
	}

	cardtag := types.CardTag{CardID: cardID, TagID: tagID}
	if err := c.BodyParser(&cardtag); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse request"})
	}

	exist, err := db.ExistCard(cardID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error finding card",
			"trace": fmt.Sprint(err),
		})
	}
	if !exist {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Card not found"})
	}

	exist, err = db.ExistTag(tagID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error finding tag",
			"trace": fmt.Sprint(err),
		})
	}
	if !exist {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Tag not found"})
	}

	count, err := db.UpdateCardTag(cardtag)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Cannot update tag of card",
			"stack": fmt.Sprint(err),
		})
	}

	if count == 0 {
		return c.SendStatus(fiber.StatusNotFound)
	}

	err = c.SendStatus(fiber.StatusNoContent)
	if err != nil {
		return err
	}

	source := c.Get("X-Request-Source");
	if source == "" {
		return nil
	}

	publish(fiber.Map{
		"object": "cardTag",
		"action": "update",
		"card_id": cardID,
		"tag_id": tagID,
		"changes": cardtag,
		"X-Request-Source": source,
	})

	return nil
}
