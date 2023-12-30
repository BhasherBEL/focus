package handlers

import (
	"fmt"
	"strconv"

	"git.bhasher.com/bhasher/focus/backend/db"
	"git.bhasher.com/bhasher/focus/backend/types"
	"github.com/gofiber/fiber/v2"
)

func tagsRouter(router fiber.Router) {
	router.Post("/", CreateTag)
	router.Get("/:id", GetTag)
	router.Delete("/:id", DeleteTag)
	router.Put("/:id", UpdateTag)
}

func CreateTag(c *fiber.Ctx) error {
	tag := types.Tag{}
	if err := c.BodyParser(&tag); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse request"})
	}

	id, err := db.CreateTag(tag)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Cannot create tag",
			"trace": fmt.Sprint(err),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"id": id})
}

func GetTag(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid tag ID"})
	}

	tag, err := db.GetTag(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Cannot retrieve tag",
			"trace": fmt.Sprint(err),
		})
	}
	if tag == nil {
		return c.SendStatus(fiber.StatusNotFound)
	}

	return c.JSON(tag)
}

func DeleteTag(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid tag ID"})
	}

	count, err := db.DeleteTag(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Cannot delete tag",
			"trace": fmt.Sprint(err),
		})
	}

	if count == 0 {
		return c.SendStatus(fiber.StatusNotFound)
	}

	return c.SendStatus(fiber.StatusNoContent)
}

func UpdateTag(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid tag ID"})
	}

	tag := types.Tag{ID: id}
	if err := c.BodyParser(&tag); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse request"})
	}

	count, err := db.UpdateTag(tag)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Cannot update tag",
			"trace": fmt.Sprint(err),
		})
	}

	if count == 0 {
		return c.SendStatus(fiber.StatusNotFound)
	}

	return c.SendStatus(fiber.StatusOK)
}
