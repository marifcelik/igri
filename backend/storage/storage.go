package storage

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/http"
	"strings"

	"go-chat/config"
	"go-chat/utils"

	"github.com/alexedwards/scs/redisstore"
	"github.com/alexedwards/scs/v2"
	clog "github.com/charmbracelet/log"
	"github.com/gomodule/redigo/redis"
)

var Session *mySession
var log = clog.WithPrefix("STORAGE")

type mySession struct {
	*scs.SessionManager
}

func (s *mySession) LoadAndServeHeader(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		key := "Authorization"
		expiryKey := "X-Session-Expiry"

		bearer := r.Header.Get(key)
		token := strings.TrimPrefix(bearer, "Bearer ")
		fmt.Printf("token: %v\n", token)

		ctx, err := s.Load(r.Context(), token)
		if err != nil {
			log.Error("session load", "err", err)
			utils.ErrResp(w, http.StatusInternalServerError)
			return
		}
		bw := &bufferedResponseWriter{ResponseWriter: w}
		sr := r.WithContext(ctx)
		next.ServeHTTP(bw, sr)

		if s.Status(ctx) == scs.Modified {
			token, expiry, err := s.Commit(ctx)
			if err != nil {
				log.Error("session commit", "err", err)
				utils.ErrResp(w, http.StatusInternalServerError)
				return
			}

			// TODO extend expiry time on session commit
			w.Header().Set(key, "Bearer "+token)
			w.Header().Set(expiryKey, expiry.Format(http.TimeFormat))
		}

		if bw.code != 0 {
			w.WriteHeader(bw.code)
		}
		w.Write(bw.buf.Bytes())
	})
}

type bufferedResponseWriter struct {
	http.ResponseWriter
	buf  bytes.Buffer
	code int
}

func (bw *bufferedResponseWriter) Write(b []byte) (int, error) {
	return bw.buf.Write(b)
}

func (bw *bufferedResponseWriter) WriteHeader(code int) {
	bw.code = code
}

func (bw *bufferedResponseWriter) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	h, ok := bw.ResponseWriter.(http.Hijacker)
	if !ok {
		return nil, nil, errors.New("http.Hijacker not implemented")
	}
	return h.Hijack()
}

func init() {
	pool := &redis.Pool{
		MaxIdle: 10,
		Dial: func() (redis.Conn, error) {
			return redis.DialURL(config.C.RedisURL)
		},
	}

	Session = &mySession{scs.New()}
	Session.Store = redisstore.New(pool)
	Session.Lifetime = config.GetExpirationTime()
	Session.IdleTimeout = config.GetIdleTimeout()
}
