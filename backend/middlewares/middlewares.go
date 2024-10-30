package middlewares

import (
	"context"
	"net/http"

	"go-chat/config"
	st "go-chat/storage"
	"go-chat/utils"

	"github.com/charmbracelet/log"
	"github.com/go-chi/chi/v5"
)

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userID := st.Session.GetString(r.Context(), config.C.SessionIDKey)

		log.Print("auth middleware", "userID", userID)
		sr := r.WithContext(context.WithValue(r.Context(), "userID", userID))

		if userID == "" {
			log.WithPrefix("AUTH MW").Warn("unauthorized request", "header", r.Header.Get(config.C.HeaderKey.Session))
			utils.ErrResp(w, http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, sr)
	})
}

// i can't use this middleware because if i do i have to parse the body twice
// TODO find a way to parse the body once or look for what other people do
func LoggedIn(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := st.Session.GetString(r.Context(), config.C.SessionIDKey)
		log.Info("userID: %v\n", user)
		if user != "" {
			st.Session.Put(r.Context(), "warn", st.Session.GetInt(r.Context(), "warn")+1)
			utils.JsonResp(w, utils.M{"warn": "you already logged in"}, http.StatusConflict)
			return
		}
		// log.Warn("loggedin middleware unimplemented")
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

		r.Header.Add(config.C.HeaderKey.Session, token)

		ctx := context.WithValue(context.Background(), chi.RouteCtxKey, r.Context().Value(chi.RouteCtxKey))
		nc, err := st.Session.Load(ctx, token)
		if err != nil {
			log.Error("session load", "err", err)
			utils.ErrResp(w, http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r.WithContext(nc))
	})
}

func UpgradeChecher(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		connHeader, upgradeHeader := r.Header.Get("connection"), r.Header.Get("upgrade")
		if !(utils.ContainsI(connHeader, "upgrade") && utils.ContainsI(upgradeHeader, "websocket")) {
			utils.ErrResp(w, http.StatusUpgradeRequired)
			return
		}
		next.ServeHTTP(w, r)
	})
}
