package handlers

import (
	"fmt"
	"strconv"
	"time"

	"git.bhasher.com/bhasher/focus/backend/db"
	"git.bhasher.com/bhasher/focus/backend/types"
	"git.bhasher.com/bhasher/focus/backend/utils"
	"github.com/gofiber/fiber/v2"
)

func projectsRouter(router fiber.Router) error {
	router.Post("/", CreateProject)
	router.Get("/", GetAllProjects)
	router.Get("/:id", GetProject)
	router.Put("/:id", UpdateProject)
	router.Delete("/:id", DeleteProject)
	router.Get(":id/cards", GetProjectCards)
	router.Get(":id/tags", GetProjectTags)
	router.Get(":id/views", GetProjectViews)
	return nil
}

var projectsLastEdit time.Time = time.Now().Truncate(time.Second);

func GetAllProjects(c *fiber.Ctx) error {
	isCached, err := utils.Cache(c, &projectsLastEdit);
	if err == nil && isCached {
		return nil;
	}

	projects, err := db.GetAllProjects()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Cannot retrieve projects",
			"trace": fmt.Sprint(err),
		})
	}
	
	err = c.Status(fiber.StatusOK).JSON(projects)
	if err != nil {
		return err;
	}

	return nil;
}

func GetProject(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid project ID"})
	}

	project, err := db.GetProject(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error fetching project",
			"trace": fmt.Sprint(err),
		})
	}
	if project == nil {
		return c.SendStatus(fiber.StatusNotFound)
	}

	return c.Status(fiber.StatusOK).JSON(project)
}

func CreateProject(c *fiber.Ctx) error {
	p := new(types.Project)
	if err := c.BodyParser(p); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Error parsing request"})
	}

	id, err := db.CreateProject(*p)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error creating project",
			"trace": fmt.Sprint(err),
		})
	}

	projectsLastEdit = time.Now().Truncate(time.Second);

	if err = c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"id": id,
	}); err != nil {
		return err;
	}

	p.ID = id;

	source := c.Get("X-Request-Source");
	if source == "" {
		return nil;
	}

	publish(fiber.Map{
		"object": "project",
		"action": "create",
		"data": p,
		"X-Request-Source": source,
	})

	return nil;
}

func UpdateProject(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid project ID"})
	}

	p := types.Project{ID: id}
	if err := c.BodyParser(&p); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Error parsing request"})
	}

	count, err := db.UpdateProject(p)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error updating project",
			"trace": fmt.Sprint(err),
		})
	}

	if count == 0 {
		return c.SendStatus(fiber.StatusNotFound)
	}

	projectsLastEdit = time.Now().Truncate(time.Second);

	if err = c.SendStatus(fiber.StatusNoContent); err != nil {
		return err;
	}

	source := c.Get("X-Request-Source");
	if source == "" {
		return nil;
	}

	publish(fiber.Map{
		"object": "project",
		"action": "update",
		"id": id,
		"changes": p,
		"X-Request-Source": source,
})

	return nil;
}

func DeleteProject(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid project ID"})
	}

	count, err := db.DeleteProject(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error deleting project",
			"trace": fmt.Sprint(err),
		})
	}

	if count == 0 {
		return c.SendStatus(fiber.StatusNotFound)
	}

	projectsLastEdit = time.Now().Truncate(time.Second);

	if err = c.SendStatus(fiber.StatusNoContent); err != nil {
		return err;
	}

	source := c.Get("X-Request-Source");
	if source == "" {
		return nil;
	}

	publish(fiber.Map{
		"object": "project",
		"action": "delete",
		"id": id,
		"X-Request-Source": source,
	})

	return nil;
}

func GetProjectCards(c *fiber.Ctx) error {
	projectID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid project ID"})
	}

	exists, err := db.ExistProject(projectID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error finding project",
			"trace": fmt.Sprint(err),
		})
	}

	if !exists {
		return c.SendStatus(fiber.StatusNotFound)
	}

	cards, err := db.GetProjectsCards(projectID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Cannot retrieve cards",
			"trace": fmt.Sprint(err),
		})
	}

	return c.Status(fiber.StatusOK).JSON(cards)
}

func GetProjectTags(c *fiber.Ctx) error {
	projectID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid project ID"})
	}

	exists, err := db.ExistProject(projectID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error finding project",
			"trace": fmt.Sprint(err),
		})
	}

	if !exists {
		return c.SendStatus(fiber.StatusNotFound)
	}

	tags, err := db.GetProjectTags(projectID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Cannot retrieve tags",
			"trace": fmt.Sprint(err),
		})
	}

	return c.Status(fiber.StatusOK).JSON(tags)
}

func GetProjectViews(c *fiber.Ctx) error {
	projectID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid project ID"})
	}

	exists, err := db.ExistProject(projectID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error finding project",
			"trace": fmt.Sprint(err),
		})
	}

	if !exists {
		return c.SendStatus(fiber.StatusNotFound)
	}

	views, err := db.GetProjectViews(projectID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Cannot retrieve views",
			"trace": fmt.Sprint(err),
		})
	}

	return c.Status(fiber.StatusOK).JSON(views)
}
