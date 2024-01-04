package utils

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

func Cache(c *fiber.Ctx, lastEdit time.Time) (bool, error) {
	clientLast, err := time.Parse(time.RFC1123, c.Get("If-Modified-Since"))
	if err != nil {
		return false, err
	}
	
	if clientLast.Before(lastEdit) {
		c.SendStatus(fiber.StatusNotModified)
		return true, nil
	}

	return false, nil
}