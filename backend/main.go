package main

import (
	"net/http"
	"strconv"

	"go-chat/config"
	"go-chat/db"
	"go-chat/internal/auth"
	"go-chat/internal/message"
	"go-chat/internal/ws"
	"go-chat/middlewares"
	st "go-chat/storage"

	"github.com/charmbracelet/log"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	app := chi.NewRouter()
	app.Use(
		middleware.Logger,
		middleware.RequestID,
		middleware.RealIP,
		middleware.RedirectSlashes,
		middleware.StripSlashes,
	)

	// XXX may be i can create an interface for setup functions
	auth.Setup(app, db.DB)
	message.Setup(app)
	ws.Setup(app, db.DB)

	app.With(middlewares.AuthMiddleware).Get("/", func(w http.ResponseWriter, r *http.Request) {
		count := st.Session.GetInt(r.Context(), "count")
		count++
		log.Info(st.Session.Status(r.Context()))
		st.Session.Put(r.Context(), "count", count)

		w.WriteHeader(http.StatusAccepted)
		w.Write([]byte(strconv.Itoa(count) + "\n"))
	})

	log.Info("Server listening on", "addr", config.GetListenAddr())
	log.Fatal(http.ListenAndServe(config.GetListenAddr(), st.Session.LoadAndServeHeader(app)))
}
