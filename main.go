package main

import (
	"go-chat/internal/routes"
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

	log.Fatal(app.Listen(utils.GetListenAddr()))
}
