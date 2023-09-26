package main

import (
	_ "go-chat/internal/db"
	"go-chat/internal/routes"
	"go-chat/internal/storage"
	"go-chat/pkg/middlewares"
	"go-chat/pkg/utils"

	"strconv"

	"github.com/charmbracelet/log"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	log.Info("main")
	app := fiber.New()

	app.Use(logger.New())

	app.Get("/", middlewares.AuthMiddleware, countHandler)

	routes.SetupAuthRoutes(app)

	log.Fatal(app.Listen(utils.GetListenAddr()))
}

func countHandler(c *fiber.Ctx) error {
	userSess, err := storage.Session.Get(c)
	if err != nil {
		log.Error(err)
	}
	log.Info(userSess.Fresh())
	log.Info(userSess.Get("user"))

	count := userSess.Get("count")
	switch temp := count.(type) {
	case int:
		count = temp + 1
		userSess.Set("count", count)

	case nil:
		count = 1
		userSess.Set("count", count)
	}

	// it resets the user session
	if err := userSess.Save(); err != nil {
		log.Error(err)
	}

	return c.SendString(strconv.Itoa(count.(int)))
}
