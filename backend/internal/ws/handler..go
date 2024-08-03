package ws

import (
	"context"
	"encoding/json"
	"time"

	"go-chat/internal/auth"
	"go-chat/storage"

	clog "github.com/charmbracelet/log"
	"github.com/marifcelik/gws"
)

// const (
// 	pingInterval = 5 * time.Second
// 	pingWait     = 20 * time.Second
// )

var log = clog.WithPrefix("WS")

type wsHandler struct {
	repo     *wsRepo
	userRepo *auth.AuthRepo

	clients *gws.ConcurrentMap[string, *gws.Conn]
}

func NewWSHandler(repo *wsRepo, userRepo *auth.AuthRepo) *wsHandler {
	return &wsHandler{
		repo:     repo,
		userRepo: userRepo,
		clients:  gws.NewConcurrentMap[string, *gws.Conn](16, 128),
	}
}

func (c *wsHandler) OnOpen(conn *gws.Conn) {
	// XXX this is a workaround for the close error
	// its basically disable the deadline
	err := conn.SetDeadline(time.Time{})
	if err != nil {
		log.Warn("setdeadline", "err", err)
	}
	username := storage.Session.GetString(conn.Context(), "user")
	c.clients.Store(username, conn)
	log.Info("connection opened", "from", conn.RemoteAddr(), "username", username)
}

func (c *wsHandler) OnClose(conn *gws.Conn, err error) {
	username := storage.Session.GetString(conn.Context(), "user")
	sharding := c.clients.GetSharding(username)
	sharding.Lock()
	defer sharding.Unlock()

	if socket, ok := sharding.Load(username); ok {
		key0, exits := socket.Session().Load("websocketKey")
		if exits {
			key0 = key0.(string)
		}
		key1, exits := conn.Session().Load("websocketKey")
		if exits {
			key1 = key1.(string)
		}
		if key0 == key1 {
			sharding.Delete(username)
		}
	}

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

	msg := &MessageDTO{}

	if err := json.Unmarshal(message.Bytes(), msg); err != nil {
		log.Warn("unmarshal message", "err", err)
	}

	from, ok := storage.Session.Get(conn.Context(), "user").(string)
	if !ok {
		log.Warn("username not found in session")
		conn.WriteString("username not found in session")
		return
	}

	log.Info("message received", "from", from)

	log.Info("message received", "message", msg.Data)

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

	if to, ok := c.clients.Load(msg.To); ok {
		msgToSend, err := json.Marshal(msg)
		if err != nil {
			log.Error("marshal message", "err", err)
			return
		}
		err = to.WriteMessage(gws.OpcodeText, msgToSend)
		if err != nil {
			log.Error("write message to", "err", err)
		}
	} else {
		log.Warn("user not connected", "username", msg.To)
		conn.WriteString("user not connected")
	}

}
