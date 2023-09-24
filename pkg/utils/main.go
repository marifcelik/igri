package utils

import (
	"os"
	"strings"

	"github.com/charmbracelet/log"
	"github.com/joho/godotenv"
)

func init() {
	if GetAppEnv() != "prod" {
		os.Setenv("APP_ENV", "dev")
		if err := godotenv.Load(); err != nil {
			log.Error(err)
		}
	}
}

func GetAppEnv() string {
	appEnv := os.Getenv("APP_ENV")
	if appEnv == "" {
		appEnv = "dev"
	}
	return appEnv
}

func GetDBURL() string {
	connStr := os.Getenv("DATABASE_URL")
	if connStr == "" {
		connStr = "ws://localhost:8080"
	}
	return connStr
}

// if REDIS_URL env is not set then that returns "" because 'redis.Config' defaults
// are same the default connection url
func GetRedisURL() string {
	connStr := os.Getenv("REDIS_URL")
	return connStr
}

func GetListenAddr() string {
	host, port := os.Getenv("HOST"), os.Getenv("PORT")
	if port == "" {
		port = "8085"
	}
	return host + ":" + port
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
