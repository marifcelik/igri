package message

import (
	"go-chat/middlewares"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	r := app.Group("message")
	r.Use(middlewares.AuthMiddleware)

	// TODO implement get message queries like sender=x, receiver=x
	r.Get("/", handleGetUserMessages)
	r.Get("/:id", handleGetMessage)
}
