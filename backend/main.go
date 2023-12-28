package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	_ "github.com/mattn/go-sqlite3"

	"git.bhasher.com/bhasher/focus/backend/db"
	"git.bhasher.com/bhasher/focus/backend/handlers"
)

func main() {
	driver := "sqlite3"
	connStr := "db.sqlite"
	port := "3000"
	origins := "*"

	if err := db.InitDB(driver, connStr); err != nil {
		log.Fatal(err)
	}

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: origins,
		AllowMethods: "GET,POST,PUT,DELETE",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	app.Get("/api/projects", handlers.GetAllProjects)
	app.Get("/api/project/:id", handlers.GetProject)
	app.Post("/api/project", handlers.CreateProject)
	app.Put("/api/project/:id", handlers.UpdateProject)
	app.Delete("/api/project/:id", handlers.DeleteProject)

	app.Post("/api/list", handlers.CreateList)
	app.Get("/api/lists/:board_id", handlers.GetAllListsOf)
	app.Get("/api/list/:id", handlers.GetList)
	app.Delete("/api/list/:id", handlers.DeleteList)
	app.Put("/api/list/:id", handlers.UpdateList)

	log.Fatal(app.Listen(fmt.Sprintf(":%v", port)))
}
