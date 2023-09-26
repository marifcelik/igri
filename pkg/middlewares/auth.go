package middlewares

import (
	"go-chat/internal/storage"

	"github.com/charmbracelet/log"
	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware(c *fiber.Ctx) error {
	log := log.WithPrefix("AUTH MW")

	userSess, err := storage.Session.Get(c)
	if err != nil {
		log.Error("session get error", "err", err)
	}

	name := userSess.Get("user")
	if name == nil || name.(string) == "" {
		log.Warn("unauthorized request", "from", c.IPs())
		return c.SendStatus(fiber.ErrUnauthorized.Code)
	}

	return c.Next()
}
