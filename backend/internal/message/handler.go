package message

import (
	"net/http"
	"time"

	"go-chat/models"
	"go-chat/utils"

	clog "github.com/charmbracelet/log"
	v "github.com/cohesivestack/valgo"
	"github.com/go-chi/chi/v5"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var log = clog.WithPrefix("MESSAGE")

type messageHandler struct {
	repo *messageRepo
}

func NewMessageHandler(repo *messageRepo) *messageHandler {
	return &messageHandler{repo: repo}
}

func (h *messageHandler) GetConversationMessages(w http.ResponseWriter, r *http.Request) {
	conversationIDParam := chi.URLParam(r, "conversationID")

	val := v.Is(v.String(conversationIDParam, "conversationID").Not().Blank().Not().Empty())
	if !val.Valid() {
		utils.JsonResp(w, utils.M{
			"status":  "error",
			"message": "Validation error",
			"data":    val.Error(),
		}, http.StatusBadRequest)
		return
	}

	conversationID, err := primitive.ObjectIDFromHex(conversationIDParam)
	if err != nil {
		utils.JsonResp(w, utils.M{
			"status":  "error",
			"message": "Invalid conversation id",
		}, http.StatusBadRequest)
		return
	}

	messages, err := h.repo.GetConversationMessages(r.Context(), conversationID)
	if err != nil {
		log.Error("get conversation messages", "err", err)
		utils.ErrResp(w, http.StatusInternalServerError)
		return
	}

	response := make([]MessageDTO, 0, len(messages))
	for _, m := range messages {
		response = append(response, h.createMessageDTO(m))
	}

	utils.JsonResp(w, utils.M{
		"status": "success",
		"data":   response,
	}, http.StatusOK)
}

func (h *messageHandler) GetUserConversations(w http.ResponseWriter, r *http.Request) {
	userIDParam := chi.URLParam(r, "userID")

	val := v.Is(v.String(userIDParam, "userID").Not().Blank().Not().Empty())
	if !val.Valid() {
		utils.JsonResp(w, utils.M{
			"status":  "error",
			"message": "Validation error",
			"data":    val.Error(),
		}, http.StatusBadRequest)
		return
	}

	userID, err := primitive.ObjectIDFromHex(userIDParam)
	if err != nil {
		utils.JsonResp(w, utils.M{
			"status":  "error",
			"message": "Invalid user id",
		}, http.StatusBadRequest)
		return
	}

	conversations, err := h.repo.GetUserConversations(r.Context(), userID)
	if err != nil {
		log.Error("get user conversations", "err", err)
		utils.ErrResp(w, http.StatusInternalServerError)
		return
	}

	log.Info("get user conversations", "user_id", userID, "conversations", conversations)

	response := h.createConversationDTO(conversations)

	utils.JsonResp(w, utils.M{
		"status": "success",
		"data":   response,
	}, http.StatusOK)
}

func (h *messageHandler) createMessageDTO(m models.Message) MessageDTO {
	result := MessageDTO{
		ID:        m.ID.Hex(),
		Content:   m.Content,
		SenderID:  m.SenderID.Hex(),
		CreatedAt: m.CreatedAt.Format(time.UnixDate),
	}

	return result
}

func (h *messageHandler) createConversationDTO(c []models.Conversation) []ConversationDTO {
	result := make([]ConversationDTO, 0, len(c))
	for _, c := range c {
		lastMsg := MessageDTO{
			ID:        c.LastMessage.ID.Hex(),
			Content:   c.LastMessage.Content,
			SenderID:  c.LastMessage.SenderID.Hex(),
			CreatedAt: c.LastMessage.CreatedAt.Format(time.UnixDate),
		}

		participants := make([]string, 0, len(c.Participants))
		for _, p := range c.Participants {
			participants = append(participants, p.Hex())
		}

		result = append(result, ConversationDTO{
			ID:           c.ID.Hex(),
			Participants: participants,
			LastMessage:  lastMsg,
			Name:         c.Name,
		})
	}

	return result
}
