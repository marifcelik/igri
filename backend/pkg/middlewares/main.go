package middlewares

import (
	"go-chat/internal/storage"
	"go-chat/pkg/utils"

	"github.com/charmbracelet/log"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware(c *fiber.Ctx) error {
	log := log.WithPrefix("AUTH MW")

	userSess, err := storage.Session.Get(c)
	if err != nil {
		log.Error("session get error", "err", err)
		// TODO send a more meaningful response
		return c.Status(fiber.ErrUnavailableForLegalReasons.Code).JSON(fiber.Map{
			"err": err.Error(),
		})
	}

	name := userSess.Get("user")
	if name == nil || name.(string) == "" {
		log.Warn("unauthorized request", "from", utils.GetIPAddr(c))
		return c.SendStatus(fiber.ErrUnauthorized.Code)
	}

	return c.Next()
}

func RequestUpgrade(c *fiber.Ctx) error {
	if websocket.IsWebSocketUpgrade(c) {
		c.Locals("allowed", true)
		return c.Next()
	}
	return fiber.ErrUpgradeRequired
}
