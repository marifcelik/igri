package middlewares

import (
	"fmt"
	"net/http"

	st "go-chat/storage"
	"go-chat/utils"

	"github.com/charmbracelet/log"
)

// TODO
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log := log.WithPrefix("AUTH MW")
		log.Print("auth", "token", r.Header.Get("Authorization"))

		fmt.Printf("st.Session.Keys(r.Context()): %v\n", st.Session.Keys(r.Context()))

		user := st.Session.GetString(r.Context(), "user")
		if user == "" {
			log.Warn("unauthorized request", "from", utils.GetIPAddr(r))
			utils.ErrResp(w, http.StatusUnauthorized)
			return
		}
		log.Print("TODO")
		next.ServeHTTP(w, r)
	})
}

func WsHeaderMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.URL.Query().Get("access_token")
		if token == "" {
			utils.ErrResp(w, http.StatusUnauthorized)
			return
		}
		r.Header.Add("Authorization", "Bearer "+token)
		next.ServeHTTP(w, r)
	})
}

func CheckIsUpgrade(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Connection") != "Upgrade" || r.Header.Get("Upgrade") != "Websocket" {
			next.ServeHTTP(w, r)
			return
		}
		utils.ErrResp(w, http.StatusUpgradeRequired)
	})
}
