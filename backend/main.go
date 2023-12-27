package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	_ "github.com/mattn/go-sqlite3"

	"git.bhasher.com/bhasher/focus/backend/db"
	"git.bhasher.com/bhasher/focus/backend/handlers"
)

func main() {
	driver := "sqlite3"
	connStr := "db.sqlite"
	port := "3000"

	if err := db.InitDB(driver, connStr); err != nil {
		log.Fatal(err)
	}

	app := fiber.New()

	app.Get("/projects", handlers.GetAllProjects)
	app.Get("/project/:id", handlers.GetProject)
	app.Post("/project", handlers.CreateProject)
	app.Put("/project/:id", handlers.UpdateProject)
	app.Delete("/project/:id", handlers.DeleteProject)

	app.Post("/list", handlers.CreateList)
	app.Get("/lists/:board_id", handlers.GetAllListsOf)
	app.Get("/list/:id", handlers.GetList)
	app.Delete("/list/:id", handlers.DeleteList)
	app.Put("/list/:id", handlers.UpdateList)

	log.Fatal(app.Listen(fmt.Sprintf(":%v", port)))
}
