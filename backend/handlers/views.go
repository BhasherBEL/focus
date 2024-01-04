package handlers

import (
	"fmt"
	"strconv"

	"git.bhasher.com/bhasher/focus/backend/db"
	"git.bhasher.com/bhasher/focus/backend/types"
	"github.com/gofiber/fiber/v2"
)

func viewsRouter(router fiber.Router) error {
	router.Post("/", CreateView)
	router.Get("/:id", GetView)
	router.Put("/:id", UpdateView)
	router.Delete("/:id", DeleteView)

	return nil
}

func CreateView(c *fiber.Ctx) error {
	view := types.View{}
	if err := c.BodyParser(&view); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse request",
			"trace": fmt.Sprint(err),
		})
	}

	id, err := db.CreateView(view)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Cannot create view",
			"trace": fmt.Sprint(err),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"id": id,
	})
}

func GetView(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid view ID"})
	}

	view, err := db.GetView(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Cannot retrieve view",
			"trace": fmt.Sprint(err),
		})
	}
	if view == nil {
		return c.SendStatus(fiber.StatusNotFound)
	}

	return c.Status(fiber.StatusOK).JSON(view)
}

func UpdateView(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid view ID"})
	}

	view := types.View{ID: id}
	if err := c.BodyParser(&view); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse request",
			"trace": fmt.Sprint(err),
		})
	}

	count, err := db.UpdateView(view)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Cannot update view",
			"trace": fmt.Sprint(err),
		})
	}
	if count == 0 {
		return c.SendStatus(fiber.StatusNotFound)
	}

	return c.SendStatus(fiber.StatusNoContent)
}

func DeleteView(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid view ID"})
	}

	count, err := db.DeleteView(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Cannot delete view",
			"trace": fmt.Sprint(err),
		})
	}

	if count == 0 {
		return c.SendStatus(fiber.StatusNotFound)
	}

	return c.SendStatus(fiber.StatusNoContent)
}
