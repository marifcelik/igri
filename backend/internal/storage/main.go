package storage

import (
	"go-chat/pkg/utils"

	fiber_session "github.com/gofiber/fiber/v2/middleware/session"
	fiber_utils "github.com/gofiber/fiber/v2/utils"
	"github.com/gofiber/storage/redis/v3"
)

var Session *fiber_session.Store

func init() {
	Session = fiber_session.New(fiber_session.Config{
		Storage: redis.New(redis.Config{
			URL: utils.GetRedisURL(),
		}),
		KeyLookup: "header:Authorization",
		KeyGenerator: func() string {
			return "Bearer " + fiber_utils.UUIDv4()
		},
		Expiration: utils.GetExpirationTime(),
	})
}
