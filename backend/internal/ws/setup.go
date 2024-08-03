package ws

import (
	"context"
	"net/http"

	"go-chat/internal/auth"
	"go-chat/middlewares"
	"go-chat/utils"

	"github.com/go-chi/chi/v5"
	"github.com/lxzan/gws"
	"go.mongodb.org/mongo-driver/mongo"
)

type WSSession struct {
	ctx context.Context
}

func Setup(c *chi.Mux, db *mongo.Database) {
	repo := NewWSRepo(db)
	authRepo := auth.NewAuthRepo(db)
	handler := NewWSHandler(repo, authRepo)

	c.Route("/_ws", func(r chi.Router) {
		r.Use(
			middlewares.UpgradeChecher,
			middlewares.WsHeader,
			middlewares.Auth,
		)

		r.Get("/", func(w http.ResponseWriter, req *http.Request) {
			upgrader := gws.NewUpgrader(handler, &gws.ServerOption{
				ParallelEnabled:   true,
				Recovery:          gws.Recovery,
				PermessageDeflate: gws.PermessageDeflate{Enabled: true},
			})

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
