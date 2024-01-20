package handlers

import (
	"fmt"
	"strconv"

	"git.bhasher.com/bhasher/focus/backend/db"
	"git.bhasher.com/bhasher/focus/backend/types"

	"github.com/gofiber/fiber/v2"
)

func filtersRouter(router fiber.Router) error {
	router.Post("/", CreateFilter)
	router.Get("/:id", GetFilter)
	router.Put("/:id", UpdateFilter)
	router.Delete("/:id", DeleteFilter)

	return nil
}

func CreateFilter(c *fiber.Ctx) error {
	filter := types.Filter{}
	if err := c.BodyParser(&filter); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse request",
			"trace": fmt.Sprint(err),
		})
	}

	exist, err := db.ExistView(filter.ViewID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error finding view",
			"trace": fmt.Sprint(err),
		})
	}
	if !exist {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "View not found"})
	}

	id, err := db.CreateFilter(filter)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Cannot create filter",
			"trace": fmt.Sprint(err),
		})
	}

	err = c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"id": id,
	})
	if err != nil {
		return err
	}

	filter.ID = id

	source := c.Get("X-Request-Source")
	if source == "" {
		return nil
	}

	publish(fiber.Map{
		"object": "filter",
		"action": "create",
		"data": filter,
		"X-Request-Source": source,
	})

	return nil
}

func GetFilter(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid filter ID"})
	}

	filter, err := db.GetFilter(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Cannot retrieve filter",
			"trace": fmt.Sprint(err),
		})
	}
	if filter == nil {
		return c.SendStatus(fiber.StatusNotFound)
	}

	return c.Status(fiber.StatusOK).JSON(filter)
}

func UpdateFilter(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid filter ID"})
	}

	filter := types.Filter{ID: id}
	if err := c.BodyParser(&filter); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse request",
			"trace": fmt.Sprint(err),
		})
	}
	exist, err := db.ExistView(filter.ViewID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error finding view",
			"trace": fmt.Sprint(err),
		})
	}
	if !exist {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "View not found"})
	}

	count, err := db.UpdateFilter(filter)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Cannot update filter",
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
		"object": "filter",
		"action": "update",
		"id": id,
		"changes": filter,
		"X-Request-Source": source,
	})

	return nil
}

func DeleteFilter(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid filter ID"})
	}

	count, err := db.DeleteFilter(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Cannot delete filter",
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
		"object": "filter",
		"action": "delete",
		"id": id,
		"X-Request-Source": source,
	})

	return nil
}