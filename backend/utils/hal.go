package utils

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func SendHAL(c *fiber.Ctx, status int, data fiber.Map) error {
	if err := c.JSON(data); err != nil {
		return err
	}

	c.Status(status)
	c.Context().SetContentType("application/hal+json")
	return nil
}

func HALProjectLinks(id int) fiber.Map {
	return fiber.Map{
		"_links": fiber.Map{
			"self":  fiber.Map{"href": fmt.Sprintf("/api/projects/%v", id)},
			"cards": fiber.Map{"href": fmt.Sprintf("/api/projects/%v/cards", id)},
			"tags":  fiber.Map{"href": fmt.Sprintf("/api/projects/%v/tags", id)},
		},
	}
}
