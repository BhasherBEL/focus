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

	app.Post("/api/projects", handlers.CreateProject)
	app.Get("/api/projects", handlers.GetAllProjects)
	app.Get("/api/projects/:id", handlers.GetProject)
	app.Put("/api/projects/:id", handlers.UpdateProject)
	app.Delete("/api/projects/:id", handlers.DeleteProject)
	app.Get("/api/projects/:project_id/cards", handlers.GetProjectCards)
	app.Get("/api/projects/:project_id/tags", handlers.GetProjectTags)

	app.Post("/api/cards", handlers.CreateCard)
	app.Get("/api/cards/:id", handlers.GetCard)
	app.Put("/api/cards/:id", handlers.UpdateCard)
	app.Delete("/api/cards/:id", handlers.DeleteCard)

	app.Post("/api/tags", handlers.CreateTag)
	app.Get("/api/tags/:id", handlers.GetTag)
	app.Delete("/api/tags/:id", handlers.DeleteTag)
	app.Put("/api/tags/:id", handlers.UpdateTag)

	app.Post("/api/cards/:card_id/tags/:tag_id", handlers.CreateCardTag)
	app.Get("/api/cards/:card_id/tags", handlers.GetCardTags)
	app.Put("/api/cards/:card_id/tags/:tag_id", handlers.UpdateCardTag)
	app.Delete("/api/cards/:card_id/tags/:tag_id", handlers.DeleteCardTag)
	app.Delete("/api/cards/:card_id/tags", handlers.DeleteCardTags)

	log.Fatal(app.Listen(fmt.Sprintf(":%v", port)))
}
