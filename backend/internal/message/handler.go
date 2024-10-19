package message

import (
	"net/http"

	"go-chat/utils"

	clog "github.com/charmbracelet/log"
)

var log = clog.WithPrefix("MESSAGE")

type messageHandler struct {
	repo *messageRepo
}

func NewMessageHandler(repo *messageRepo) *messageHandler {
	return &messageHandler{repo: repo}
}

func (h *messageHandler) GetUserMessages(w http.ResponseWriter, r *http.Request) {
	utils.ErrResp(w, http.StatusNotImplemented)
}

func (h *messageHandler) GetGroupMessages(w http.ResponseWriter, r *http.Request) {
	utils.ErrResp(w, http.StatusNotImplemented)
}
