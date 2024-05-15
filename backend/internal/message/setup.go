package message

import (
	"go-chat/middlewares"

	"github.com/go-chi/chi/v5"
)

func Setup(c *chi.Mux) {
	c.Route("/message", func(r chi.Router) {
		r.Use(middlewares.Auth)

		// TODO implement get message queries like sender=x, receiver=x
		r.Get("/", handleGetUserMessages)
		r.Get("/:id", handleGetMessage)
	})
}
