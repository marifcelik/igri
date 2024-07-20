package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"strings"

	clog "github.com/charmbracelet/log"
)

var log = clog.WithPrefix("UTILS")

type M map[string]any

// ContainsI checks if a string contains a substring case-insensitively
func ContainsI(s, substr string) bool {
	return strings.Contains(strings.ToLower(s), strings.ToLower(substr))
}

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
func ErrResp(w http.ResponseWriter, status int, err ...any) {
	var text string
	if len(err) > 0 && err[0] != nil {
		switch t := err[0].(type) {
		case string:
			text = t
		case error:
			text = t.Error()
		default:
			log.Warnf("unknown error type: %T", t)
			text = http.StatusText(status)
		}
	} else {
		text = http.StatusText(status)
	}

	http.Error(w, text, status)
}

func InternalErrResp(w http.ResponseWriter, err error) {
	http.Error(w, err.Error(), http.StatusInternalServerError)
}

// CopyFields copies fields from src to dst.
// dst and src must be pointers to structs
// FIX doesn't work
func CopyFields(dst, src any) error {
	dstVal := reflect.ValueOf(dst)
	srcVal := reflect.ValueOf(src)

	if dstVal.Kind() != reflect.Ptr || dstVal.Elem().Kind() != reflect.Struct {
		return fmt.Errorf("dst must be a pointer to a struct")
	}
	if srcVal.Kind() != reflect.Ptr || srcVal.Elem().Kind() != reflect.Struct {
		return fmt.Errorf("src must be a pointer to a struct")
	}

	dstType := dstVal.Elem().Type()
	srcType := srcVal.Elem().Type()

	for i := 0; i < dstType.NumField(); i++ {
		dstField := dstType.Field(i)
		srcField, ok := srcType.FieldByName(dstField.Name)
		if !ok {
			log.Warnf("field %s not found in src", dstField.Name)
			continue
		}

		log.Printf("dst type: %s, src type: %s", dstField.Type, srcField.Type)
		if dstField.Type != srcField.Type {
			log.Warnf("field %s has different types in dst and src", dstField.Name)
			return fmt.Errorf("field %s has different types in dst and src", dstField.Name)
		}

		dstVal.Elem().Field(i).Set(srcVal.Elem().FieldByName(dstField.Name))
	}

	return nil
}
