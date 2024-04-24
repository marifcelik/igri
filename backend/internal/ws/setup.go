package ws

import (
	"net/http"

	"go-chat/internal/auth"
	"go-chat/middlewares"
	"go-chat/utils"

	"github.com/go-chi/chi/v5"
	"github.com/lxzan/gws"
	"go.mongodb.org/mongo-driver/mongo"
)

func Setup(c *chi.Mux, db *mongo.Database) {
	repo := NewWSRepo(db)
	authRepo := auth.NewAuthRepo(db)
	handler := NewWSHandler(repo, authRepo)

	c.Route("/_ws", func(r chi.Router) {
		r.Use(
			middlewares.CheckIsUpgrade,
			middlewares.WsHeaderMiddleware,
			middlewares.AuthMiddleware,
		)

		upgrader := gws.NewUpgrader(handler, &gws.ServerOption{
			ParallelEnabled: true,
			Recovery:        gws.Recovery,
		})

		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			conn, err := upgrader.Upgrade(w, r)
			if err != nil {
				log.WithPrefix("WS").Error("websocket upgrade error", "err", err, "ip", r.RemoteAddr)
				utils.ErrResp(w, http.StatusInternalServerError)
				return
			}
			go func() {
				conn.ReadLoop()
			}()
		})
	})
}
