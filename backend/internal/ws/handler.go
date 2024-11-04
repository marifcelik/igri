package ws

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"go-chat/config"
	"go-chat/enums"
	"go-chat/internal/auth"
	"go-chat/models"
	st "go-chat/storage"

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

var (
	ErrInvalidSenderID    = errors.New("invalid sender id")
	ErrInvalidRecipientID = errors.New("invalid recipient id")
	ErrSenderNotFound     = errors.New("user in \"senderID\" field is not found")
	ErrRecipientNotFound  = errors.New("user in \"recipientID\" field is not found")
	ErrorServer           = errors.New("server error")
)

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
	userID := st.Session.GetString(conn.Context(), config.C.SessionIDKey)
	h.clients.Store(userID, conn)
	log.Info("connection opened", "from", conn.RemoteAddr(), "userID", userID)
}

func (h *wsHandler) OnClose(conn *gws.Conn, err error) {
	username := st.Session.GetString(conn.Context(), config.C.SessionIDKey)
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

	incomingMessage := MessageDTO{}

	if err := json.Unmarshal(message.Bytes(), &incomingMessage); err != nil {
		log.Warn("unmarshal message", "err", err)
	}

	log.Info("message received", "from", conn.RemoteAddr(), "msg", incomingMessage)

	_, ok := st.Session.Get(conn.Context(), config.C.SessionIDKey).(string)
	if !ok {
		log.Warn("username not found in session")
		conn.WriteString("username not found in session")
		return
	}

	if !h.validateUserID(incomingMessage.SenderID, conn, ErrInvalidSenderID, ErrSenderNotFound) {
		return
	}

	resp := WSMessage{Type: enums.WSMessageResult}

	messageID, conversationID, recipient, err := h.saveMessage(incomingMessage, conn.Context())
	if err != nil {
		log.Error("save message", "err", err)
		resp.Data = ResultDTO{
			Status:         enums.ResultFailed,
			ConversationID: conversationID,
			MessageID:      messageID,
			Message:        err.Error(),
		}
		goto SendResponse
	}

	if to, ok := h.clients.Load(recipient.ID.Hex()); ok {
		messageToSend := WSMessage{
			Type: enums.WSMessageConversation,
			Data: incomingMessage,
		}

		jsonMessage, _ := json.Marshal(messageToSend)

		if err := to.WriteMessage(gws.OpcodeText, jsonMessage); err != nil {
			log.Error("write message to", "err", err)
			resp.Data = ResultDTO{
				Status:         enums.ResultFailed,
				ConversationID: conversationID,
				MessageID:      messageID,
				Message:        err.Error(),
			}
		}
	} else {
		log.Warn("user not connected", "userID", recipient)
	}

	resp.Data = ResultDTO{
		Status:         enums.ResultSuccess,
		ConversationID: conversationID,
		MessageID:      messageID,
	}

SendResponse:
	respJson, _ := json.Marshal(resp)
	conn.WriteMessage(gws.OpcodeText, respJson)
}

// saveMessage checks the type of the message and saves it to the database
// if the message is a new conversation, it creates a new conversation and saves the message
func (h *wsHandler) saveMessage(incomingMessage MessageDTO, ctx context.Context) (messageID, conversationID string, recipient models.User, err error) {
	senderID, err := primitive.ObjectIDFromHex(incomingMessage.SenderID)
	if err != nil {
		log.Error(ErrInvalidSenderID, "id", incomingMessage.SenderID, "err", err)
		err = fmt.Errorf("%s: %w", ErrInvalidSenderID.Error(), err)
		return
	}

	if !incomingMessage.Type.IsValid() {
		log.Warn("unknown message type", "type", incomingMessage.Type)
		err = errors.New("unknown message type")
		return
	}

	message := models.Message{
		SenderID: senderID,
		Content:  incomingMessage.Content,
		M: models.M{
			ID:        primitive.NewObjectID(),
			CreatedAt: time.Now(),
			Version:   1,
		},
	}

	recipient, err = h.establishConversation(senderID, &incomingMessage, &message, ctx)
	if err != nil {
		return
	}

	// TODO consider removing the id from the return value
	messageIDObject, err := h.repo.CreateMessage(message, ctx)
	if err != nil {
		err = fmt.Errorf("save message: %w", err)
		return
	}
	return messageIDObject.Hex(), message.ConversationID.Hex(), recipient, nil
}

// establishConversation checks if the conversation exists and creates a new one if it doesn't
// it also updates the last message of the conversation and adds conversationID to the message
//
// TODO this function should return `conversationID`, at this moment, app only has normal conversation
// so its fine until group conversation is added
func (h *wsHandler) establishConversation(senderID primitive.ObjectID, incomingMessage *MessageDTO, message *models.Message, ctx context.Context) (recipient models.User, err error) {
	if incomingMessage.ConversationID == "" {
		if incomingMessage.Type == enums.GroupConversation {
			err = errors.New("conversation id is required for group conversation")
			return
		}
		if incomingMessage.RecipientUsername == "" {
			err = errors.New("recipient id is required for new normal conversation")
			return
		}

		recipient, err = h.userRepo.GetUserByUsername(incomingMessage.RecipientUsername, ctx)
		if err != nil {
			if errors.Is(err, mongo.ErrNoDocuments) {
				log.Error("invalid recipient username", "username", incomingMessage.RecipientUsername, "err", err)
				err = fmt.Errorf("invalid recipient username: %w", err)
				return
			}
			log.Error("check recipient username", "err", err)
			err = fmt.Errorf("check recipient username: %w", err)
			return
		}
		message.ConversationID = primitive.NewObjectID()

		conversation := models.Conversation{
			Type:         incomingMessage.Type,
			Participants: []primitive.ObjectID{senderID, recipient.ID},
			LastMessage:  message,
			M: models.M{
				ID:        message.ConversationID,
				CreatedAt: time.Now(),
				Version:   1,
			},
		}

		// TODO consider removing the id from the return value
		_, err = h.repo.CreateConversation(conversation, ctx)
		if err != nil {
			err = fmt.Errorf("save conversation: %w", err)
			return
		}
	} else {
		message.ConversationID, err = primitive.ObjectIDFromHex(incomingMessage.ConversationID)
		if err != nil {
			err = fmt.Errorf("invalid conversation id: %w", err)
			return
		}

		if err = h.repo.UpdateLastMessageOfConversation(*message, ctx); err != nil {
			err = fmt.Errorf("update conversation: %w", err)
			return
		}

		var recipientID primitive.ObjectID
		recipientID, err = h.repo.GetRecipientIDByConversationID(message.ConversationID, message.SenderID, ctx)
		if err != nil {
			err = fmt.Errorf("get recipient id: %w", err)
			return
		}
		recipient, err = h.userRepo.GetUserByID(recipientID, ctx)
	}

	return
}

// validateUserID checks if the user ID is valid and exists.
// errMessages[0] is the error message for invalid user id and
// errMessages[1] is the error message for user not found
func (h *wsHandler) validateUserID(userIDHex string, conn *gws.Conn, errMessages ...error) bool {
	idErr := errors.New("invalid user id")
	if len(errMessages) > 0 {
		idErr = errMessages[0]
	}
	notFoundErr := errors.New("user not found")
	if len(errMessages) > 1 {
		notFoundErr = errMessages[1]
	}

	userID, err := primitive.ObjectIDFromHex(userIDHex)
	if err != nil {
		log.Error(idErr, "id", userIDHex, "err", err)
		conn.WriteString(idErr.Error())
		return false
	}

	exists, err := h.userRepo.CheckUserID(userID, conn.Context())
	if err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		log.Error("check user id", "err", err)
		conn.WriteString(ErrorServer.Error())
		return false
	}
	if !exists {
		log.Warn(notFoundErr, "userID", userID.String())
		conn.WriteString(notFoundErr.Error())
		return false
	}

	return true
}
