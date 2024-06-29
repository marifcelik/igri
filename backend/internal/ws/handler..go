package ws

import (
	"time"

	clog "github.com/charmbracelet/log"
	"github.com/lxzan/gws"

	"go-chat/internal/auth"
)

const (
	PingInterval = 5 * time.Second
	PingWait     = 10 * time.Second
)

var log = clog.WithPrefix("WS")

type wsHandler struct {
	repo     *wsRepo
	userRepo *auth.AuthRepo
}

func NewWSHandler(repo *wsRepo, userRepo *auth.AuthRepo) *wsHandler {
	return &wsHandler{repo: repo, userRepo: userRepo}
}

// FIX there is an close error after 15 seconds, idk why. i tried other ws packages, but still the same
// on client side, it says "close 1006"
func (c *wsHandler) OnOpen(conn *gws.Conn) {
	err := conn.SetDeadline(time.Now().Add(PingInterval + PingWait))
	if err != nil {
		log.Warn("setdeadline", "err", err)
	}
}

func (c *wsHandler) OnClose(conn *gws.Conn, err error) {
	log.Warn("connection closed", "err", err)
}

func (c *wsHandler) OnPing(conn *gws.Conn, payload []byte) {
	err := conn.SetDeadline(time.Now().Add(PingInterval + PingWait))
	if err != nil {
		log.Error("setdeadline", "err", err)
	}
	err = conn.WritePong(nil)
	if err != nil {
		log.Error("write pong", "err", err)
	}
}

func (c *wsHandler) OnPong(conn *gws.Conn, payload []byte) {}

func (c *wsHandler) OnMessage(conn *gws.Conn, message *gws.Message) {
	defer message.Close()

	log.Info("message received", "message", message.Data.String())

	conn.WriteMessage(message.Opcode, []byte("hii"))
}
