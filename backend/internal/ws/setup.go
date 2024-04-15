package ws

import (
	"go-chat/internal/auth"
	"go-chat/middlewares"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

func Setup(app *fiber.App, db *mongo.Database) {
	ws := app.Group("_ws")
	ws.Use(
		middlewares.RequestUpgrade,
		middlewares.WsHeaderMiddleware,
		middlewares.AuthMiddleware,
	)

	repo := NewWSRepo(db)
	authRepo := auth.NewAuthRepo(db)
	handler := NewWSHandler(repo, authRepo)

	ws.Get("/", websocket.New(handler.HandleMessages))
}
