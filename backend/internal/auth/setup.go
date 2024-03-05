package auth

import (
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

func Setup(app *fiber.App, db *mongo.Database) {
	r := app.Group("auth")

	repo := NewAuthRepo(db)
	handler := NewAuthHandler(repo)

	r.Post("/login", handler.Login)
	r.Post("/logout", handler.Logout)
	r.Post("/register", handler.Register)
}
