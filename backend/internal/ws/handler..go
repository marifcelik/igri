package ws

import (
	"context"
	"encoding/json"
	"time"

	"go-chat/enums"
	"go-chat/internal/auth"
	"go-chat/storage"

	clog "github.com/charmbracelet/log"
	"github.com/lxzan/gws"
)

// const (
// 	pingInterval = 5 * time.Second
// 	pingWait     = 20 * time.Second
// )

var log = clog.WithPrefix("WS")

type wsHandler struct {
	repo     *wsRepo
	userRepo *auth.AuthRepo

	ctx     context.Context
	clients map[string]*gws.Conn
}

func NewWSHandler(repo *wsRepo, userRepo *auth.AuthRepo, ctx context.Context) *wsHandler {
	return &wsHandler{repo: repo, userRepo: userRepo, ctx: ctx, clients: make(map[string]*gws.Conn)}
}

func (c *wsHandler) OnOpen(conn *gws.Conn) {
	// XXX this is a workaround for the close error
	// its basically disable the deadline
	err := conn.SetDeadline(time.Time{})
	if err != nil {
		log.Warn("setdeadline", "err", err)
	}
	username := storage.Session.GetString(c.ctx, "username")
	c.clients[username] = conn
}

func (c *wsHandler) OnClose(conn *gws.Conn, err error) {
	log.Warn("connection closed", "err", err)
}

func (c *wsHandler) OnPing(conn *gws.Conn, payload []byte) {
	log.Info("ping received", "from", conn.RemoteAddr())
	// FIX recall SetDeadline is not working
	// err := conn.SetDeadline(time.Time{})
	// if err != nil {
	// 	log.Error("setdeadline", "err", err)
	// }
	err := conn.WritePong(nil)
	if err != nil {
		log.Error("write pong", "err", err)
	}
}

func (c *wsHandler) OnPong(conn *gws.Conn, payload []byte) {}

func (c *wsHandler) OnMessage(conn *gws.Conn, message *gws.Message) {
	defer message.Close()

	// for chrome
	if b := message.Bytes(); len(b) == 4 && string(b) == "ping" {
		c.OnPing(conn, nil)
		return
	}

	type Message struct {
		Type enums.MessageType `json:"type"`
		From string            `json:"from"`
		To   string            `json:"to"`
		Data string            `json:"data"`
	}

	msg := &Message{}

	err := json.Unmarshal(message.Bytes(), msg)
	if err != nil {
		log.Warn("unmarshal message", "err", err)
	}

	log.Info("message received", "message", msg)

	// shitty, refactor this later
	exits, err := c.userRepo.CheckUsername(msg.From, context.TODO())
	if err != nil {
		log.Error("check username", "err", err)
		conn.WriteString("server error")
		return
	}

	if !exits {
		log.Warn("user not found", "username", msg.From)
		conn.WriteString("user in \"from\" field is not found")
		return
	}

	exits, err = c.userRepo.CheckUsername(msg.To, context.TODO())
	if err != nil {
		log.Error("check username", "err", err)
		conn.WriteString("server error")
		return
	}

	if !exits {
		log.Warn("user not found", "username", msg.To)
		conn.WriteString("user in \"to\" field is not found")
		return
	}

	conn.WriteString("you sent the message \"" + msg.Data + "\" to " + msg.To)

	to, ok := c.clients[msg.To]
	if !ok {
		log.Warn("user not connected", "username", msg.To)
		conn.WriteString("user not connected")
		return
	}

	if err = to.WriteString(msg.Data); err != nil {
		log.Error("write message to", "err", err)
	}
}
