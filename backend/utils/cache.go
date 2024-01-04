package utils

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

func Cache(c *fiber.Ctx, lastEdit *time.Time) (bool, error) {
	ifModifiedSince := c.Get("If-Modified-Since")
	if ifModifiedSince == "" {
		return false, nil
	}

	clientLast, err := time.Parse(time.RFC1123, ifModifiedSince)
	if err != nil {
		return false, err
	}

	if !lastEdit.After(clientLast) {
		c.SendStatus(fiber.StatusNotModified)
		return true, nil
	}

	c.Set("Last-Modified", lastEdit.Format(time.RFC1123))
	return false, nil
}