package utils

import (
	"os"
	"strings"

	"github.com/charmbracelet/log"
)

func GetDBConnStr() string {
	connStr := os.Getenv("DB_CONN_STR")
	if connStr == "" {
		connStr = "ws://localhost:8080"
	}
	return connStr
}

func GetAppEnv() string {
	appEnv := os.Getenv("APP_ENV")
	if appEnv == "" {
		appEnv = "dev"
	}
	return appEnv
}

// Check the error and exit if its not nil.
// The parameters after the second parameter will be joined into a single string
func CheckErr(err error, msgParams ...string) {
	msg := strings.Join(msgParams, "")

	if err != nil {
		if msg != "" {
			log.Fatal(msg, "err", err)
		} else {
			log.Fatal(err)
		}
	}
}
