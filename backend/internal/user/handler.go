package user

import (
	"go-chat/utils"
	"net/http"

	clog "github.com/charmbracelet/log"
)

// TODO implement user handler

var log = clog.WithPrefix("USER")

type userHandler struct {
	repo *userRepo
}

func NewHandler(repo *userRepo) *userHandler {
	return &userHandler{repo}
}

func (h *userHandler) GetUserMessages(w http.ResponseWriter, r *http.Request) {
	log.Warn("GetUserMessages not implemented")
	utils.ErrResp(w, http.StatusNotImplemented)
}

func (h *userHandler) GetGroupMessages(w http.ResponseWriter, r *http.Request) {
	log.Warn("GetGroupMessages not implemented")
	utils.ErrResp(w, http.StatusNotImplemented)
}
