package main

import (
	"go-chat/internal/routes"
	"go-chat/internal/storage"
	"go-chat/pkg/middlewares"
	"go-chat/pkg/utils"

	"github.com/charmbracelet/log"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	log.Info("main")
	app := fiber.New()
	app.Use(logger.New())

	routes.SetupAuthRoutes(app)
	routes.SetupMessageRoutes(app)

	app.Get("/", middlewares.AuthMiddleware, func(c *fiber.Ctx) error {
		userSess, err := storage.Session.Get(c)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).SendString("unauthorized")
		}

		count := userSess.Get("count")
		if count == nil {
			userSess.Set("count", 0)
			count = 0
		}

		userSess.Set("count", count.(int)+1)
		if err := userSess.Save(); err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		return c.JSON(fiber.Map{
			"count": count,
		})
	})

	log.Fatal(app.Listen(utils.GetListenAddr()))
}
