package handlers

import (
	"encoding/json"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

var (
	subscribers = make(map[*websocket.Conn]bool)
)

func wsRouter(router fiber.Router) error {
	router.Use("/", upgradeWebsocket)
	router.Get("/", websocket.New(handleWebsocket))
	return nil;
}

func upgradeWebsocket(c *fiber.Ctx) error {
	if websocket.IsWebSocketUpgrade(c) {
		c.Locals("allowed", true)
		return c.Next()
	}
	return fiber.ErrUpgradeRequired
}

func handleWebsocket(c *websocket.Conn) {
	subscribers[c] = true
	defer func() {
		delete(subscribers, c)
		c.Close()
	}()

	for {
		_, _, err := c.ReadMessage()
		if err != nil {
			return
		}
	}
}

func publish(content fiber.Map) {
	jsonMessage, err := json.Marshal(content)
	if err != nil {
		log.Println("Error marshalling JSON:", err)
		return
	}

	for s := range subscribers {
		if err := s.WriteMessage(websocket.TextMessage, jsonMessage); err != nil {
			log.Println("Error writing to websocket:", err)
			s.Close()
		}
	}
}