package routes

import (
	handlers "go-chat/internal/handlers/http"

	"github.com/gofiber/fiber/v2"
)

func SetupAuthRoutes(app *fiber.App) {
	router := app.Group("auth")

	router.Post("/login", handlers.LoginHandler)
	router.Post("/logout", handlers.LogoutHandler)
	router.Post("/register", handlers.RegisterHandler)
}
