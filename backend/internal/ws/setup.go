package ws

import (
	"go-chat/middlewares"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	ws := app.Group("_ws")
	ws.Use(middlewares.AuthMiddleware, middlewares.RequestUpgrade)

	ws.Get("/", websocket.New(handleMessages))
}
