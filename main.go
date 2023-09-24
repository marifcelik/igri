package main

import (
	"go-chat/pkg/utils"
	"strconv"
	"time"

	"github.com/charmbracelet/log"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	fiber_session "github.com/gofiber/fiber/v2/middleware/session"
	fiber_utils "github.com/gofiber/fiber/v2/utils"
	"github.com/gofiber/storage/redis/v3"
)

var store *redis.Storage
var session *fiber_session.Store
var app *fiber.App

func init() {
	app = fiber.New()

	store = redis.New(redis.Config{
		URL: utils.GetRedisURL(),
	})

	session = fiber_session.New(fiber_session.Config{
		Storage:   store,
		KeyLookup: "header:Authorization",
		KeyGenerator: func() string {
			return "Bearer " + fiber_utils.UUIDv4()
		},
		Expiration: time.Minute,
	})
}

func main() {
	log.Info("main")

	app.Use(logger.New())

	app.Get("/", func(c *fiber.Ctx) error {
		userSess, err := session.Get(c)
		if err != nil {
			log.Error(err)
		}

		count := userSess.Get("count")
		switch temp := count.(type) {
		case int:
			count = temp + 1
			userSess.Set("count", count)

		case nil:
			count = 1
			userSess.Set("count", count)
		}

		if err := userSess.Save(); err != nil {
			log.Error(err)
		}

		return c.SendString(strconv.Itoa(count.(int)))
	})

	log.Fatal(app.Listen(utils.GetListenAddr()))
}
