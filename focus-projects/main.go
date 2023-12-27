package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	driver := "sqlite3"
	connStr := "db.sqlite"
	port := "3000"

	if err := InitDB(driver, connStr); err != nil {
		log.Fatal(err)
	}

	app := fiber.New()

	app.Get("/projects", getAllProjectsHandler)
	app.Get("/project/:id", getProjectHandler)
	app.Post("/project", createProjectHandler)
	app.Put("/project/:id", updateProjectHandler)
	app.Delete("/project/:id", deleteProjectHandler)

	log.Fatal(app.Listen(fmt.Sprintf(":%v", port)))
}

func getAllProjectsHandler(c *fiber.Ctx) error {
	projects, err := GetAll()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Cannot retrieve projects"})
	}

	fmt.Println(projects)

	return c.JSON(projects)
}

func getProjectHandler(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid project ID"})
	}

	project, err := Get(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error fetching project"})
	}
	if project == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Project not found"})
	}
	return c.JSON(project)
}

func createProjectHandler(c *fiber.Ctx) error {
	p := new(Project)
	if err := c.BodyParser(p); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Error parsing request"})
	}

	id, err := Create(*p)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error creating project"})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"id": id})
}

func updateProjectHandler(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid project ID"})
	}

	p := Project{ID: id}
	if err := c.BodyParser(&p); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Error parsing request"})
	}

	err = Update(p)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error updating project"})
	}

	return c.SendStatus(fiber.StatusOK)
}

func deleteProjectHandler(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid project ID"})
	}

	err = Delete(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error deleting project"})
	}

	return c.SendStatus(fiber.StatusOK)
}
