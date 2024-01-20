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
	router.Get("/:id/filters", GetView)
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

	err = c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"id": id,
	})
	if err != nil {
		return err
	}

	view.ID = id

	source := c.Get("X-Request-Source")
	if source == "" {
		return nil
	}

	publish(fiber.Map{
		"object": "view",
		"action": "create",
		"data": view,
		"X-Request-Source": source,
	})

	return nil
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

	err = c.SendStatus(fiber.StatusNoContent)
	if err != nil {
		return err
	}

	source := c.Get("X-Request-Source")
	if source == "" {
		return nil
	}

	publish(fiber.Map{
		"object": "view",
		"action": "update",
		"id": id,
		"changes": view,
		"X-Request-Source": source,
	})

	return nil
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

	err = c.SendStatus(fiber.StatusNoContent)
	if err != nil {
		return err
	}

	source := c.Get("X-Request-Source")
	if source == "" {
		return nil
	}

	publish(fiber.Map{
		"object": "view",
		"action": "delete",
		"id": id,
		"X-Request-Source": source,
	})

	return nil
}

func GetViewFilters(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid view ID"})
	}

	filters, err := db.GetViewFilters(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Cannot retrieve view filters",
			"trace": fmt.Sprint(err),
		})
	}

	return c.Status(fiber.StatusOK).JSON(filters)
}