package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func main() {
	driver := "sqlite3"
	connStr := "db.sqlite"
	port := "3001"

	if err := InitDB(driver, connStr); err != nil {
		log.Fatal(err)
	}

	app := fiber.New()

	app.Post("/list", createListHandler)
	app.Get("/lists/:board_id", getAllListsHandler)
	app.Get("/list/:id", getListHandler)
	app.Delete("/list/:id", deleteListHandler)
	app.Put("/list/:id", updateListHandler)

	log.Fatal(app.Listen(fmt.Sprintf(":%v", port)))
}

func createListHandler(c *fiber.Ctx) error {
	list := List{}
	if err := c.BodyParser(&list); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse request"})
	}

	id, err := Create(list)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Cannot create list"})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"id": id})
}

func getAllListsHandler(c *fiber.Ctx) error {
	boardID, err := strconv.Atoi(c.Params("board_id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid board ID"})
	}

	lists, err := GetAll(boardID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Cannot retrieve lists"})
	}

	return c.JSON(lists)
}

func getListHandler(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid list ID"})
	}

	list, err := Get(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Cannot retrieve list"})
	}
	if list == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "List not found"})
	}

	return c.JSON(list)
}

func deleteListHandler(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid list ID"})
	}

	err = Delete(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Cannot delete list"})
	}

	return c.SendStatus(fiber.StatusNoContent)
}

func updateListHandler(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid list ID"})
	}

	list := List{ID: id}
	if err := c.BodyParser(&list); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse request"})
	}

	err = Update(list)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Cannot update list"})
	}

	return c.SendStatus(fiber.StatusOK)
}
