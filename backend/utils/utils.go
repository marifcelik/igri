package utils

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/charmbracelet/log"
)

type M map[string]any

func GetIPAddr(r *http.Request) string {
	switch {
	case r.RemoteAddr == "127.0.0.1" || r.RemoteAddr == "::1":
		return r.RemoteAddr
	case len(r.Header.Get("X-Forwarded-For")) > 0:
		return r.Header.Get("X-Forwarded-For")
	case len(r.Header.Get("X-Real-IP")) > 0:
		return r.Header.Get("X-Real-IP")
	default:
		return strings.Split(r.RemoteAddr, ":")[0]
	}
}

func JsonResp(w http.ResponseWriter, data any, status ...int) {
	w.Header().Set("Content-Type", "application/json")
	if len(status) > 0 {
		w.WriteHeader(status[0])
	}
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		log.Error("json encode", "err", err)
		InternalErrResp(w, err)
	}
}

// TODO follow the google api design guide. see: https://cloud.google.com/apis/design/errors
func ErrResp(w http.ResponseWriter, status int, err ...error) {
	var text string
	if len(err) > 0 && err[0] != nil {
		text = err[0].Error()
	} else {
		text = http.StatusText(status)
	}

	http.Error(w, text, status)
}

func InternalErrResp(w http.ResponseWriter, err error) {
	http.Error(w, err.Error(), http.StatusInternalServerError)
}

// Check the error and exit if its not nil.
// The parameters after the second parameter will be joined into a single string
func CheckErr(err error, msgParams ...string) {
	msg := strings.Join(msgParams, ", ")

	if err != nil {
		if msg != "" {
			log.Fatal(msg, "err", err)
		} else {
			log.Fatal(err)
		}
	}
}
