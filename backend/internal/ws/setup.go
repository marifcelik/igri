package ws

import (
	"net/http"

	"go-chat/internal/auth"
	"go-chat/middlewares"
	"go-chat/utils"

	"github.com/go-chi/chi/v5"
	"github.com/marifcelik/gws"
	"go.mongodb.org/mongo-driver/mongo"
)

func Setup(c *chi.Mux, db *mongo.Database) {
	repo := NewWSRepo(db)
	authRepo := auth.NewAuthRepo(db)
	handler := NewWSHandler(repo, authRepo)

	upgrader := gws.NewUpgrader(handler, &gws.ServerOption{
		ParallelEnabled:   true,
		Recovery:          gws.Recovery,
		PermessageDeflate: gws.PermessageDeflate{Enabled: true},
	})

	c.With(
		middlewares.UpgradeChecher,
		middlewares.WsHeader,
		middlewares.Auth,
	).Route("/_ws", func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, req *http.Request) {
			conn, err := upgrader.Upgrade(w, req)
			if err != nil {
				log.WithPrefix("WS").Error("websocket upgrade error", "err", err, "ip", req.RemoteAddr)
				utils.ErrResp(w, http.StatusInternalServerError)
				return
			}
			go conn.ReadLoop()
		})
	})
}
