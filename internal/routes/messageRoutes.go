package routes

import (
	handlers "go-chat/internal/handlers/http"
	"go-chat/pkg/middlewares"

	"github.com/gofiber/fiber/v2"
)

func SetupMessageRoutes(app *fiber.App) {
	router := app.Group("message")
	router.Use(middlewares.AuthMiddleware)

	router.Get("/")
	router.Get("/:id", handlers.LogoutHandler)
	router.Post("/register", handlers.RegisterHandler)
}
