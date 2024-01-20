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
	router.Post("/:id/options", CreateTagOption)
	router.Get("/:id/options", GetTagOptions)
	router.Delete("/:id/options/:option_id", DeleteTagOption)
	router.Put("/:id/options/:option_id", UpdateTagOption)
	router.Delete("/:id/options", DeleteTagOptions)
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

	err = c.Status(fiber.StatusCreated).JSON(fiber.Map{"id": id})
	if err != nil {
		return err
	}

	tag.ID = id

	source := c.Get("X-Request-Source")
	if source == "" {
		return nil
	}

	publish(fiber.Map{
		"object": "projectTag",
		"action": "create",
		"data": tag,
		"X-Request-Source": source,
	})

	return nil
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

	return c.Status(fiber.StatusOK).JSON(tag)
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

	err = c.SendStatus(fiber.StatusNoContent)
	if err != nil {
		return err
	}

	source := c.Get("X-Request-Source")
	if source == "" {
		return nil
	}

	publish(fiber.Map{
		"object": "projectTag",
		"action": "delete",
		"id": id,
		"X-Request-Source": source,
	})

	return nil
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

	err = c.SendStatus(fiber.StatusNoContent)
	if err != nil {
		return err
	}

	source := c.Get("X-Request-Source")
	if source == "" {
		return nil
	}

	publish(fiber.Map{
		"object": "projectTag",
		"action": "update",
		"id": id,
		"changes": tag,
		"X-Request-Source": source,
	})

	return nil
}

func CreateTagOption(c *fiber.Ctx) error {
	tagID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid tag ID"})
	}

	option := types.TagOption{TagID: tagID}
	if err := c.BodyParser(&option); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse request"})
	}

	exist, err := db.ExistTag(tagID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error finding tag",
			"trace": fmt.Sprint(err),
		})
	}
	if !exist {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Tag not found"})
	}

	id, err := db.CreateTagOption(option)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Cannot create tag option",
			"trace": fmt.Sprint(err),
		})
	}

	err = c.Status(fiber.StatusCreated).JSON(fiber.Map{"id": id})
	if err != nil {
		return err
	}

	option.ID = id

	source := c.Get("X-Request-Source")
	if source == "" {
		return nil
	}

	publish(fiber.Map{
		"object": "tagOption",
		"action": "create",
		"data": option,
		"X-Request-Source": source,
	})

	return nil
}

func GetTagOptions(c *fiber.Ctx) error {
	tagID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid tag ID"})
	}

	exist, err := db.ExistTag(tagID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error finding tag",
			"trace": fmt.Sprint(err),
		})
	}
	if !exist {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Tag not found"})
	}

	options, err := db.GetTagOptions(tagID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Cannot retrieve tag options",
			"trace": fmt.Sprint(err),
		})
	}

	return c.Status(fiber.StatusOK).JSON(options)
}

func DeleteTagOption(c *fiber.Ctx) error {
	tagID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid tag ID"})
	}

	optionID, err := strconv.Atoi(c.Params("option_id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid tag option ID"})
	}

	exist, err := db.ExistTag(tagID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error finding tag",
			"trace": fmt.Sprint(err),
		})
	}
	if !exist {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Tag not found"})
	}

	count, err := db.DeleteTagOption(optionID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Cannot delete tag option",
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
		"object": "tagOption",
		"action": "delete",
		"tag_id": tagID,
		"option_id": optionID,
		"X-Request-Source": source,
	})

	return nil
}

func UpdateTagOption(c *fiber.Ctx) error {
	tagID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid tag ID"})
	}

	optionID, err := strconv.Atoi(c.Params("option_id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid tag option ID"})
	}

	option := types.TagOption{ID: optionID, TagID: tagID}
	if err := c.BodyParser(&option); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse request"})
	}

	exist, err := db.ExistTag(tagID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error finding tag",
			"trace": fmt.Sprint(err),
		})
	}
	if !exist {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Tag not found"})
	}

	count, err := db.UpdateTagOption(option)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Cannot update tag option",
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
		"object": "tagOption",
		"action": "update",
		"tag_id": tagID,
		"option_id": optionID,
		"changes": option,
		"X-Request-Source": source,
	})

	return nil
}

func DeleteTagOptions(c *fiber.Ctx) error {
	tagID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid tag ID"})
	}

	exist, err := db.ExistTag(tagID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error finding tag",
			"trace": fmt.Sprint(err),
		})
	}
	if !exist {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Tag not found"})
	}

	err = db.DeleteTagOptions(tagID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Cannot delete tag options",
			"trace": fmt.Sprint(err),
		})
	}

	return c.SendStatus(fiber.StatusNoContent)

	// TODO: publish event
}
