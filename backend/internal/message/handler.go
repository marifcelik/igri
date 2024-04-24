package message

import (
	"go-chat/utils"
	"net/http"
)

func handleGetUserMessages(w http.ResponseWriter, r *http.Request) {
	utils.ErrResp(w, http.StatusNotImplemented)
}

func handleGetMessage(w http.ResponseWriter, r *http.Request) {
	utils.ErrResp(w, http.StatusNotImplemented)
}
