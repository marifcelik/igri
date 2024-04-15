package ws

import (
	"context"

	clog "github.com/charmbracelet/log"
	"github.com/gofiber/contrib/websocket"

	"go-chat/internal/auth"
)

var log = clog.WithPrefix("WS")

type wsHandler struct {
	repo     *wsRepo
	userRepo *auth.AuthRepo
}

func NewWSHandler(repo *wsRepo, userRepo *auth.AuthRepo) *wsHandler {
	return &wsHandler{repo: repo, userRepo: userRepo}
}

func (w *wsHandler) HandleMessages(c *websocket.Conn) {
	log.SetPrefix(log.GetPrefix() + " - " + c.LocalAddr().String())
	log.Print("new connection")

	oldCloseHandler := c.CloseHandler()
	c.SetCloseHandler(func(code int, text string) error {
		log.Info("connection closed", "code", code, "text", text)
		return oldCloseHandler(code, text)
	})

	c.SetPingHandler(func(appData string) error {
		log.Info("ping received", "appData", appData)
		c.WriteMessage(websocket.PongMessage, []byte("PONG"))
		return nil
	})

	// wait for incoming messages
	w.handleIncomingMessage(c)
}

func (w *wsHandler) handleIncomingMessage(c *websocket.Conn) {
	var (
		msg WsMessageDTO = WsMessageDTO{}
		err error
	)

	ctx := context.Background()

	for {
		err = c.ReadJSON(&msg)
		if err != nil {
			log.Error("error reading message", "error", err)
			break
		}
		isExist, err := w.userRepo.CheckUsername(msg.Sender, ctx)
		if err != nil {
			log.Warn("error getting user", "error", err)
			break
		}
		if !isExist {
			log.Warn("user not found", "username", msg.Sender)
			break
		}

		log.Info("message received", "message", msg)
	}
}
