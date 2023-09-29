package routes

import (
	authHandler "go-chat/internal/handlers/http/auth"

	"github.com/gofiber/fiber/v2"
)

func SetupAuthRoutes(app *fiber.App) {
	router := app.Group("auth")

	router.Post("/login", authHandler.Login)
	router.Post("/logout", authHandler.Logout)
	router.Post("/register", authHandler.Register)
}
