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

func (c *wsHandler) OnOpen(conn *gws.Conn) {
	_ = conn.SetDeadline(time.Now().Add(PingInterval + PingWait))
}

func (c *wsHandler) OnClose(conn *gws.Conn, err error) {}

func (c *wsHandler) OnPing(conn *gws.Conn, payload []byte) {
	_ = conn.SetDeadline(time.Now().Add(PingInterval + PingWait))
	_ = conn.WritePong(nil)
}

func (c *wsHandler) OnPong(conn *gws.Conn, payload []byte) {}

func (c *wsHandler) OnMessage(conn *gws.Conn, message *gws.Message) {
	defer message.Close()
	conn.WriteMessage(message.Opcode, message.Bytes())
}
