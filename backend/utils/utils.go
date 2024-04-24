package utils

import (
	"strings"

	"github.com/charmbracelet/log"
	"github.com/gofiber/fiber/v2"
)

func GetIPAddr(c *fiber.Ctx) string {
	switch {
	case c.IsFromLocal():
		return c.Context().LocalIP().String()
	case len(c.IPs()) != 0:
		return strings.Join(c.IPs(), "")
	case c.IP() != "":
		return c.IP()
	default:
		return c.Context().RemoteAddr().String()
	}
}

func InternalErr(c *fiber.Ctx, err error) error {
	return c.Status(fiber.ErrInternalServerError.Code).JSON(fiber.Map{
		"err": err.Error(),
	})
}

// Check the error and exit if its not nil.
// The parameters after the second parameter will be joined into a single string
func CheckErr(err error, msgParams ...string) {
	msg := strings.Join(msgParams, ", ")

	if err != nil {
		if msg != "" {
			log.Fatal(msg, "err", err)
		} else {
			log.Fatal(err)
		}
	}
}
