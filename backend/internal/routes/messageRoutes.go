package routes

import (
	messageHandler "go-chat/internal/handlers/http/message"
	websocketHandler "go-chat/internal/handlers/ws"
	"go-chat/pkg/middlewares"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
)

func SetupMessageRoutes(app *fiber.App) {
	router := app.Group("message")
	router.Use(middlewares.AuthMiddleware)

	// TODO implement get message queries like sender=x, receiver=x
	router.Get("/", messageHandler.GetUserMessages)
	router.Get("/:id", messageHandler.GetMessage)

	ws := app.Group("ws")
	ws.Use(middlewares.AuthMiddleware, middlewares.RequestUpgrade)

	ws.Get("/", websocket.New(websocketHandler.SendMessage))
}
