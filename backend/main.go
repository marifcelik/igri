package main

import (
	"go-chat/config"
	"go-chat/db"
	"go-chat/internal/auth"
	"go-chat/internal/message"
	"go-chat/internal/ws"
	"go-chat/middlewares"
	"go-chat/storage"

	"github.com/charmbracelet/log"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	log.Info("main")
	log.Error("deneme")
	log.Warn("denem 2")
	app := fiber.New()
	app.Use(logger.New())

	// XXX may be i can create an interface for setup functions
	auth.Setup(app, db.DB)
	message.Setup(app)
	ws.Setup(app)

	app.Get("/", middlewares.AuthMiddleware, func(c *fiber.Ctx) error {
		userSess, err := storage.Session.Get(c)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).SendString("unauthorized")
		}

		count, ok := userSess.Get("count").(int)
		if !ok {
			userSess.Set("count", 0)
		}

		userSess.Set("count", count+1)
		if err := userSess.Save(); err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		return c.JSON(fiber.Map{
			"count": count,
		})
	})

	log.Fatal(app.Listen(config.GetListenAddr()))
}
