package handlers

import "github.com/gofiber/fiber/v2"

func APIRouter(router fiber.Router) error {
	return v1Router(router.Group("/v1"))
}

func v1Router(router fiber.Router) error {
	projectsRouter(router.Group("/projects"))
	cardsRouter(router.Group("/cards"))
	tagsRouter(router.Group("/tags"))
	viewsRouter(router.Group("/views"))
	filtersRouter(router.Group("/filters"))
	return nil
}

// app.Post("/api/v1/cards", handlers.CreateCard)
// app.Get("/api/v1/cards/:id", handlers.GetCard)
// app.Put("/api/v1/cards/:id", handlers.UpdateCard)
// app.Delete("/api/v1/cards/:id", handlers.DeleteCard)

// app.Post("/api/v1/tags", handlers.CreateTag)
// app.Get("/api/v1/tags/:id", handlers.GetTag)
// app.Delete("/api/v1/tags/:id", handlers.DeleteTag)
// app.Put("/api/v1/tags/:id", handlers.UpdateTag)
