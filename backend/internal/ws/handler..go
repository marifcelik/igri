package ws

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"go-chat/enums"
	"go-chat/internal/auth"
	"go-chat/models"
	"go-chat/storage"

	clog "github.com/charmbracelet/log"
	"github.com/marifcelik/gws"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
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

func (h *wsHandler) OnOpen(conn *gws.Conn) {
	// XXX this is a workaround for the close error
	// its basically disable the deadline
	err := conn.SetDeadline(time.Time{})
	if err != nil {
		log.Warn("setdeadline", "err", err)
	}
	userID := storage.Session.GetString(conn.Context(), "user")
	h.clients.Store(userID, conn)
	log.Info("connection opened", "from", conn.RemoteAddr(), "userID", userID)
}

func (h *wsHandler) OnClose(conn *gws.Conn, err error) {
	username := storage.Session.GetString(conn.Context(), "user")
	sharding := h.clients.GetSharding(username)
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

func (h *wsHandler) OnPing(conn *gws.Conn, payload []byte) {
	log.Info("ping received", "from", conn.RemoteAddr())
	// FIX recall SetDeadline is not working
	// err := conn.SetDeadline(time.Time{})
	// if err != nil {
	// 	log.Error("setdeadline", "err", err)
	// }
	if err := conn.WritePong(nil); err != nil {
		log.Error("write pong", "err", err)
	}
}

func (h *wsHandler) OnPong(conn *gws.Conn, payload []byte) {}

func (h *wsHandler) OnMessage(conn *gws.Conn, message *gws.Message) {
	defer message.Close()

	// for chrome
	if b := message.Bytes(); len(b) == 4 && string(b) == "ping" {
		h.OnPing(conn, nil)
		return
	}

	msg := MessageDTO{}

	if err := json.Unmarshal(message.Bytes(), &msg); err != nil {
		log.Warn("unmarshal message", "err", err)
	}

	from, ok := storage.Session.Get(conn.Context(), "user").(string)
	if !ok {
		log.Warn("username not found in session")
		conn.WriteString("username not found in session")
		return
	}

	log.Info("message received", "from", from, "message", msg)

	if !(h.validateUserID(msg.SenderID, conn, "invalid sender id", "user in \"from\" field is not found") &&
		h.validateUserID(msg.ReceiverID, conn, "invalid receiver id", "user in \"to\" field is not found")) {
		return
	}

	if err := h.saveMessage(msg, conn.Context()); err != nil {
		// TODO handle invalid id error and log it
		log.Error("save message", "err", err)
		conn.WriteString("server error")
		return
	}

	if to, ok := h.clients.Load(msg.ReceiverID); ok {
		if err := to.WriteString(message.Data.String()); err != nil {
			log.Error("write message to", "err", err)
		}
	} else {
		log.Warn("user not connected", "username", msg.ReceiverID)
		conn.WriteString("user not connected")
	}
}

// saveMessage checks the type of the message and saves it to the database
func (h *wsHandler) saveMessage(msg MessageDTO, ctx context.Context) error {
	sender, err := primitive.ObjectIDFromHex(msg.SenderID)
	if err != nil {
		log.Error("invalid sender id", "id", msg.SenderID, "err", err)
		return fmt.Errorf("invalid sender id: %w", err)
	}

	switch msg.Type {
	case enums.NormalMessage:
		receiver, err := primitive.ObjectIDFromHex(msg.ReceiverID)
		if err != nil {
			return fmt.Errorf("invalid receiver id: %w", err)
		}

		message := models.UserMessage{
			SenderID:   sender,
			ReceiverID: receiver,
			Message:    msg.Data,
		}
		return h.repo.SaveMessage(message, ctx)

	case enums.GroupMessage:
		group, err := primitive.ObjectIDFromHex(msg.GroupID)
		if err != nil {
			return fmt.Errorf("invalid group id: %w", err)
		}

		message := models.GroupMessage{
			SenderID: sender,
			GroupID:  group,
			Message:  msg.Data,
		}
		return h.repo.SaveGroupMessage(message, ctx)

	default:
		log.Warn("unknown message type", "type", msg.Type)
		return errors.New("unknown message type")
	}
}

// validateUserID checks if the user ID is valid and exists.
// errMessages[0] is the error message for invalid user id and
// errMessages[1] is the error message for user not found
func (h *wsHandler) validateUserID(userIDHex string, conn *gws.Conn, errMessages ...string) bool {
	idErr := "invalid user id"
	if len(errMessages) > 0 {
		idErr = errMessages[0]
	}
	notFoundErr := "user not found"
	if len(errMessages) > 1 {
		notFoundErr = errMessages[1]
	}

	userID, err := primitive.ObjectIDFromHex(userIDHex)
	if err != nil {
		log.Error(idErr, "id", userIDHex, "err", err)
		conn.WriteString(idErr)
		return false
	}

	exists, err := h.userRepo.CheckUserID(userID, conn.Context())
	if err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		log.Error("check user id", "err", err)
		conn.WriteString("server error")
		return false
	}
	if !exists {
		log.Warn(notFoundErr, "userID", userID.String())
		conn.WriteString(notFoundErr)
		return false
	}

	return true
}
