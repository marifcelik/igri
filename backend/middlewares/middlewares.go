package middlewares

import (
	"context"
	"net/http"

	st "go-chat/storage"
	"go-chat/utils"

	"github.com/charmbracelet/log"
	"github.com/go-chi/chi/v5"
)

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := st.Session.GetString(r.Context(), "user")

		if user == "" {
			log.WithPrefix("AUTH MW").Warn("unauthorized request", "from", utils.GetIPAddr(r))
			utils.ErrResp(w, http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func WsHeader(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.URL.Query().Get("token")
		if token == "" {
			utils.ErrResp(w, http.StatusUnauthorized)
			return
		}
		r.Header.Add("Authorization", "Bearer "+token)
		ctx := context.WithValue(context.Background(), chi.RouteCtxKey, r.Context().Value(chi.RouteCtxKey))
		nc, err := st.Session.Load(ctx, token)
		if err != nil {
			log.Error("session load", "err", err)
		}
		sr := r.WithContext(nc)
		next.ServeHTTP(w, sr)
	})
}

func UpgradeChecher(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !(utils.ContainsI(r.Header.Get("connection"), "upgrade") && utils.ContainsI(r.Header.Get("upgrade"), "websocket")) {
			utils.ErrResp(w, http.StatusUpgradeRequired)
			return
		}
		next.ServeHTTP(w, r)
	})
}
